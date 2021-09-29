package vault

import (
	"context"
	"fmt"
	"log"
	"mongo-vault/notify"
	"strings"

	"github.com/hashicorp/vault/api"
	vault "github.com/mittwald/vaultgo"
)

var ErrRoleIsEmpty = fmt.Errorf("role can not be empty")
var ErrInvalidIncrement = fmt.Errorf("invalid increment value")

type VaultClient struct {
	vClient  *vault.Client
	firstRun bool
}

// New will instantiate a client that will automatically authenticate with kubernetes using
// kubernetes role and accountservice jwt token obtained from the file /run/secrets/kubernetes.io/serviceaccount/token
func New(vaultAddress string, kubernetesRole string) (*VaultClient, error) {
	client, err := vault.NewClient(
		vaultAddress,
		vault.WithCaPath(""),
		vault.WithKubernetesAuth(
			kubernetesRole,
		),
	)
	if err != nil {
		return nil, err
	}

	return &VaultClient{
		vClient:  client,
		firstRun: true,
	}, nil
}

// NewWithToken will instantiate a client that will authenticate using token (for local test)
func NewWithToken(vaultAddress string, token string) (*VaultClient, error) {
	client, err := vault.NewClient(
		vaultAddress,
		vault.WithCaPath(""),
		vault.WithAuthToken(token),
	)
	if err != nil {
		return nil, err
	}

	return &VaultClient{
		vClient:  client,
		firstRun: true,
	}, nil
}

//Read will read credentials from vault and return the first credentials with a notifier (for the future renewed credentials)
// each Read has its own notifier, because it depends on the role, and usually each role connects to a specific database.
func (v *VaultClient) Read(ctx context.Context, role string, increment int) (map[string]interface{}, notify.Notifier, error) {

	if strings.TrimSpace(role) == "" {
		return nil, nil, ErrRoleIsEmpty
	}

	if increment <= 0 {
		return nil, nil, ErrInvalidIncrement
	}

	firstResult := make(chan map[string]interface{})

	notifier := notify.New()

	go func() {
		for {
			select {
			// exit before event starting if context was canceled
			case <-ctx.Done():
				log.Println("context canceled")
				return

			default:
				err := v.renew(ctx, role, increment, notifier, firstResult)
				if err != nil {
					log.Printf("error on Renew: %v", err)
				}
			}
		}

	}()
	return <-firstResult, notifier, nil
}

// renew will renew the ttl of the lease until reach the max_ttl , then it will re-read again (ROTATE).
// the secrets are sent to the channel on the first run, then on the notifier in the next runs.
func (v *VaultClient) renew(ctx context.Context, role string, increment int, notifier notify.Notifier, result chan map[string]interface{}) error {
	var response api.Secret
	err := v.vClient.Read([]string{role}, &response, &vault.RequestOptions{
		SkipRenewal: false,
	})

	if err != nil {

		if v.firstRun {
			close(result)
		}
		notifier.Notify(notify.MongoEvent{
			Data:  nil,
			Error: err,
		})
		return err
	}
	if v.firstRun {
		v.firstRun = false
		result <- response.Data
		close(result)
	} else {
		//this is a blocking operation, depends on how the onNotify is implemented by the observer.
		notifier.Notify(notify.MongoEvent{
			Data:  response.Data,
			Error: nil,
		})
	}

	secretesTokenWatcher := api.LifetimeWatcherInput{
		Secret:        &response,
		Increment:     increment,
		RenewBehavior: api.RenewBehaviorIgnoreErrors,
	}

	// this will not return error, because we are passing everything.
	watcher, err := v.vClient.NewLifetimeWatcher(&secretesTokenWatcher)
	if err != nil {
		log.Println("error creating watcher: ", err)
		return err
	}

	go watcher.Start()

	for {
		select {
		// gracefully stop the goroutine in case the context canceled.
		case <-ctx.Done():
			watcher.Stop()
			return nil

		case rawData := <-watcher.RenewCh():
			log.Printf("received renewal Warning: %+v \n", rawData.Secret.Warnings)

			// this will return when max_ttl reached
		case er := <-watcher.DoneCh():
			log.Println("exiting watcher, error: ", er)
			watcher.Stop()
			return er
		}
	}
}

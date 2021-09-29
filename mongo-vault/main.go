package main

import (
	"context"
	"log"
	repository2 "mongo-vault/repository"
	"mongo-vault/vault"
	"time"
)

var (
	address        = "https://127.0.0.1:8200"
	dbRole         = "database/creds/db-app"
	kubernetesRole = "demo"
	mongodbURI     = "localhost:27017/test"
)

func main() {

	token := "s.xxxxxxxxxxxxxxxx"
	//vaultClient, err := vault.New(address, kubernetesRole)
	vaultClient, err := vault.NewWithToken(address, token)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data, notifier, err := vaultClient.Read(ctx, dbRole, 60)
	if err != nil || data == nil {
		log.Fatalf("was not able to get secrets, check logs, err: %v", err)
	}

	dbRepository := repository2.NewRepository(mongodbURI)
	err = dbRepository.ReplaceClient(data)
	if err != nil {
		panic(err)
	}

	// registering the repository in the notifier to automatically replace the mongoClient
	// with the new credentials received from vault
	notifier.Register(dbRepository)
	time.Sleep(10 * time.Minute)
}

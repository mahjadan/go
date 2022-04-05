package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"go-proto-poc/pkg/model"
	"log"
)

const AuthenticationClientsColName = "authentication_clients"

type Repository struct {
	Client *firestore.Client
}

func NewRepository(client *firestore.Client) Repository {
	return Repository{
		Client: client,
	}
}

func (r Repository) CreateAuthentication(ctx context.Context, authClient model.AuthenticationClient) (model.AuthenticationClient, error) {
	persistence, clientId := toPersistence(authClient)

	doc := r.Client.Collection(AuthenticationClientsColName).Doc(clientId)
	writeResult, err := doc.Create(ctx, persistence)
	log.Println(writeResult)
	if err != nil {
		return model.AuthenticationClient{}, err
	}
	return toModel(persistence), nil
}

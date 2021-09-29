package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mongo-vault/notify"
	"sync"
)

type repository struct {
	sync.Mutex
	client *mongo.Client
	dbURI  string
}

func NewRepository(uri string) *repository {
	return &repository{
		dbURI:  uri,
		client: &mongo.Client{},
	}
}

func (r *repository) OnNotify(event notify.MongoEvent) {
	if event.Error != nil {
		log.Fatalf("error onNotify: %v", event.Error)
	}
	err := r.ReplaceClient(event.Data)
	if err != nil {
		log.Fatalf("error onNofity: %v", err)
	}
}

func (r *repository) ReplaceClient(data map[string]interface{}) error {
	client, err := getClient(data, r.dbURI)
	if err != nil {
		return err
	}
	r.Lock()
	defer r.Unlock()
	*r.client = *client
	return nil
}

func getClient(data map[string]interface{}, uri string) (*mongo.Client, error) {
	fullURI := fmt.Sprintf("mongodb://%v:%v@%v", data["username"].(string), data["password"].(string), uri)
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(fullURI)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("error creating mongodb client: %v", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("error connecting to mongodb: %v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

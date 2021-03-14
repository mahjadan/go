package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var client2 *mongo.Client

// var resource *dockertest.Resource
// var resource2 *dockertest.Resource
var pool *dockertest.Pool

func TestMain(m *testing.M) {
	var err error
	pool, err = dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	resource, err := getClient(pool, "3.6")
	resource2, err := getWiremock(pool, "wiremock")
	fmt.Println("got : ", resource, client, err)
	fmt.Println("got2 : ", resource2, client2, err)
	code := m.Run()

	purgeResouce(resource, resource2)

	os.Exit(code)
}

func TestHome(t *testing.T) {
	e := client.Ping(context.TODO(), nil)
	fmt.Println("after ping client: ", e)
	dbs, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		t.Error("got ERROR : ", err)
	}
	t.Logf("1 - databsenames : %v ", dbs)
}

// func TestHome2(t *testing.T) {
// 	e := client2.Ping(context.TODO(), nil)
// 	fmt.Println("after ping client2 : ", e)
// 	dbs, err := client2.ListDatabaseNames(context.TODO(), bson.M{})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Logf("2 - databsenames : %v ", dbs)
// }

func getClient(pool *dockertest.Pool, mongoVersion string) (*dockertest.Resource, error) {

	// pulls an image, creates a container based on it and runs it
	r, err := pool.Run("mongo", mongoVersion, nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
		return nil, err
	}

	fmt.Println(mongoVersion+" - connecting to : ", r.GetPort("27017/tcp"))
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error

		clientOptions := options.Client().ApplyURI("mongodb://" + fmt.Sprintf("localhost:%s", r.GetPort("27017/tcp")))
		// Connect to MongoDB
		client, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			return err
		}

		return client.Ping(context.TODO(), nil)

	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return r, nil
}

func getWiremock(pool *dockertest.Pool, mongoVersion string) (*dockertest.Resource, error) {

	// pulls an image, creates a container based on it and runs it
	r, err := pool.Run("rodolpheche/wiremock", "2.27.2", nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
		return nil, err
	}

	fmt.Println(mongoVersion+" - connecting to : ", r.GetPort("8443/tcp"))
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {

		resp, err := http.Get("http://localhost:" + r.GetPort("8443/tcp") + "/__admin")
		if err != nil {
			fmt.Printf("trying to connect on localhost:%s, got : %v \n", r.GetPort("8443/tcp"), err)
			return err
		}

		fmt.Println("status: ", resp.StatusCode)
		rs, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("RESPONSE: %v", rs)
		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return r, nil
}
func purgeResouce(reouseces ...*dockertest.Resource) {

	fmt.Println("removing resources.")
	for _, resource := range reouseces {

		// You can't defer this because os.Exit doesn't care for defer
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}

	}
}

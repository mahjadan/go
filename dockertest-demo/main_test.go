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

var (
	client       *mongo.Client
	pool         *dockertest.Pool
	wiremockPort string
)

func TestMain(m *testing.M) {
	var err error
	pool, err = dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	resource, err := getClient(pool, "3.6")
	resource2, err := getWiremock(pool, "latest")

	code := m.Run()

	purgeResource(resource, resource2)
	//_,_=resource, resource2
	os.Exit(code)
}

func TestMongoContainer(t *testing.T) {
	e := client.Ping(context.TODO(), nil)
	fmt.Println("after ping client: ", e)
	dbs, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		t.Error("got ERROR : ", err)
	}
	t.Logf("1 - databsenames : %v ", dbs)
}

func TestWiremockContainer(t *testing.T) {
	url := "http://localhost:" + wiremockPort + "/__admin"
	fmt.Println("calling wiremock on : ", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("calling wiremock err : ", err)
		t.Error(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	t.Logf("2 - wiremock resp : %s \n", bytes)
}

func getClient(pool *dockertest.Pool, mongoVersion string) (*dockertest.Resource, error) {

	// pulls an image, creates a container based on it and runs it
	r, err := pool.Run("mongo", mongoVersion, nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
		return nil, err
	}

	fmt.Printf("mongo-%s - connecting to : %s \n", mongoVersion, r.GetPort("27017/tcp"))
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

func getWiremock(pool *dockertest.Pool, tag string) (*dockertest.Resource, error) {

	// pulls an image, creates a container based on it and runs it
	r, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "rodolpheche/wiremock",
		Tag:        tag,
		Mounts: []string{
			//mounting /tmp/tx local machine to /home/wiremock in container
			"/tmp/tx:/home/wiremock",
		},
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
		return nil, err
	}

	wiremockPort = r.GetPort("8080/tcp")
	fmt.Println("wiremock - connecting to : ", wiremockPort)
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {

		resp, err := http.Get("http://localhost:" + wiremockPort + "/__admin")
		if err != nil {
			fmt.Printf("trying to connect on localhost:%s, got : %v \n", wiremockPort, err)
			return err
		}

		fmt.Println("status: ", resp.StatusCode)
		rs, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("RESPONSE: %s", rs)
		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return r, nil
}
func purgeResource(resources ...*dockertest.Resource) {

	fmt.Println("removing resources.")
	for _, resource := range resources {
		// You can't defer this because os.Exit doesn't care for defer
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}
}

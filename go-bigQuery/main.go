package main

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/cloudkms/v1"
	"io"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

// implicit uses Application Default Credentials to authenticate.
func implicit(projectId string) {
	ctx := context.Background()

	// For API packages whose import path is starting with "cloud.google.com/go",
	// such as cloud.google.com/go/storage in this case, if there are no credentials
	// provided, the client library will look for credentials in the environment.
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer storageClient.Close()

	it := storageClient.Buckets(ctx, projectId)
	for {
		bucketAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(bucketAttrs.Name)
	}

	// For packages whose import path is starting with "google.golang.org/api",
	// such as google.golang.org/api/cloudkms/v1, use NewService to create the client.
	kmsService, err := cloudkms.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_ = kmsService
}

// queryBasic demonstrates issuing a query and reading results.
func queryBasic(w io.Writer, projectID string) error {
	// projectID := "my-project-id"
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	fmt.Println("successfully created client.")
	defer client.Close()

	q := client.Query("SELECT * FROM `big-data-demo-319815.ct_event_30d.customer-from-json-file`")
	//"SELECT * FROM `big-data-demo-319815.ct_event_30d.customer-from-json-file` " +
	//"SELECT _id, email, lastModification, telephones.cellphone.number, telephones.cellphone.ddd,  FROM `big-data-demo-319815.ct_event_30d.customer-from-json-file` " +
	//	"LIMIT 100")

	// Location must match that of the dataset(s) referenced in the query.
	//q.Location = "US"
	// Run the query and print results when the query job is completed.
	job, err := q.Run(ctx)

	if err != nil {
		return err
	}
	js, err := job.Status(ctx)
	fmt.Println("job status: ", js.Statistics.StartTime)
	fmt.Println("job status: ", js.Statistics.EndTime)
	status, err := job.Wait(ctx)
	//fmt.Println("job status 2: ", status.Statistics.TotalBytesProcessed)
	if err != nil {
		return err
	}
	if err := status.Err(); err != nil {
		return err
	}
	it, err := job.Read(ctx)
	for {
		//var row []bigquery.Value
		var rowMap map[string]bigquery.Value
		//err := it.Next(&row)
		err := it.Next(&rowMap)
		fmt.Println("total rows: ", it.TotalRows)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		//printList( row)
		printMap(rowMap)
	}
	return nil
}

func printList(row []bigquery.Value) {
	for inx, value := range row {
		fmt.Printf("%d - %v\n", inx, value)
	}
}

func printMap(rowMap map[string]bigquery.Value) {
	bytes, err := json.MarshalIndent(rowMap, " ", "  ")
	if err != nil {
		fmt.Printf("error MARSHALLINIG: %v\n", err)
	}
	fmt.Println("result: ", string(bytes))
}

func main() {
	// need to set GOOGLE_APPLICATION_CREDENTIALS env variable with credentials
	// export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/service-account-file.json"

	fmt.Println("starting...")
	projectID := "big-data-demo-319815"
	implicit(projectID)
	fmt.Println("successfully authenticated..")
	err := queryBasic(os.Stdout, projectID)
	if err != nil {
		fmt.Printf("error while Querying : %v", err)
	}
}

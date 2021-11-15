package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// run program -input email-to-validate.csv -workers 5 -timeout 1
// will split the file into 5 parts and run each part in a seperate goroutine to validate the domain.

var timeout time.Duration

func main() {
	inputFileName := flag.String("input", "data.csv", "name of the csv input file.")
	numOfWorkers := flag.Int("workers", 1, "number of workers")
	timeoutArg := flag.Int("timeout", 1, "number of seconds waiting for domain validation")
	flag.Parse()
	timeout = time.Duration(*timeoutArg) * time.Second

	fmt.Println("reading from : ", *inputFileName)
	fmt.Println("starting with total number of workers: ", *numOfWorkers)
	fmt.Println("with timeout: ", timeout)

	f, err := os.Open(*inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	size := len(data) / *numOfWorkers

	fmt.Println("size of each part is:", size)

	numOfParts := *numOfWorkers
	//adding one more part of data if len(data) is odd number
	if size*numOfParts < len(data) {
		numOfParts += 1
	}

	wg := sync.WaitGroup{}

	for i := 0; i < numOfParts; i++ {

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			outputFilename := "new-list-" + strconv.Itoa(i+1) + ".csv"
			csvFile, err := os.Create(outputFilename)
			if err != nil {
				log.Fatalf("failed creating file: %s", err)
			}
			csvwriter := csv.NewWriter(csvFile)

			// split data
			workerData := getSlice(data, size, i)
			processData(workerData, csvwriter, outputFilename)
			csvFile.Close()
		}(i)
	}

	wg.Wait()
	fmt.Println("validation finished...")

}

func validateEmail(email string) bool {

	at := strings.LastIndex(email, "@")
	if at >= 0 {
		_, domain := email[:at], email[at+1:]

		r := &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: time.Millisecond * time.Duration(100),
				}
				return d.DialContext(ctx, network, "8.8.8.8:53")
			},
		}
		ctx, _ := context.WithTimeout(context.Background(), timeout)

		_, err := r.LookupHost(ctx, domain)
		if err == nil {
			return true
		}
	}
	return false
}
func getSlice(data [][]string, size int, partNum int) [][]string {

	from := partNum * size
	to := from + size
	if to > len(data) {
		return data[from:]
	}
	return data[from:to]
}

func processData(data [][]string, csvwriter *csv.Writer, outputFilename string) {
	var counter int64
	fmt.Printf("starting validation for %s, with size of : %d, first: %+v \n", outputFilename, len(data), data[0])

	for i, line := range data {
		fmt.Printf("%d,", i)
		if validateEmail(line[0]) {
			_ = csvwriter.Write(line)
			counter++
		}
	}

	fmt.Printf("found %d in %s \n", counter, outputFilename)
	csvwriter.Flush()
}

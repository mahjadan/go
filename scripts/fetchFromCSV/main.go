package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	url         = "http://localhost:8080/update"
	contentType = "application/json"
)

func main() {
	start := time.Now()
	var csvFilename string

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s csvFilename", filepath.Base(os.Args[0]))
		return
	}
	csvFilename = os.Args[1]

	csvFile, err := os.Open(csvFilename)
	if err != nil {
		fmt.Println("Error opening the file: ", err)
		return
	}
	errFilename := getErrorFilename(csvFilename)
	errFile, err := os.Create(errFilename)
	if err != nil {
		fmt.Println("Error while creating error file: ", err)
		return
	}

	run(csvFile, errFile)

	defer csvFile.Close()
	defer errFile.Close()

	fmt.Printf("This script took %s\n", time.Since(start))

	fmt.Printf("Please check %s for errors.\n", errFilename)
}

func run(csvFile io.ReadCloser, errFile io.WriteCloser) {
	errChannel := make(chan []string, 500)
	readChannel := make(chan []string, 50)

	var wg sync.WaitGroup
	wg.Add(3)

	fmt.Println("Start reading csv file.")
	// start routine to write on readChan and errChan , wg.Done() and close(readChan) but keep errChan open and continue writing to it.
	go readCSVFile(csvFile, errChannel, readChannel, &wg)

	fmt.Println("the script is now processing your file.")
	// start routine to read from readChan and write to errChan and then wg.Done() and close(errChan)
	go processFile(readChannel, errChannel, &wg)
	// start routine to read from errChan then wg.Done().
	go generateErrorReport(errFile, errChannel, &wg)

	// wait for all the 3 go routines to finish
	wg.Wait()
}

// return the file name without its path and extension
// and appends the error file name
func getErrorFilename(filename string) string {

	fn := filepath.Base(filename)

	if strings.Contains(fn, ".") {
		fn = strings.Split(fn, ".")[0]
	}
	return fmt.Sprintf("%s_errors.csv", fn)
}

// read the csv file line by line and write the output to a channel
// and if there is any parsing errors , write the error to the error channel (with parsed line if exists)
func readCSVFile(csvFile io.ReadCloser, errCh chan []string, readChannel chan []string, wg *sync.WaitGroup) {

	reader := csv.NewReader(csvFile)
	reader.Comma = ';'
	reader.FieldsPerRecord = 3

	defer wg.Done()
	defer close(readChannel)

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			errSlice := []string{"PARSING-ERROR", err.Error()}
			line = append(line, errSlice...)
			errCh <- line
			continue
		}
		readChannel <- line
	}
}

// read from a channel and process each line by calling the API
// and write any error occur during the process to the error channel
func processFile(readChannel chan []string, errCh chan []string, wg *sync.WaitGroup) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	recency := &updateRecency{}

	defer wg.Done()
	count := 0
	for line := range readChannel {
		count++
		recency.Document = line[0]
		recency.OrderID = line[1]
		recency.Date = line[2]

		recencyBytes, err := json.Marshal(recency)
		if err != nil {
			fmt.Println("marshalling body err :", err)
			errCh <- line
			continue
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(recencyBytes))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Access-Control-Realm", "internal")

		if err != nil {
			fmt.Println("Error creating request :", err)
			continue
		}
		response, err := client.Do(req)
		if err != nil {
			errSlice := []string{"POST-ERROR", err.Error()}
			line = append(line, errSlice...)
			errCh <- line
			continue
		}

		if response.StatusCode != 204 {
			// write to csv file errs
			responseBody, _ := ioutil.ReadAll(response.Body)
			errSlice := []string{response.Status, string(responseBody)}
			line = append(line, errSlice...)
			errCh <- line
		}

		response.Body.Close()
	}
	fmt.Printf("%d records were processed.\n", count)
	fmt.Println("Generating error report please wait.")
	close(errCh)
}

// generate the error report by reading from the error channel and writing
// the errors to a csv file
func generateErrorReport(errFile io.WriteCloser, errCh chan []string, wg *sync.WaitGroup) {
	writer := csv.NewWriter(errFile)
	// writing header
	header := []string{"LINE", "RESPONSE_CODE", "ERROR_MESSAGE"}
	writer.Write(header)
	defer wg.Done()

	for errLine := range errCh {
		writer.Write(errLine)
		if writer.Error() != nil {
			fmt.Println("Error writing to errors.csv file :", writer.Error())
		}
	}
	writer.Flush()
}

type updateRecency struct {
	Document string `json:"document,omitempty"`
	Date     string `json:"date,omitempty"`
	OrderID  string `json:"orderId,omitempty"`
}

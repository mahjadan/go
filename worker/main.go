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
	"sync"
	"time"
)

func main() {
	start := time.Now()
	csvFilename := "cr_optin_exemplo.csv"
	csvFile, err := os.Open(csvFilename)
	if err != nil {
		panic(fmt.Errorf("Error opening the file: ", err))
	}
	errFilename := csvFilename + "_error"
	errFile, err := os.Create(errFilename)
	if err != nil {
		panic(fmt.Errorf("Error while creating error file: ", err))
	}
	defer csvFile.Close()
	defer errFile.Close()

	run(csvFile, errFile)

	fmt.Printf("This script took %s\n", time.Since(start))

	fmt.Printf("Please check %s for errors.\n", errFilename)
}

// read csv file into a readChannel, then split readChannel into 5 channels and start 5 workers to read from each channel
//the program will finish when all workers are done.
// errors are handled by writing into errChannel and then reading from this channel and writing to csv file.
// only parsing errors are being handled for now,
// 1- TODO add more functionality to workers to parse the line into struct
// 2- TODO also handle the errors of parsing to struct in the same way mentioned above
// 3- TODO make the workers call the API and also handle errors.

func run(csvFile io.ReadCloser, errFile io.WriteCloser) {

	fmt.Println("Start reading csv file.")
	readChannel,errChannel := readCSVFile(csvFile)

	errWg := sync.WaitGroup{}
	errWg.Add(1)
	handleErrors(errChannel,errFile,&errWg)

	fmt.Println("the script is now processing your file.")

	const numberOfJobs = 5
	var jobs [numberOfJobs]chan Data

	// Distribute the job work across the numberOfJobs goroutines that all read from readChannel.
	for i := 0; i < numberOfJobs; i++ {
		jobName := fmt.Sprintf("job-%d", i+1)
		jobs[i] = job(jobName, readChannel)
	}

	workerWg := sync.WaitGroup{}
	workerWg.Add(len(jobs)) // len of array or numberOfJobs

	// Start the work on each job, reading from the JobChannel array.
	for i := 0; i < numberOfJobs; i++ {
		jobName := fmt.Sprintf("job-%d", i+1)
		go worker(jobName, jobs[i], &workerWg)
	}

	//wait for handling file errors
	errWg.Wait()
	//wait for workers to finish their work on each jobChannel.
	workerWg.Wait()
}

func handleErrors(errChannel chan []string, errFile io.WriteCloser,w *sync.WaitGroup) {
	writer := csv.NewWriter(errFile)
	// writing header
	header := []string{"LINE", "RESPONSE_CODE", "ERROR_MESSAGE"}
	writer.Write(header)

	go func() {

		defer w.Done()
		for errLine:= range errChannel {

			writer.Write(errLine)

			if writer.Error() != nil {
				fmt.Errorf("Error writing to errors.csv file :", writer.Error(),"\n")
			}
			//fmt.Println("got ERROR: ",errLine)
			writer.Flush()
		}

	}()
}

func worker(jobId string, jobChan <-chan Data, wg *sync.WaitGroup) {
	fmt.Println("working on - ", jobId)
	defer wg.Done()

	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	url:="localhost:8080/customer"

	for data := range jobChan {
		fmt.Println("job-", jobId, "processing line : ", data)
		dataBytes, err := json.Marshal(data)
		if err != nil {
			fmt.Println("marshalling body err :", err)
			errCh <- data
			continue
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(dataBytes))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Access-Control-Realm", "internal")

		if err != nil {
			fmt.Println("Error creating request :", err)
			continue
		}
		response, err := client.Do(req)
		if err != nil {
			errSlice := []string{"POST-ERROR", err.Error()}
			originalLine:=[]string{data.Email,data.Date}
			line := append(originalLine, errSlice...)
			errCh <- line
			continue
		}
		// success case - generate log or write to csv file.
		if response.StatusCode != 204 {// 200
			// write to csv file errs
			responseBody, _ := ioutil.ReadAll(response.Body)
			errSlice := []string{response.Status, string(responseBody)}
			originalLine:=[]string{data.Email,data.Date}
			line := append(originalLine, errSlice...)
			errCh <- line
		}

		response.Body.Close()
	}
}
//Starts new goroutine to fan-out to newly created channel.
func job(jobId string, readChannel <-chan []string) chan Data {

	fmt.Println("createJob - ", jobId)
	jobChannel := make(chan Data)
	go func() {
		count := 0
		for line := range readChannel {
			count++
			data:= Data{Email: line[0], Date: line[1]}
			jobChannel <- data
		}
		close(jobChannel)
		fmt.Printf("%d records were sent to channel[%s].\n", count, jobId)
	}()

	return jobChannel
}

type Data struct{
	Email string
	Date string
}
// job the csv file line by line and write the output to a channel
// and if there is any parsing errors , write the error to the error channel (with parsed line if exists)
func readCSVFile(csvFile io.ReadCloser) (chan []string, chan []string) {

	readCh:=make(chan []string)
	errCh:=make(chan []string)

	reader := csv.NewReader(csvFile)
	reader.Comma = ','
	reader.FieldsPerRecord = 2

	go func() {

		defer close(readCh)
		defer close(errCh)
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

			readCh <- line
		}

	}()

	return readCh, errCh
}


func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

//
//type updateRecency struct {
//	Document string
//	OrderID  string
//	Date     string
//}
//
//// job from a channel and process each line by calling the API
//// and write any error occur during the process to the error channel
//func processFile(readChannel chan []string, errCh chan []string, wg *sync.WaitGroup) {
//	client := http.Client{
//		Timeout: time.Duration(5 * time.Second),
//	}
//	recency := &updateRecency{}
//
//	defer wg.Done()
//	count := 0
//	for line := range readChannel {
//		count++
//		recency.Document = line[0]
//		recency.OrderID = line[1]
//		recency.Date = line[2]
//
//		recencyBytes, err := json.Marshal(recency)
//		if err != nil {
//			fmt.Println("marshalling body err :", err)
//			errCh <- line
//			continue
//		}
//
//		req, err := http.NewRequest("POST", "localhost:8080/customer", bytes.NewBuffer(recencyBytes))
//		req.Header.Set("Content-Type", "application/json")
//		req.Header.Set("X-Access-Control-Realm", "internal")
//
//		if err != nil {
//			fmt.Println("Error creating request :", err)
//			continue
//		}
//		response, err := client.Do(req)
//		if err != nil {
//			errSlice := []string{"POST-ERROR", err.Error()}
//			line = append(line, errSlice...)
//			errCh <- line
//			continue
//		}
//
//		if response.StatusCode != 204 {
//			// write to csv file errs
//			responseBody, _ := ioutil.ReadAll(response.Body)
//			errSlice := []string{response.Status, string(responseBody)}
//			line = append(line, errSlice...)
//			errCh <- line
//		}
//
//		response.Body.Close()
//	}
//	fmt.Printf("%d records were processed.\n", count)
//	fmt.Println("Generating error report please wait.")
//	close(errCh)
//}

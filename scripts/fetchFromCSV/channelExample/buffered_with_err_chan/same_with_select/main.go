package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {

	file, err := os.Open("../input-gen.csv")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	reader := csv.NewReader(file)
	readChan := make(chan []string)
	readChA := make(chan []string)
	readChB := make(chan []string)
	errChan := make(chan []string, 4)

	go readCSV(reader, readChan, errChan) // 20 items         // read csv and write to readChan, and any error write to errChan
	go fanout(readChan, readChA, readChB) // 10 items each // split the readChan to 2 channels ChA, ChB and close the channels.

	sem := make(chan int, 4) // channel to control how many can run at the same time. buffered chan. no need to close it.

	// waitgroup to know when all the go routines has finished their reading from ChA,ChB and also finished writing on the errChan
	// var wgGlobal sync.WaitGroup

	// wgGlobal.Add(2)
	aClosed := false
	bClosed := false

	go func() {
		var lineWG_a sync.WaitGroup
		var lineWG_b sync.WaitGroup

		for {
			if aClosed && bClosed {
				lineWG_a.Wait()
				lineWG_b.Wait()
				close(errChan)
				break
			}
			select {
			case data, ok := <-readChA:
				if !ok {
					fmt.Println("chA is closed.")
					aClosed = true
					readChA = nil
				} else {
					lineWG_a.Add(1)
					go doFetch(data, errChan, &lineWG_a, sem)
					// fmt.Println("chA: ", data)
					// go doFetchSimple(data, errChan)
				}
			case data, ok := <-readChB:
				if !ok {
					fmt.Println("chB is closed.")
					bClosed = true
					readChB = nil
				} else {
					lineWG_b.Add(1)
					go doFetch(data, errChan, &lineWG_b, sem)
					// go doFetchSimple(data, errChan)
					// fmt.Println("chB: ", data)
					// go doFetchSimple(data, errChan)
				}
			}
		}
	}()

loop:
	for {

		select {
		case data, ok := <-errChan:
			if !ok {
				fmt.Println("errChan is closed.")
				break loop
			}
			fmt.Println("err_recieved: ", data)
		}
	}

	// go readFromChannel(readChA, sem, errChan, &wgGlobal)
	// go readFromChannel(readChB, sem, errChan, &wgGlobal)

	// you need to make sure the writing is do in seperate go routine,
	// otherwise in the above 2 routine of readFromChannel, there is waitingGroup and waiting group .Wait() will
	// block , but we are still writing on errChan , and writing on the channel will block, thats why you need to start
	// the reading process at the same time in order to clear the channel and not block,
	// now how to close the channel on the receiving is by knowing the both the readingFromCh routines finished then
	// only u wil be able to close the channel. and this will help the writingErrors which is reading from
	// the errChan will not give deadlock because we closed the channel on the right time.
	// go writeErrors(errChan)

	// wgGlobal.Wait()
	// close(errChan)
}

func writeErrors(errCh chan []string) {
	file, _ := os.Create("errors.csv")
	writer := csv.NewWriter(file)

	for data := range errCh {
		fmt.Println("err: ", data)
		writer.Write(data)
		writer.Flush()
	}
}
func readFromChannel(chAorB chan []string, sem chan int, errCh chan []string, channelWG *sync.WaitGroup) {
	defer channelWG.Done()
	var lineWG sync.WaitGroup
	for data := range chAorB {
		lineWG.Add(1)
		// each doFetch takes 0.5 sec.
		// because the sem is 4 , it will allow only 4 to run at the same time.
		// processing 10 from the channel  10/4 2.5 sec total program with 20 items split by 2 ch.
		go doFetch(data, errCh, &lineWG, sem)
	}
	lineWG.Wait()
}

func doFetch(data []string, errCh chan []string, wg *sync.WaitGroup, sem chan int) {
	defer wg.Done()
	sem <- 1 //it will block here is the buffer size has reached.
	time.Sleep(500 * time.Millisecond)
	id, _ := strconv.Atoi(data[0])
	if id > 12 {
		errCh <- data
	}
	fmt.Printf("fetching: %v\n", data)
	<-sem // here is to clear it from the buffer.
}
func doFetchSimple(data []string, errCh chan []string) {
	// defer wg.Done()
	// sem <- 1 //it will block here is the buffer size has reached.
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("fetching: %v\n", data)
	id, _ := strconv.Atoi(data[0])
	if id > 7 {
		fmt.Println("writing to errCha.")
		errCh <- data
	}
	// <-sem // here is to clear it from the buffer.
}

func readCSV(reader *csv.Reader, ch chan<- []string, errCh chan<- []string) {
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			errCh <- line
		}

		ch <- line
	}
	defer close(ch)
}

func fanout(in <-chan []string, chA, chB chan []string) {
	aCount, bCount := 0, 0
	for data := range in {
		select {
		case chA <- data: //+ "-chA":
			aCount++
			fmt.Println("chA")
		case chB <- data: //+ "-chB":
			bCount++
			fmt.Println("chB")
		}
	}
	fmt.Printf("%d sent to chA and %d sent to ChB\n-", aCount, bCount)
	close(chA)
	close(chB)
	fmt.Println("testing", chA)
	var ar []string
	var x string
	ar = append(ar, x)
	// TODO:
	// TODO: SDFSDFSDHSH
}

// Test implementation.
type Code struct {
}

func (v *Code) Vest(s int) string {
	return ""
}

// VV implementation.
func (c *Code) VV(s string) string {
	return ""
}
func SX(s string) []string {
	return []string{}
}
func ss(s string) []string {

	return []string{}
}

// Codes .
type Codes struct {
	a int
}

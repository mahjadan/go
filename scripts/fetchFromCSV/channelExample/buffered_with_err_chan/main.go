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

	file, _ := os.Open("input-gen.csv")
	reader := csv.NewReader(file)
	readChan := make(chan []string)
	readChA := make(chan []string)
	readChB := make(chan []string)
	errChan := make(chan []string)

	go readCSV(reader, readChan, errChan) // 20 items         // read csv and write to readChan, and any error write to errChan
	go fanout(readChan, readChA, readChB) // 10 items each // split the readChan to 2 channels ChA, ChB and close the channels.

	sem := make(chan int, 4) // channel to control how many can run at the same time. buffered chan. no need to close it.

	// waitgroup to know when all the go routines has finished their reading from ChA,ChB and also finished writing on the errChan
	var wgGlobal sync.WaitGroup

	wgGlobal.Add(2)
	go readFromChannel(readChA, sem, errChan, &wgGlobal)
	go readFromChannel(readChB, sem, errChan, &wgGlobal)

	// you need to make sure the writing is do in seperate go routine,
	// otherwise in the above 2 routine of readFromChannel, there is waitingGroup and waiting group .Wait() will
	// block , but we are still writing on errChan , and writing on the channel will block, thats why you need to start
	// the reading process at the same time in order to clear the channel and not block,
	// now how to close the channel on the receiving is by knowing the both the readingFromCh routines finished then
	// only u wil be able to close the channel. and this will help the writingErrors which is reading from
	// the errChan will not give deadlock because we closed the channel on the right time.
	go writeErrors(errChan)

	wgGlobal.Wait()
	close(errChan)
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
func readFromChannel(readCh chan []string, sem chan int, errCh chan []string, channelWG *sync.WaitGroup) {
	defer channelWG.Done()
	var lineWG sync.WaitGroup
	for data := range readCh {
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
	fmt.Println(data)
	<-sem // here is to clear it from the buffer.
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
	for data := range in {
		select {
		case chA <- data: //+ "-chA":
			fmt.Println("chA")
		case chB <- data: //+ "-chB":
			fmt.Println("chB")
		}
	}
	close(chA)
	close(chB)
}

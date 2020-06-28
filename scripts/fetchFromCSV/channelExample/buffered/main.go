package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

func main() {

	// generate()
	file, _ := os.Open("input-gen.csv")
	reader := csv.NewReader(file)
	readChan := make(chan []string)
	readChA := make(chan []string)
	readChB := make(chan []string)

	go readCSV(reader, readChan)          // 20 items         // read csv and write to readChan
	go fanout(readChan, readChA, readChB) // 10 items each // split the readChan to 2 channels ChA, ChB and close the channels.

	sem := make(chan int, 4) // channel to control how many can run at the same time. buffered chan. no need to close it.

	var wgGlobal sync.WaitGroup // waitgroup to know when all the go routines has finished their work, doWork
	wgGlobal.Add(2)
	go readFromChannel(readChA, sem, &wgGlobal)
	go readFromChannel(readChB, sem, &wgGlobal)
	wgGlobal.Wait()
}

func readFromChannel(readCh chan []string, sem chan int, channelWG *sync.WaitGroup) {
	defer channelWG.Done()
	var lineWG sync.WaitGroup
	for data := range readCh {
		lineWG.Add(1)
		// each doFetch takes 0.5 sec.
		// because the sem is 4 , it will allow only 4 to run at the same time.
		// processing 10 from the channel  10/4 2.5 sec total program with 20 items split by 2 ch.
		go doFetch(data, &lineWG, sem)
	}
	lineWG.Wait()
}

func doFetch(data []string, wg *sync.WaitGroup, sem chan int) {
	defer wg.Done()
	sem <- 1 //it will block here is the buffer size has reached.
	time.Sleep(500 * time.Millisecond)
	fmt.Println(data)
	<-sem // here is to clear it from the buffer.
}

func readCSV(reader *csv.Reader, ch chan<- []string) {
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
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

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	in := make(chan string)
	err := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			in <- fmt.Sprintf("in-%d", i)
		}
		close(in)
	}()

	chA := make(chan string)
	chB := make(chan string)
	var wg sync.WaitGroup
	go fanout(in, chA, chB)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range chA {
			// fmt.Println("chA: ", data)
			time.Sleep(time.Second)
			err <- data
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range chB {
			// fmt.Println("chB: ", data)
			time.Sleep(time.Second)

			err <- data

		}
	}()

	go func() { // this will wait for bothe the channels to finish writing on err and then close it.
		wg.Wait()
		close(err)
	}()

	// go f(in, chA)
	// for data := range chA {
	// 	fmt.Println("chA-", data)
	// }

	for data := range err {
		fmt.Println("recieved: ", data)
	}

}

// func f(in <-chan string, chA chan string) {
// 	for data := range in {
// 		select {
// 		case chA <- data:
// 		}
// 	}
// 	close(chA)
// }

func fanout(in <-chan string, chA, chB chan string) {
	for data := range in {
		select {
		case chA <- data + "-chA":
		case chB <- data + "-chB":
		}
	}
	close(chA)
	close(chB)
}

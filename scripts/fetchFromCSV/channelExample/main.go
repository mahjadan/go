package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	asyncGenerator(10)
}

func send(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	// time.Sleep(time.Duration(id*100) * time.Millisecond)
	time.Sleep(1 * time.Second)
	ch <- id
	fmt.Println("done writing:", id)
}

func asyncGenerator(max int) {
	ch := make(chan int, 3)
	ch2 := make(chan int)
	var wg2 sync.WaitGroup

	// runLoop(ch, max)
	var wg sync.WaitGroup
	for i := 1; i < max; i++ {
		wg.Add(1)
		fmt.Printf("firing %d\n", i)
		go send(i, ch, &wg)
	}
	wg2.Add(1)
	go func() {
		wg2.Done()
		ch2 <- 1
	}()
	go func() {
		wg2.Wait()
		fmt.Println("closing ch2")
		close(ch2)
	}()

	fmt.Println("finish loop.")
	go func() {
		wg.Wait()
		close(ch)
	}()

	fmt.Println("now starting to recieve.")
loop:
	for {

		select {
		case id := <-ch:
			if id == 0 {
				break loop
			}
			fmt.Println("id:", id)
		case ex := <-ch2:
			if ex != 0 {
				fmt.Println("finished....:", ex)
			}
			continue
		}
	}
	// for id := range ch {
	// fmt.Println("id: ", id)
	// }
}

func runLoop(ch chan int, max int) {
	var wg sync.WaitGroup
	for i := 0; i < max; i++ {
		wg.Add(1)
		fmt.Printf("firing %d\n", i)
		go send(i, ch, &wg)
	}

	fmt.Println("finish loop.")
	go func() {
		wg.Wait()
		close(ch)
	}()

}

// func AsyncHTTP(users []string) ([]string, error) {
// 	ch := make(chan string)
// 	var responses []string
// 	var user string
// 	var wg sync.WaitGroup
// 	for _, user = range users {
// 		wg.Add(1)
// 		go sendUser(user, ch, &wg)
// 	}

// 	// close the channel in the background
// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()
// 	// read from channel as they come in until its closed
// 	for res := range ch {
// 		responses = append(responses, res)
// 	}

// 	return responses, nil
// }
// func sendUser(user string, ch chan<- string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	resp, err := http.Get("URL/" + user)
// 	if err != nil {
// 		log.Println("err handle it")
// 	}
// 	defer resp.Body.Close()
// 	b, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println("err handle it")
// 	}
// 	ch <- string(b)
// }

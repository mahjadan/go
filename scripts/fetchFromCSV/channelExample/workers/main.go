package main

import "time"

func main() {

}

// think of sendind the line as struct and it has it errCh to write to
type Line struct {
	errCh chan string
	line  string
}

func Run() {
	in, out := make(chan *Work), make(chan *Work)
	maxWorkers := 10

	for i := 0; i < maxWorkers; i++ {
		go worker(in, out)
	}
	go sendLotsOfWork(in)
	receiveLotsOfResults(out)
}

type Work struct {
	x, y, z int
}

func worker(in <-chan *Work, out chan<- *Work) {
	// think of this "in" as line coming from csv, no, should be a list of line.
	// why dont we write on different channel and split the file to multiple channel, and
	// each worker will receive "in" and the work struct will hold the name of the channel need to write the errors.
	// or all they can write to the same err channel.
	for w := range in {
		w.z = w.y * w.x
		time.Sleep(500 * time.Millisecond)
		out <- w
	}
}

func sendLotsOfWork(in chan *Work) {}

func receiveLotsOfResults(out chan *Work) {}

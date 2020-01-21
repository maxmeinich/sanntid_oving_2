// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
)

// Control signals
const (
	GetNumber = iota
	Exit
)

func number_server(add_number <-chan int, control <-chan bool, number chan<- int) {
	var i = 0
	// This for-select pattern is one you will become familiar with if you're using go "correctly".
	for {
		select {
			// TODO: receive different messages and handle them correctly
			// You will at least need to update the number and handle control signals.
			case <- control:
				number <- i
			case tall := <-add_number:
				i = i + tall
		}
	}
}

func incrementing(add_number chan<-int, finished chan<- bool) {
	for j := 0; j<1000; j++ {
		add_number <- 1
		Println("Perry")
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func decrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j<1000; j++ {
		add_number <- -1
		Println("Nebbdyret")
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the required channels
	// Think about wether the receptions of the number should be unbuffered, or buffered with a fixed queue size.
	control := make(chan bool)
	number := make(chan int)
	add_number := make(chan int)
	// TODO: Spawn the required goroutines
	go number_server(add_number, control, number)
	go incrementing(add_number, control)
	go decrementing(add_number, control)
	// TODO: block on finished from both "worker" goroutines
	var temp = 0
	for ; temp < 2;{
		select{
			case <-control:
				temp = temp + 1
		}
	}

	control<-true
	Println("The magic number is:", <-number)
}

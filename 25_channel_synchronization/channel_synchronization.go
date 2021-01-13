package main

import (
	"fmt"
	"time"
)

// we can use channels to synchronize ececution across goroutines.

// Example: using a blocking receive to wait for a goroutine to finish
// this function will run in a goroutine
// the done channel will be used to notify another goroutine that this function's work is done
func worker(done chan bool) {
	fmt.Print("working... ")
	time.Sleep(time.Second)
	fmt.Println("done")

	// send a value to notify that we are done
	done <- true
}

func main() {
	// start a worker goroutine and give it a channel to notify on
	done := make(chan bool, 1)
	go worker(done)

	// blocking receive
	// block until we receive notification from the worker on the channel
	// if we removed this line, the program would exit before the worker even started
	<-done
}

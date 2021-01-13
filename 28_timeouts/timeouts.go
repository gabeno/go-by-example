package main

import (
	"fmt"
	"time"
)

/*
   Timeouts are important for programs that connect to external resources or that otherwise need
   to bound execution time. Implemented in Go using channels and select.
*/

func main() {
	// simulate an external call that returns its result on channel c1 after 2s
	// note that the channel is buffered so the send in the goroutine is non-blocking - a common
	// pattern to prevent goroutine leaks in case the channel is never read
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// await results and await value to be sent after the timeout of 1s
	// since select proceeds with the first receive that is ready, we will take the timeout case
	// if the operation takes more than allowed 1s
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// if we allow a longer timeout of 3s then the receive from c2 will succeed and we will print
	// out the results
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

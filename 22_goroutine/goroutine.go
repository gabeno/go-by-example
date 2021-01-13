package main

import (
	"fmt"
	"time"
)

// goroutine is a lightweight thread of execution

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// synchronous call
	f("direct")

	// invoke this function in a goroutine
	// executes concurrently with the calling* one
	go f("goroutine")

	// goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// two function calls running asynchronously in separate goroutines now
	// wait for them to finish by sleeping (more robustly use WaitGroup)
	time.Sleep(time.Second)
	fmt.Println("done")
}

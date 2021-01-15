package main

import (
	"fmt"
	"sync"
	"time"
)

// this function is run in every goroutine
// a WaitGroup must be passed to functions through pointers
func worker(id int, wg *sync.WaitGroup) {

	// on return, notify WaitGroup that we are done
	defer wg.Done()

	// simulate an expensive task
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// the WaitGroup is used to wait for all the goroutines launched here to finish
	var wg sync.WaitGroup

	// launch several goroutines and increment the WaitGroup counter for each
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// block until the WaitGroup counter goes to 0; all thw workers notified they are done
	wg.Wait()
}

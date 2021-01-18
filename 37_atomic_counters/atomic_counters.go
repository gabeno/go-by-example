package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
   The primary mechanism for managing state in Go is communication over channels. We saw this
   for example with worker pools. There are a few other options for managing state though.
   Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple
   goroutines.
*/
func main() {
	// use unsigned integer to represent our (always positive counter)
	var ops uint64

	// A WaitGroup helps us wait for all goroutines to finish their work
	var wg sync.WaitGroup

	// start 50 goroutines that each increment the counter exactly 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// to atomically increment the counter we use AddUint64, giving it a memory address of
		// our ops counter
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// wait for all goroutines to be Done
	wg.Wait()

	// It’s safe to access ops now because we know no other goroutine is writing to it. Reading
	// atomics safely while they are being updated is also possible, using functions like
	// atomic.LoadUint64.
	fmt.Println("ops", ops)
}

/*
   $ go run 37_atomic_counters/atomic_counters.go
   ops 50000

   We expect to get exactly 50,000 operations. Had we used the non-atomic ops++ to increment the
   counter, we’d likely get a different number, changing between runs, because the goroutines
   would interfere with each other. Moreover, we’d get data race failures when running with
   the -race flag.
*/

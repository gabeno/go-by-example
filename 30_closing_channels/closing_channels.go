package main

import "fmt"

/*
   Closing a channel indicates that no more value will be sent on it. this can be useful to
   communicate completion to the channel's receivers
*/

func main() {
	// example with a jobs channel to communicate work done from the main goroutine to a worker
	// goroutine. when we have no more jobs for the worker we will close the jobs channel

	jobs := make(chan int, 5)
	done := make(chan bool)

	// worker goroutine
	go func() {
		for {
			// special 2-value form of receive
			// more value will be false if jobs has been closed and all values in the channel have
			// been received. we use this to nofify on done when we have worked on all our jobs.
			j, more := <-jobs
			if more {
				fmt.Println("received job:", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// send 3 jobs to the worker over jobs channel then close it
	for j := 1; j < 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// await worker using synchronization approach
	<-done
}

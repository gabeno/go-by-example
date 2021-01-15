package main

import (
	"fmt"
	"time"
)

/*
   implementing worker pools using goroutines and channels
*/

func worker(id int, jobs <-chan int, results chan<- int) {
	// receive work on the jobs channel and send results on results channel
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5

	// in order to use our pool of workers, we need to send them work and collect their results
	// so we need to create these two channels
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// start up 3 workers, initially blocked as there are no jobs yet
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send 5 jobs then close that channel to indicate thats all the work we have
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// finally collect all results of the work. this also ensure the worker goroutines have finished
	// alternatively we can use WaitGroup
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

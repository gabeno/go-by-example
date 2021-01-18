package main

import (
	"fmt"
	"time"
)

// rate limiting: control resource utilization and maintain quality of service
// Go support rate limiting with goroutines, channels and tickers

func main() {
	// Example 1

	// limit handling of incoming requests through a channel
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// limiter channel will receive a value every 200ms. this is the regulator in our rate
	// limiting scheme
	limiter := time.Tick(200 * time.Millisecond)

	// by blocking on a receive from the limiter channel before serving each request, we limit
	// ourselves to 1 request every 200ms
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Example 2

	// we may want to allow short bursts of requests in our rate limiting scheme while preserving
	// the overall rate limit. We can accomplish this by buffering our limiter channel. This
	// burstyLimiter channel allows bursts of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)

	// fill up the channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// every 200ms we try to add  a new value to burstyLimiter, up to its limit of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// simluate 5 more incoming requests. the first 3 of these will benefit from the burst
	// capability of burstyLimiter
	burstyRequests := make(chan int, 5)
	for i := 0; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

// For example 1:   requests handled every 200ms
// For example 2:   first 3 requests served almost imeediately due to burstable rate limiting, then
//                  serve the remaining 2 with 200ms delay

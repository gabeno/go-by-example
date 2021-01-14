package main

import (
	"fmt"
	"time"
)

/*
   timers are for when you want to do something once in the future.
   tickers are for when you want to do something repeatedly at regular intervals
*/

func main() {
	// example of ticker that ticks periodically until we stop it

	// tickers use a similar mechanism to timers: a channel that is sent values
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// use select on the channel to await values as they arrive after every 500ms
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// tickers can be stopped like timers and once stopped they will not be able to receive any more
	// values.
	// stop our ticker after 1600ms
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

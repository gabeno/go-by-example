package main

import (
	"fmt"
	"time"
)

/*
   A means to execute code at some point in the future or repeatedly at some interval.
   Timers represent a single event in the future.
*/

func main() {
	// tell the timer how long you want to wait and it provides a channel that will be notified
	// at that time, eg 2s in this example
	timer1 := time.NewTimer(2 * time.Second)

	// blocks on the timer's channel C until it sends a value that the timer fired
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// if you just want to wait, you may use time.Sleep. one reason a timer is useful is that you
	// can cancel a timer before it fires:
	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// give timer2 enough time to fire, if it ever was going to, to show infact it stopped.
	time.Sleep(2 * time.Second)

	// the first timer will fire ~2s after we start the program, but the second should be stopped
	// before it has a chance to fire.
}

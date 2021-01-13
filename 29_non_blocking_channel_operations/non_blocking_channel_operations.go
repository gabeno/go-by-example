package main

import (
	"fmt"
)

/*
   Basic sends and receives are blocking. however we can use select with a default clause
   to implement non-blocking sends, receives and even non-blocking multi-way selects
*/

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// a non-blocking receive
	// if a value is available on messages then select will take the <-messages case with that
	// value, if not it will imeediately take the default case
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// a non-blocking send works similarly
	// msg can not be sent to the messages channel because the channel has no buffer and there is
	// no receiver, therefore the default case is selected.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// use multiple cases above the default clause to implement a multiway non-blocking select.
	// we attempt non-blocking receives on both messages and signals channels.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

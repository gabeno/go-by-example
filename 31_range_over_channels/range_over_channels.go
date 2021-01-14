package main

import "fmt"

/*
   use range to iterate over values received from a channel
*/

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// iterate over each element as it is received from queue. since the channel is closed, the
	// iteration terminates after receiving two elements.
	for elem := range queue {
		fmt.Println(elem)
	}

	// its possible to close a non-empty channel and still have the remaining values received
}

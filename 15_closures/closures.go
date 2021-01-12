package main

import "fmt"

// intSeq returns another function which is defined anonymously in the body of intSeq
// the returned function closes over the variable i to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// we call intSeq, assigning the result (a function) to nextInt
	// this function value captures its own i value, which will be updated each time we call nextInt
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// state is unique to that particular function
	// create and test a new one
	nextInts := intSeq()
	fmt.Println(nextInts())
}

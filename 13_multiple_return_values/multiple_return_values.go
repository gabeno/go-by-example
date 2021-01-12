package main

import "fmt"

// multiple return values are used in idiomatic Go for example to return both result and error
// values from a function.

// (int, int) signature shows that the function returns 2 ints
func vals() (int, int) {
	return 3, 7
}

func main() {
	// use the two return values
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// ignore one of the return values
	_, c := vals()
	fmt.Println(c)
}

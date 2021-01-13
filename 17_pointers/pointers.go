package main

import "fmt"

// pointers allow you to pass references to values and records within your program

// zeroval has int parameter so arguments will be passed to it by value
// zeroval will get a copy of ival distinct from the one in the calling function
func zeroval(ival int) {
	ival = 0
}

// zeroptr has *int parameter meaning it takes an int pointer
// the *iptr dereferences the pointer from its memory address to the current value at that address
// assigning a value to a dereferenced pointer changes the value at the referenced address
func zeroptr(iptr *int) {
	*iptr = 0
}

// zeroval does not change the i in main
// zeroptr does as it has reference to the memory address of that variable
func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// printing the pointer
	fmt.Println("pointer:", &i)
}

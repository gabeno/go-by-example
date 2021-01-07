package main

import "fmt"

// Go has only one looping construct: for
func main() {
	// single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// intialize:condition:after
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without a condition loops repeatedly until your break out of the loop or
	// return from the enclosing function
	for {
		fmt.Println("continuous loop")
		break
	}

	// you can also continue to the next iteration of the loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

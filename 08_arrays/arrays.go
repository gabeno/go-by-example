package main

import "fmt"

func main() {
	// an array holding 5 integers
	// by default an array is zero valued
	var a [5]int
	fmt.Println("empty: ", a)

	// set a value at index
	a[4] = 100
	fmt.Println("set value at index 4: ", a)
	// get a value at index
	fmt.Println("set value at index 4: ", a[4])

	// get array length
	fmt.Println("array length: ", len(a))

	// declare and initialize an array in one line
	d := [5]int{1, 2, 3, 5, 4}
	fmt.Println("declare: ", d)

	// we can create a 2d array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D array: ", twoD)
}

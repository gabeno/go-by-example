package main

import "fmt"

func main() {
	// var keyword declares variables
	var a = "Initial"
	fmt.Println(a)

	// one can declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// go will infer types of initialized variables
	var d = true
	fmt.Println(d)

	// variables declared without corresponding initialization are zero-valued
	var e int
	fmt.Println(e)

	// the := syntax is shorthand for declaring and initializing a variable
	// var f string = "apple"
	f := "apple"
	fmt.Println(f)
}

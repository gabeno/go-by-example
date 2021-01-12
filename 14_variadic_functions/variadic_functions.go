package main

import "fmt"

// variadic functions accept any number of arguments
// they can be called with any number of trailing arguments

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// varidadic functions called in the usual way with individual arguments
	sum(1, 2)
	sum(1, 2, 3)

	// if you have multiple args in a slice, apply them to a variadic function using
	// func(slice...)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

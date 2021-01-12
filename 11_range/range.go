package main

import "fmt"

// range iterates over elements in a variety of data structures

func main() {
	// use range to sum numbers in a slice
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range on arrays and slices provides both index and value for each entry
	for idx, num := range nums {
		fmt.Println("index, num:", idx, num)
	}

	// range on map iterates over key/value pairs
	fruits := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range fruits {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can also iterate over the keys of a map
	for k := range fruits {
		fmt.Println("key:", k)
	}

	// range on strings iterates over unicode code points
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

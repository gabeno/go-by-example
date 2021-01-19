package main

import (
	"fmt"
	"sort"
)

// sorting for builtins

func main() {

	// sorting is specific to builtin type
	// for string it sorts in-place
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// sorting ints
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	// use sort to check is a slice is already in sorted order
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)

}

/*
   $ go run 40_sort/sort.go
   Strings: [a b c]
   Ints: [2 4 7]
   Sorted: true
*/

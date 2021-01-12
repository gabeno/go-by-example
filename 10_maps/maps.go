package main

import "fmt"

// maps are Go's associative data type
// also called hashes or dict in other langauges

func main() {
	// create using make
	// make(map[key-type]val-type)
	m := make(map[string]int)

	// set key:value pairs
	m["k1"] = 1
	m["k2"] = 2

	// printing a map will show all its key:value pairs
	fmt.Println("map: ", m)

	// get a value for a key
	v := m["k1"]
	fmt.Println("v: ", v)

	// builtin len returns the number of key/value pairs
	fmt.Println("len :", len(m))

	// builtin delete removes a key/value pair from a map
	delete(m, "k2")
	fmt.Println("map: ", m)

	// optional second return value when getting a value from a map indicates if the key was present
	// in the map. this can be used to disambiguate between missing keys and keys with zero values
	// like 0 or ""
	_, present := m["k2"]
	fmt.Println("present:", present)

	// declare and initialize a map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n: ", n)
}

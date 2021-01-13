package main

import "fmt"

// structs are typed collections of fields used for grouping data together to form records
type person struct {
	name string
	age  int
}

// construct a new person struct with the given name
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	// you may safely return a pointer to a local variable as local variable will survive
	// the scope of the function
	return &p
}

func main() {
	// create a new struct
	fmt.Println(person{"Bob", 20})

	// create with named fields
	fmt.Println(person{"Alice", 30})

	// omitted fields are zero valued
	fmt.Println(person{name: "Fred"})

	// an & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// its idiomatic to encapsulate new struct creations in constructor functions
	fmt.Println(newPerson("Jon"))

	// accessing struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// you can also use dots with struct pointers (the pointers are automatically dereferenced)
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}

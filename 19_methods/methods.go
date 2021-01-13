package main

import "fmt"

type rect struct {
	width, height int
}

// methods can be defined for either pointer or value receiver types.
// this method has a receiver type of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// this method has a value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perimeter: ", r.perim())

	// go auto handles conversions between values and pointers for method calls
	// you may want to use pointer receiver type to avoid copying on method calls or to allow the
	// method to mutate receiving struct
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perim())
}

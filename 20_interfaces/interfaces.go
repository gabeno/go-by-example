package main

import (
	"fmt"
	"math"
)

// interfaces are named collections of method signatures

// interface for geometric shapes
type geometry interface {
	area() float64
	perim() float64
}

// lets implement this interface on rect and circle types
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// implement interface on rect type
// ie implement all methods in the interface
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// implement interface on circle type
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// if a variable has an interface type then we can call methods that are in the named interface.
// Here is a generic measure function taking advantage of this to work on any geometry.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// circle and rect struct types both implement geometry interface so we can use instances of these
// structs as arguments to measure
func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

package main

import (
	"fmt"
	"math"
)

// GoLang supports constants of character, string, boolean and numeric values

// const statement can appear anywhere a var statement can
const s string = "some constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	// constant expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// a numeric constant has no type until it is given one
	fmt.Println(int64(d))

	// a number can be given type by using it in a context that requires one
	// eg math.Sin expects a float64 type
	fmt.Println(math.Sin(n))
}

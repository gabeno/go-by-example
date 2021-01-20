package main

import (
	"fmt"
	"strconv"
)

func main() {
	// With ParseFloat, this 64 tells how many bits of precision to parse.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// For ParseInt, the 0 means infer the base from the string. 64 requires that the result fit
	// in 64 bits.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt will recognize hex-formatted numbers.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// A ParseUint is also available.
	u, _ := strconv.ParseInt("789", 0, 64)
	fmt.Println(u)

	// Atoi is a convenience function for basic base-10 int parsing.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parse functions return an error on bad input.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}

/*
   Test:
   ----

   $ go run 54_number_parsing/number_parsing.go
   1.234
   123
   456
   789
   135
   strconv.Atoi: parsing "wat": invalid syntax
*/
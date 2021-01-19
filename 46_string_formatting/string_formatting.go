package main

import (
	"fmt"
	"os"
)

/*
   Go offers excellent support for string formatting in the printf tradition. Here are some
   examples of common string formatting tasks.
*/

type point struct {
	x, y int
}

func main() {
	// Go offers several printing “verbs” designed to format general Go values. For example, this
	// prints an instance of our point struct.
	p := point{1, 2}
	fmt.Printf("struct instance: %v\n", p)

	// If the value is a struct, the %+v variant will include the struct’s field names.
	fmt.Printf("struct field names: %+v\n", p)

	// The %#v variant prints a Go syntax representation of the value, i.e. the source code snippet
	// that would produce that value.
	fmt.Printf("struct syntax representation: %#v\n", p)

	//To print the type of a value, use %T.
	fmt.Printf("type of a value: %T\n", p)

	//Formatting booleans is straight-forward.
	fmt.Printf("booleans: %t\n", true)

	// There are many options for formatting integers. Use %d for standard, base-10 formatting.
	fmt.Printf("integers: %d\n", 123)

	// This prints a binary representation.
	fmt.Printf("binary representation: %b\n", 14)

	// This prints the character corresponding to the given integer.
	fmt.Printf("character from integer: %c\n", 33)

	// %x provides hex encoding.
	fmt.Printf("hex encoding: %x\n", 456)

	// There are also several formatting options for floats. For basic decimal formatting use %f.
	fmt.Printf("floats: %f\n", 78.9)

	//%e and %E format the float in (slightly different versions of) scientific notation.
	fmt.Printf("scientific 1: %e\n", 123400000.0)
	fmt.Printf("scientific 2: %E\n", 123400000.0)

	// For basic string printing use %s.
	fmt.Printf("string basic: %s\n", "\"string\"")

	// To double-quote strings as in Go source, use %q.
	fmt.Printf("double quote strings: %q\n", "\"string\"")

	// As with integers seen earlier, %x renders the string in base-16, with two output characters
	// per byte of input.
	fmt.Printf("%x\n", "hex this")

	//To print a representation of a pointer, use %p.
	fmt.Printf("pointer: %p\n", &p)

	// When formatting numbers you will often want to control the width and precision of the
	// resulting figure. To specify the width of an integer, use a number after the % in the verb.
	// By default the result will be right-justified and padded with spaces.
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// You can also specify the width of printed floats, though usually you’ll also want to
	// restrict the decimal precision at the same time with the width.precision syntax.
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// To left-justify, use the - flag.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// You may also want to control width when formatting strings, especially to ensure that they
	// align in table-like output. For basic right-justified width.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// To left-justify use the - flag as with numbers.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// So far we’ve seen Printf, which prints the formatted string to os.Stdout. Sprintf formats
	// and returns a string without printing it anywhere.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// You can format+print to io.Writers other than os.Stdout using Fprintf.
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}

/*
   Test:
   ----
   $ go run 46_string_formatting/string_formatting.go
   struct instance: {1 2}
   struct field names: {x:1 y:2}
   struct syntax representation: main.point{x:1, y:2}
   type of a value: main.point
   booleans: true
   integers: 123
   binary representation: 1110
   character from integer: !
   hex encoding: 1c8
   floats: 78.900000
   scientific 1: 1.234000e+08
   scientific 2: 1.234000E+08
   string basic: "string"
   double quote strings: "\"string\""
   6865782074686973
   pointer: 0xc00001a100
   |    12|   345|
   |  1.20|  3.45|
   |1.20  |3.45  |
   |   foo|     b|
   |foo   |b     |
   a string
   an error
*/

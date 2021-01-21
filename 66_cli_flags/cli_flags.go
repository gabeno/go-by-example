package main

import (
	"flag"
	"fmt"
)

/*
   Command-line flags are a common way to specify options for command-line programs. For example,
   in wc -l the -l is a command-line flag.

   Go provides a flag package supporting basic command-line flag parsing. We’ll use this package
   to implement our example command-line program.
*/

func main() {
	// Basic flag declarations are available for string, integer, and boolean options. Here we
	// declare a string flag word with a default value "foo" and a short description. This
	// flag.String function returns a string pointer (not a string value); we’ll see how to use
	// this pointer below.
	wordPtr := flag.String("word", "foo", "a string")

	// other types
	numPtr := flag.Int("num", 42, "a number")
	boolPtr := flag.Bool("fork", false, "a bool")

	// It’s also possible to declare an option that uses an existing var declared elsewhere in the
	// program. Note that we need to pass in a pointer to the flag declaration function.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Once all flags are declared, call flag.Parse() to execute the command-line parsing.
	flag.Parse()

	// Here we’ll just dump out the parsed options and any trailing positional arguments. Note that
	// we need to dereference the pointers with e.g. *wordPtr to get the actual option values.
	fmt.Println("word:", *wordPtr)
	fmt.Println("num:", *numPtr)
	fmt.Println("bool:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

/*
   Test:
   ----

   $ cd 66_cli_flags/
   $ go build cli_flags.go
   $ ./cli_flags -word=opt -num=7 -fork -svar=flag
   word: opt
   num: 7
   bool: true
   svar: flag
   tail: []
   $ ./cli_flags -word=opt
   word: opt
   num: 42
   bool: false
   svar: bar
   tail: []
   $ ./cli_flags -word=opt a1 a2 a3
   word: opt
   num: 42
   bool: false
   svar: bar
   tail: [a1 a2 a3]
   $ ./cli_flags -word=opt a1 a2 a3 -num=7 // -> Note that the flag package requires all flags to appear before positional arguments (otherwise the flags will be interpreted as positional arguments).
   word: opt
   num: 42
   bool: false
   svar: bar
   tail: [a1 a2 a3 -num=7]
   $ ./cli_flags -h
   Usage of ./cli_flags:
     -fork
       	a bool
     -num int
       	a number (default 42)
     -svar string
       	a string var (default "bar")
     -word string
       	a string (default "foo")
   $ ./cli_flags -wat
   flag provided but not defined: -wat
   Usage of ./cli_flags:
     -fork
       	a bool
     -num int
       	a number (default 42)
     -svar string
       	a string var (default "bar")
     -word string
       	a string (default "foo")
*/

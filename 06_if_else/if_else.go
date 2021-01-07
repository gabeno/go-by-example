package main

import "fmt"

// you do not need parentheses around conditions in GoLang but braces are required
// there is no ternary if in GoLang

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// you can have an if statement w/o an else
	if 8%4 == 0 {
		fmt.Println("8 is divisible")
	}

	// a statement can precede conditionals
	// any variables declared in this statement are available in all branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

package main

import (
	"errors"
	"fmt"
)

// errors in Go are passed as an explicit return value
// This approach makes it easy to see which functions returns errors and to handle
// them using the same language constructs employed for any other non-error tasks

// by convention errors are the last return value and have type error which is a builtin interface
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message
		return -1, errors.New("can't work with 42")
	}

	// a nil value in the error position indicates there was no error
	return arg + 3, nil
}

// using custom types as errors by implementing Error() method
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// use &argError syntax to build a new struct
		return -1, &argError{arg, "can't work with 42"}
	}

	return arg + 3, nil
}

func main() {
	// test builtin error returning function
	// inline error check is a common idiom in Go
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	// test custom error returning function
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// if you want to programatically use the data in a custom error, you will need to get
	// the error as an instance of the custom error type via type assertion
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

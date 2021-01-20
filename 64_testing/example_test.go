package main

import (
	"fmt"
	"testing"
)

/*
   Unit testing is an important part of writing principled Go programs. The testing package
   provides the tools we need to write unit tests and the go test command runs tests.

   For the sake of demonstration, this code is in package main, but it could be any package.
   Testing code typically lives in the same package as the code it tests.
*/

// We’ll be testing this simple implementation of an integer minimum. Typically, the code we’re
// testing would be in a source file named something like intutils.go, and the test file for it
// would then be named intutils_test.go.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// A test is created by writing a function with a name beginning with Test.
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Error* will report test failures but continue executing the test. t.Fatal* will report
		// test failures and stop the test immediately.
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

//Writing tests can be repetitive, so it’s idiomatic to use a table-driven style, where test inputs
// and expected outputs are listed in a table and a single loop walks over them and performs the
// test logic.
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d, %d", tt.a, tt.b)
		// t.Run enables running “subtests”, one for each table entry. These are shown separately
		// when executing go test -v.
		t.Run(testName, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

/*
   Test:
   ----
   $ cd 64_testing/
   $ go test -v
   === RUN   TestIntMinBasic
   --- PASS: TestIntMinBasic (0.00s)
   === RUN   TestIntMinTableDriven
   === RUN   TestIntMinTableDriven/0,_1
   === RUN   TestIntMinTableDriven/1,_0
   === RUN   TestIntMinTableDriven/2,_-2
   === RUN   TestIntMinTableDriven/0,_-1
   === RUN   TestIntMinTableDriven/-1,_0
   --- PASS: TestIntMinTableDriven (0.00s)
       --- PASS: TestIntMinTableDriven/0,_1 (0.00s)
       --- PASS: TestIntMinTableDriven/1,_0 (0.00s)
       --- PASS: TestIntMinTableDriven/2,_-2 (0.00s)
       --- PASS: TestIntMinTableDriven/0,_-1 (0.00s)
       --- PASS: TestIntMinTableDriven/-1,_0 (0.00s)
   PASS
   ok  	_/home/gm/Workbench/go-by-example/64_testing	0.001s
*/

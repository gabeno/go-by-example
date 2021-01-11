package main

import "fmt"

func main() {
	// slice is a more powerful interface to sequences in GO than arrays

	// intialize and declare a slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("declared in one line ", t)

	// empty slice with non-zero length
	s := make([]string, 3)

	// get and set just like in arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get :", s[2])

	// returns length as expected
	fmt.Println("len: ", len(s))

	// append
	// we need to accept a return value from append as we may get a new slice value
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("appended: ", s)

	// slices can be copied
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	// slices support slice operator
	// syntax slice[low:high] excluding upperbound
	l := s[2:5]
	fmt.Println("s[2:5]: ", l)

	l = s[:5]
	fmt.Println("s[:5]: ", l)

	l = s[2:]
	fmt.Println("s[2:]: ", l)

	// slices can be composed into multidimensional data structures
	// the legth of inner slices can vary, unlike with multidimensional arrays
	twoD := make([][]int, 3)
	fmt.Println("init: ", twoD)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("init: ", twoD)
}

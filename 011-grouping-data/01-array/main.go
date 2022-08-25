package main

import "fmt"

func main() {
	// array has fixed size
	var x [5]int

	// assign value
	x[3] = 42
	fmt.Println(x)

	// size is a part of array. y is of different type as x
	var y [6]int
	fmt.Println(y)

	// ***: array is more considered as the building block of slice.
	//		use slice instead
}

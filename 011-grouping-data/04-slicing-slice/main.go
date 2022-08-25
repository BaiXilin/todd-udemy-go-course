package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}

	// slicing a slice using : operator
	// 1. all elements
	fmt.Println(x[:])

	// 2. start from
	y := x[1:]
	fmt.Println(y)

	// 3. end at
	y = x[:3]
	fmt.Println(y)

	// 4. start from and end at
	y = x[1:3] // [a:b], include a, exclude b
	fmt.Println(y)
}

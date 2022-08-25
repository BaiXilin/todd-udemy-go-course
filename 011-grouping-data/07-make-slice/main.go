package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	y := make([]int, 5, 5)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	// appending to a slice when len == cap (all slots taken), the cap will grow
	y = append(y, 6)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))
}

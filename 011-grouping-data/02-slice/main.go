package main

import "fmt"

func main() {
	// Composite literal:
	// x := type{value};

	// slice allows to group together data of the same type
	x := []int{4, 5, 6, 7, 42}
	fmt.Println(x)
}

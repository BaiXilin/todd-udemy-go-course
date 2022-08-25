package main

import "fmt"

func main() {
	// append values to a slice
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6, 7)
	fmt.Println(x)

	// append slice to a slice
	y := []int{42, 43, 44, 45}
	x = append(x, y...) // ... takes all elements in y out
	fmt.Println(x)
}

package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("originla slice:", x)

	// suppose we want to delete element on index 3 and 4
	x = append(x[:3], x[5:]...)
	fmt.Println("after delete:", x)
}

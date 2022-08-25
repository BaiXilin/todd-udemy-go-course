package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}

	// properties of slice
	fmt.Println("length =", len(x))
	fmt.Println("capacity =", cap(x))

	// ranging all values in the slice
	for i := range x { // simplicication of i, _ := range x
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// alter
	for index, value := range x {
		fmt.Printf("x[%d] = %d\n", index, value)
	}
}

package main

import "fmt"

func main() {
	total := sum(2, 4, 6, 8, 42)
	fmt.Println("The sum is", total)

	// when passed in a single argument
	total = sum(1)
	fmt.Println("The sum is", total)
	
	// when passed in no argument
	total = sum()
	fmt.Println("The sum is", total)
}

// variadic function is used when the number of arguments is only known until run-time
func sum(x ...int) int {
	// find out what is passed in and the type
	fmt.Println("The argument:", x)
	fmt.Printf("The type of argument: %T\n", x)

	total := 0
	for _, value := range x {
		total += value
	}
	return total
}

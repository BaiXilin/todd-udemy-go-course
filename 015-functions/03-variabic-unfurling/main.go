package main

import "fmt"

func main() {
	xi := []int{1,2,3,4,5,6,7,8}
	sum := calc_sum(xi...)	// take out the ints in the slice and supply them to the variadic funciton
	fmt.Println("sum =", sum)
}

func calc_sum(x ...int) int {
	sum := 0
	for _, i := range x {
		sum += i
	}
	return sum
}

// Note: when a slice is passed as an argument, it is modifiable by the function

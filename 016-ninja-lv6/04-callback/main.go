// A “callback” is when we pass a func into a func as an argument. For this exercise,
// ● pass a func into a func as an argument

package main

import (
	"fmt"
)

func main() {
	find_total := func(xi ...int) int {
		total := 0
		for _, val := range xi {
			total += val
		}
		return total
	}

	series := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	print_sum(find_total, series...)
}

func print_sum(summation func(xi ...int) int, series ...int) {
	sum := summation(series...)
	fmt.Println("The sum of the series is", sum)
}

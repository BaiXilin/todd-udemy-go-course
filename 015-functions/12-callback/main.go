// callback is passing functions as arguments to functions
package main

import "fmt"

func main() {
	int_slice := []int{1,2,3,4,5,6,7,8,9}
	
	all_total := sum(int_slice...)
	fmt.Println("The sum of all integers is", all_total)
	
	// sum is a function. passed as an argument here
	even_total := sum_even(sum, int_slice...)
	fmt.Println("The sum of all even integers is", even_total)

	sum_total := sum_odd(sum, int_slice...)
	fmt.Println("The sum of all odd integer is", sum_total)
}

func sum(i ...int) int {
	total := 0
	for _, value := range i {
		total += value
	}
	return total
}

func sum_even(sum_func func(i ...int) int, xi ... int) int {
	even_i := []int{}
	for _, val := range xi {
		if val % 2 == 0 {
			even_i = append(even_i, val)
		}
	}
	// callback to sum()
	return sum_func(even_i...)
}

func sum_odd(sum_func func(i ...int) int, xi ...int) int {
	odd_i := []int{}
	for _, val := range xi {
		if val % 2 != 0 {
			odd_i = append(odd_i, val)
		}
	}
	// callback to sum()
	return sum_func(odd_i...)
}

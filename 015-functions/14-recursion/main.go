// calculating factorial using recursion
package main

import "fmt"

func main() {
	x := 10
	fact := factorial(x)
	fmt.Println("The factorial of", x, "is", fact)
}

func factorial(x int) int {
	if x == 1 {
		return 1
	}
	return x * factorial(x - 1)
}


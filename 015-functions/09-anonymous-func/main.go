package main

import "fmt"

func main() {
	// anonymous function has no identifier
	func() {
		fmt.Println("Inline function speaking")
	}()

	func(x int) {
		fmt.Println("The meaning of life is", x)
	}(42)

	// anonymous function can also have return value
	y := func(x int) string {
		return fmt.Sprintf("The meaning of life is %d.\n", x)
	}(42)
	fmt.Println(y)
}

package main

import "fmt"

func main() {
	defer foo()	// foo will be executed after bar(), when main() comes to end
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

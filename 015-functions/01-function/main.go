package main

import "fmt"

func main() {
	foo()

	boo("James")

	hello := woo("Moneypenny")
	fmt.Println(hello)

	helloworld, greeted := greeting("Miss")
	fmt.Printf("%s. Are you greeted? %t\n", helloworld, greeted) 
}

// function format:
// func (r receiver) identifier(parameters) (return types) { ... }

// function with no parameters and no return values
func foo() {
	fmt.Println("Hello from foo")
}

// function with parameters and no return values
func boo(s string) {
	fmt.Println("Hello from boo,", s)
}

// function with parameters and one return value
func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

// function with parameters and multiple return values
func greeting(s string) (string, bool) {
	greet := true
	str := woo(s)
	return str, greet
}

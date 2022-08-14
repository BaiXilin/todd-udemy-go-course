package main

import "fmt"

func main() {
	var a = 42
	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 42:
		fmt.Println("The answer to life and the universe")
	default:
		fmt.Println("The answer to life and the universe is NULL")
	}
}

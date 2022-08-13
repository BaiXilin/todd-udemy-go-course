// Write a program that prints a number in decimal, binary, and hex
package main

import "fmt"

func main() {
	var x int = 42
	fmt.Printf("BIN\t\tDEC\tHEX\n")
	fmt.Printf("%#b\t%d\t%#x\n", x, x, x)
}

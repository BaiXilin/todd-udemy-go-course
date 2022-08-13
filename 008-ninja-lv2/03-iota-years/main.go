// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
package main

import "fmt"

const (
	TwentyTwo   = 2022 + iota
	TwentyThree = 2022 + iota
	TwentyFour  = 2022 + iota
	TwentyFive  = 2022 + iota
)

func main() {
	fmt.Println("This year is", TwentyTwo)
	fmt.Println("Next year is", TwentyThree)
	fmt.Println("Next year is", TwentyFour)
	fmt.Println("Next year is", TwentyFive)
}

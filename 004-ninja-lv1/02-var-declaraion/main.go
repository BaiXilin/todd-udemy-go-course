// 1. Use var to DECLARE three VARIABLES. The variables should have package level
//    scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the
//    variables and make sure the variables are of the following TYPE (meaning they can
//    store VALUES of that TYPE).
// 		a. identifier “x” type int
// 		b. identifier “y” type string
// 		c. identifier “z” type bool
// 2. in func main
// 		a. print out the values for each identifier
// 		b. The compiler assigned values to the variables. What are these values called?

// Using the code from the previous exercise,
// 1. At the package level scope, assign the following values to the three variables
// 		a. for x assign 42
// 		b. for y assign “James Bond”
// 		c. for z assign true
// 2. in func main
// 		a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the
// 		   returned value of TYPE string using the short declaration operator to a
// 		   VARIABLE with the IDENTIFIER “s”
// 		b. print out the value stored by variable “s”

package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z) // zero values

	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprintf("%v, %v, %v\n", x, y, z)
	fmt.Print(s)
}

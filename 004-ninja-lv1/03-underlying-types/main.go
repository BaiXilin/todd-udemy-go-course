// 1. Create your own type. Have the underlying type be an int.
// 2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
// 3. in func main
// 		a. print out the value of the variable “x”
// 		b. print out the type of the variable “x”
// 		c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
// 		d. print out the value of the variable “x”

// 1. at the package level scope, using the “var” keyword, create a VARIABLE with the
//    IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
// 2. in func main
// 		a. this should already be done
// 			i. print out the value of the variable “x”
// 			ii. print out the type of the variable “x”
// 			iii. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
// 			iv. print out the value of the variable “x”
// 		b. now do this
// 			i. now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
// 				1. then use the “=” operator to ASSIGN that value to “y”
// 				2. print out the value stored in “y”
// 				3. print out the type of “y”

package main

import "fmt"

type electro int

var x electro
var y int

func main() {
	fmt.Printf("Variable x has value %v of type %T\n", x, x)
	x = 42
	fmt.Println("x now has value", x)

	y = int(x)
	fmt.Printf("Variable y has value %v of type %T\n", y, y)
}

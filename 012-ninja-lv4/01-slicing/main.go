// Create a slice of 10 ints, and use SLICING to create the following new slices
// which are then printed:
// ● [42 43 44 45 46]
// ● [47 48 49 50 51]
// ● [44 45 46 47 48]
// ● [43 44 45 49 50]

package main

import "fmt"

func main() {
	s := make([]int, 10)

	for i := 0; i < 10; i++ {
		s[i] = i + 42
	}
	fmt.Println("Original:", s)

	slicing_1 := s[:5]
	fmt.Println("Slicing 1:", slicing_1)

	slicing_2 := s[5:]
	fmt.Println("Slicing 2:", slicing_2)

	slicing_3 := s[2:7]
	fmt.Println("Slicing 3:", slicing_3)

	slicing_4 := append(s[1:4], s[7:9]...)
	fmt.Println("Slicing 4:", slicing_4)
}

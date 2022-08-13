package main

import "fmt"

const (
	Mon = iota
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
	NumberofDaysInWeek
)

// iota is resetted after the next const keyword
const (
	a = iota
	b
)

func main() {
	fmt.Println(Mon, Tue, Wed, Fri, Sat, Sun)
	fmt.Printf("The number of days in a week is %d\n", NumberofDaysInWeek)

	fmt.Println(a, b)
}

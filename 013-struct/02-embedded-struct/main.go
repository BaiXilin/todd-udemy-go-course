package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person      // a secret agent is a person, so embed person struct here
	ltk    bool // licsence to kill
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	fmt.Printf("sa1 is of type %T\n", sa1)

	// fields of the embedded type are promoted up
	fmt.Printf("%s %s is %d years old, he has license to kill? %t\n", sa1.first, sa1.last, sa1.age, sa1.ltk)
	// or include field name to access embedded values to avoid collision
	fmt.Printf("%s %s is %d years old, he has license to kill? %t\n", sa1.person.first, sa1.person.last, sa1.person.age, sa1.ltk)
}

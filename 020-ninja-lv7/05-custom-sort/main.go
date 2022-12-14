package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (u user) String() string {
	return fmt.Sprintf("%s %s is %d years old, they says %s\n", u.First, u.Last, u.Age, u.Sayings)
}

type byLast []user

func (bl byLast) Len() int {
	return len(bl)
}

func (bl byLast) Swap(i, j int) {
	bl[i], bl[j] = bl[j], bl[i]
}

func (bl byLast) Less(i, j int) bool {
	return bl[i].Last < bl[j].Last
}

type byAge []user

func (ba byAge) Len() int {
	return len(ba)
}

func (ba byAge) Swap(i, j int) {
	ba[i], ba[j] = ba[j], ba[i]
}

func (ba byAge) Less(i, j int) bool {
	return ba[i].Age < ba[j].Age
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	sort.Sort(byLast(users))
	fmt.Println("By last name:\n", users)

	sort.Sort(byAge(users))
	fmt.Println("By age:\n", users)
}

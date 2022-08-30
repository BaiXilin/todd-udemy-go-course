package main

import (
	"fmt"

	"github.com/BaiXilin/todd-udemy-go-course/tree/master/015-functions/08-interface-composition/contractor"
	"github.com/BaiXilin/todd-udemy-go-course/tree/master/015-functions/08-interface-composition/tools"
)

func main() {
	// Inventory of old boards to remove, and the new boards
	// that will replace them.
	boards := []tools.Board{
		// Rotted boards to be removed.
		{NailsDriven: 3},
		{NailsDriven: 1},
		{NailsDriven: 6},

		// Fresh boards to be fastened.
		{NailsNeeded: 6},
		{NailsNeeded: 9},
		{NailsNeeded: 4},
	}

	// Fill a toolbox
	tb := contractor.Toolbox{
		NailDriver: tools.Mallet{},
		NailPuller: tools.Crowbar{},
		Nails:      10,
	}

	displayStates(&tb, boards)

	// Hire a contractor to work
	var c contractor.Contractor
	c.ProcessBoards(&tb, &tb.Nails, boards)

	displayStates(&tb, boards)
}

func displayStates(tb *contractor.Toolbox, boards []tools.Board) {
	fmt.Printf("Box: %#v\n", tb)
	fmt.Println("Boards:")

	for _, b := range boards {
		fmt.Printf("\t%+v\n", b)
	}

	fmt.Println()
}

package contractor

import (
	"fmt"

	"github.com/BaiXilin/todd-udemy-go-course/tree/master/015-functions/08-interface-composition/tools"
)

type Contractor struct{}

// contractor drives nail into board using a NailDriver tool
func (Contractor) Fasten(d tools.NailDriver, nailSupply *int, b *tools.Board) {
	// keep driving nails until all needed in board
	for b.NailsDriven < b.NailsNeeded {
		d.DriveNail(nailSupply, b)
	}
}

// contractor removes nail from board using a NailPuller tool
func (Contractor) Unfasten(p tools.NailPuller, nailSupply *int, b *tools.Board) {
	// keep removing nails until there is no excess
	for b.NailsDriven > b.NailsNeeded {
		p.PullNail(nailSupply, b)
	}
}

func (c Contractor) ProcessBoards(dp tools.NailDriverPuller, nailSupply *int, boards []tools.Board) {
	for i := range boards {
		b := &boards[i]

		fmt.Printf("Contractor: examine board #%d: %+v\n", i+i, b)

		switch {
		case b.NailsDriven < b.NailsNeeded:
			c.Fasten(dp, nailSupply, b)
		case b.NailsNeeded < b.NailsDriven:
			c.Unfasten(dp, nailSupply, b)
		}
	}
}

// every good contractor should have a toolbox
type Toolbox struct {
	tools.NailDriver
	tools.NailPuller
	Nails int
}

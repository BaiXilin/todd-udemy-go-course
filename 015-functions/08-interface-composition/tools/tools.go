package tools

import "fmt"

// Board represents the surface we are working on
type Board struct {
	NailsNeeded int // fixed total number of nails needed
	NailsDriven int // number of nails already in place
}

// defines behavior to drive nail into board
type NailDriver interface {
	DriveNail(nailSupply *int, b *Board)
}

// defines behavior to remove nail from board
type NailPuller interface {
	PullNail(nailSupply *int, b *Board)
}

// Composition
// defines behavior to drive and remove nails from board
type NailDriverPuller interface {
	NailDriver
	NailPuller
}

// ACTUAL TOOLS
type Mallet struct{} // empty because it has no state. Only action

func (Mallet) DriveNail(nailSupply *int, b *Board) {
	// take a nail from supply and drive into the board
	*nailSupply--

	b.NailsDriven++

	fmt.Println("Mallet: pounded a nail into the board")
}

type Crowbar struct{}

func (Crowbar) PullNail(nailSupply *int, b *Board) {
	// remove a nail from the board and put into supply
	b.NailsDriven--

	*nailSupply++

	fmt.Println("Crowbar: yanked nail out of board")
}

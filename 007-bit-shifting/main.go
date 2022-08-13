package main

import "fmt"

const (
	_ = iota
	// kb is 1024 bytes (or 2^10, figure out the binary)
	kb = 1 << (10 * iota) // << and >> used for bit shifting
	// mb is 1024 kb
	mb = 1 << (10 * iota)
	// tb is 1024 mb
	tb = 1 << (10 * iota)
)

func main() {
	fmt.Printf("%d\t\t%#b\n", kb, kb)
	fmt.Printf("%d\t\t%#b\n", mb, mb)
	fmt.Printf("%d\t%#b\n", tb, tb)
}

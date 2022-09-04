package main

import (
	"fmt"
	"sort"
)

type Author struct {
	Name     string
	NumBooks int
}

func (a Author) String() string {
	return fmt.Sprintf("%s: %d books", a.Name, a.NumBooks)
}

type ByName []Author

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].Name < bn[j].Name }

type ByNumBooks []Author

func (bb ByNumBooks) Len() int           { return len(bb) }
func (bb ByNumBooks) Swap(i, j int)      { bb[i], bb[j] = bb[j], bb[i] }
func (bb ByNumBooks) Less(i, j int) bool { return bb[i].NumBooks < bb[j].NumBooks }

func main() {
	authors := []Author{
		{
			Name:     "JRR Tolkin",
			NumBooks: 4,
		},
		{
			Name:     "JRR Martin",
			NumBooks: 7,
		},
		{
			Name:     "Bradley",
			NumBooks: 18,
		},
	}

	fmt.Println("Original:", authors)
	sort.Sort(ByName(authors))
	fmt.Println("By Name:", authors)
	sort.Sort(ByNumBooks(authors))
	fmt.Println("By Book Nums:", authors)
}

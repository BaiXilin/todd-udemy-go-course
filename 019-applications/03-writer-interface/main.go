package main

import (
	"encoding/json"
	"io"
	"os"
)

type author struct {
	Name string
	Book string
	Length int
}

// w could be a http.ResponseWriter, could be a file
// by using io.Writer as argument type, author outputs same content to different Writer
func (a author) write(w io.Writer) {
	bb, _ := json.Marshal(a)
	w.Write(bb)
}

func main() {
	a1 := author {
		Name: "JRRT",
		Book: "Lord of the Rings",
		Length: 1000,
	}

	// a1 writes to a json file
	myFile, _ := os.Create("author.json")
	a1.write(myFile)

	// a1 writes to Stdout (os.Stdout is also io.Writer)
	a1.write(os.Stdout)
}

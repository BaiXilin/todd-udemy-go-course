package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string 	`json:"First"` // 'Tags', json name to golang member variable names
	Last string 	`json:"Last"`
	Age int		`json:"Age"`
}

func main() {
	data_string := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	byte_slice := []byte(data_string)
	var people []person

	err := json.Unmarshal(byte_slice, &people) // unmarshal requires a pointer as argument
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("%+v\n", people)
}

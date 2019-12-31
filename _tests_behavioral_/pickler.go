package main

import (
	"fmt"

	"github.com/abhishekkr/ficklepickle"
)

type Person struct {
	Name    string `json:"fullname"`
	Address string
	phone   Phone
}

type Phone struct {
	Home   string
	Office string
}

func main() {
	johnny := Person{Name: "Johnny", Address: "Wherever"}
	fmt.Println(johnny)

	pickle, err := ficklepickle.Pickle(johnny)
	fmt.Println(pickle)
	fmt.Println(err)

	j, err := ficklepickle.Unpickle(pickle, Person{})
	fmt.Println(j)
	fmt.Println(err)
}

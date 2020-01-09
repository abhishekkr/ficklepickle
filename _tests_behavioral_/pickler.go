package main

import (
	"fmt"

	"github.com/abhishekkr/ficklepickle"
	"github.com/abhishekkr/ficklepickle/config"

	golassert "github.com/abhishekkr/gol/golassert"
)

func init() {
	config.PickleDir = "./_tests_behavioral_"
}

type Person struct {
	Name    string `json:"fullname"`
	Address string
	Xphone  Phone
}

type Phone struct {
	Home   string
	Office string
}

func testPickleUnpickle(johnny Person) {
	pickle, err := ficklepickle.Pickle(johnny)
	golassert.AssertEqual(err, nil)

	j := Person{}
	err = ficklepickle.Unpickle(pickle, &j)
	golassert.AssertEqual(err, nil)

	golassert.AssertEqual(johnny.Name, j.Name)
	golassert.AssertEqual(johnny.Address, j.Address)
	golassert.AssertEqual(johnny.Xphone.Home, j.Xphone.Home)
	golassert.AssertEqual(johnny.Xphone.Office, j.Xphone.Office)
}

func testReadWrite(johnny Person) {
	johnny.Xphone = Phone{Home: "010101"}

	err := ficklepickle.Write("vanilla-file", "johnny", johnny)
	fmt.Println("~write:")
	fmt.Println(johnny)
	golassert.AssertEqual(err, nil)

	j := Person{}
	err = ficklepickle.Read("vanilla-file", "johnny", &j)
	fmt.Println("~read:")
	fmt.Println(j)
	golassert.AssertEqual(err, nil)

	golassert.AssertEqual(johnny.Name, j.Name)
	golassert.AssertEqual(johnny.Address, j.Address)
	golassert.AssertEqual(johnny.Xphone.Home, j.Xphone.Home)
	golassert.AssertEqual(johnny.Xphone.Office, j.Xphone.Office)
}

func main() {
	johnny := Person{Name: "Johnny", Address: "Wherever"}
	testPickleUnpickle(johnny)
	testReadWrite(johnny)
	fmt.Println("done.")
}

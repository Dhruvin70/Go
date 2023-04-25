package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	obj1 := Account{"Dhruvin", 52521, "IND"}
	Marshaling(obj1)
	// obj1.SetName("om")
	// obj2 := Account{obj1.GetName(), 541615, "cad"}
	// Marshaling(obj2)
}

func Marshaling(v Account) {

	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	//JSON objects only support strings as keys; to encode a Go map type it must be of
	//the form map[string]T (where T is any Go type supported by the json package).
	// b is of type byte

	fmt.Println(string(b))

}

type Account struct {
	Name    string
	ID      int64
	Country string
}

func (p *Account) SetName(Name string) {
	p.Name = Name
}

func (p Account) GetName() string {
	return p.Name
}

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Mor",
		Last:  "Fall",
		Age:   15,
	}
	p2 := person{
		First: "Coumba",
		Last:  "Diop",
		Age:   40,
	}
	a := []person{p1, p2}
	fmt.Println(a)
	s, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(s))

}

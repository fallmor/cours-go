package main

import "fmt"

type person struct {
	firstname        string
	lastname         string
	villedenaissance string
	age              int
}
type agentsecret struct {
	person
	nationalite string
	bl          bool
}

func main() {
	p1 := person{
		firstname:        "Mor",
		lastname:         "fall",
		villedenaissance: "Thies",
		age:              50,
	}
	p2 := agentsecret{
		person: person{
			firstname:        "Amadou",
			lastname:         "Diop",
			villedenaissance: "Dakar",
			age:              20,
		},
		nationalite: "français et senegalais",
		bl:          false,
	}
	fmt.Println(p1.firstname, p1.lastname, p1.villedenaissance, p1.age)
	fmt.Println(p2.firstname, p2.lastname, p2.villedenaissance, p2.age, p2.nationalite, p2.bl)
}
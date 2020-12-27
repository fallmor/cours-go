package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
}
type agentsecret struct {
	person
	nationalite string
	ok          bool
}

func (c agentsecret) test() {
	fmt.Println("Bonjour: ", c.firstname, c.lastname, ";Etre un agent secret: ", c.ok)
}
func (d person) test() {
	fmt.Println("Bonjour: ", d.firstname, d.lastname, ";la personne parle")
}

type human interface {
	test()
}

func foo(h human) {
	switch h.(type) {
	case agentsecret:
		fmt.Println("i ve been called by func barrr", h.(agentsecret).firstname)
	case person:
		fmt.Println("i ve been called by func barrr", h.(person).firstname)
	}
	fmt.Println("i ve been called by func bar", h)
}

func main() {
	s1 := agentsecret{
		person: person{
			firstname: "Mor",
			lastname:  "Fall",
			age:       50,
		},
		nationalite: "Française",
		ok:          true,
	}
	p1 := person{
		firstname: "Astou",
		lastname:  "FALL",
		age:       30,
	}
	fmt.Println(s1.ok)
	fmt.Println(p1)
	s1.test()
	p1.test()
	foo(s1)

	foo(p1)
}

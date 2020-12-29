package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "And I am", p.age, "years old")
}
func main() {

	p1 :=
		person{
			first: "Astou",
			last:  "Fall",
			age:   25,
		}
	p2 :=
		person{
			first: "Coumba",
			last:  "Diop",
			age:   40,
		}
	p1.speak()
	p2.speak()
}

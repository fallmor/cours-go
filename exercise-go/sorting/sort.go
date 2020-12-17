package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type SortByAge []person

func (a SortByAge) Len() int {
	return len(a)
}
func (a SortByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	p1 := person{
		Name: "Mor Fall",
		Age:  25,
	}
	p2 := person{
		Name: "Coumba Diop",
		Age:  35,
	}
	p3 := person{
		Name: "Astou Diop",
		Age:  20,
	}
	people := []person{p1, p2, p3}
	fmt.Println(people)
	sort.Sort(SortByAge(people))
	fmt.Println(people)
}

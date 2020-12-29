package main

import "fmt"

func main() {
	mp := map[string]int{
		"Mor":    15,
		"Astou":  10,
		"Coumba": 8,
		"Khady":  8,
	}
	fmt.Println(mp["Astou"])
	fmt.Println(mp["Anta"])
	if x, v := mp["Coumba"]; v {
		fmt.Printf("This %d is printed\n", x)
	}
	delete(mp, "Coumba")
	for v, k := range mp {
		fmt.Println(v, k)
	}

}

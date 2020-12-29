package main

import "fmt"

type person struct {
	nom              map[string]string
	favoriteIceCream []string
}

func main() {

	p1 := person{
		nom:              map[string]string{"Mor": "Fall", "Coumba": "Diop"},
		favoriteIceCream: []string{"Coco", "Vanilla", "Chocolate", "Strawberry"},
	}
	p2 := person{
		nom:              map[string]string{"Astour": "Seck", "Khady": "Sarr"},
		favoriteIceCream: []string{"Banana", "Apple", "Orange", "Coffe"},
	}
	fmt.Println(p1)
	fmt.Println(p2)

	for v, i := range p1.favoriteIceCream {
		fmt.Println(v, i)
	}
	for k, j := range p1.nom {
		fmt.Println(k, j)
	}

}

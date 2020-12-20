package main

import "fmt"

func main() {
	test(2, 4, 5, 7, 10, 15, 20, 33, 40)
}

func test(x ...int) {
	fmt.Println(x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("a l'index", i, "on a le numero", v, "on a le somme de num et des précédents", sum)
	}
}

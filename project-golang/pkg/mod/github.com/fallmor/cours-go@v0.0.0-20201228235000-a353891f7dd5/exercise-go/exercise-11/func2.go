package main

import "fmt"

func main() {
	z := []int{1, 2, 3, 4, 5}
	c, s := mafonc("Mor", z...)
	fmt.Println(c)
	fmt.Println(s)
}

func mafonc(b string, x ...int) (d string, a int) {
	fmt.Println(x)
	fmt.Println(b)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("Pour l'index", i, "On  la valeur", x, "et la somme avec la précédente fait", sum)
	}
	return d, a
}

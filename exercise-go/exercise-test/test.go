package main

import "fmt"

func main() {

	c := evensum(sum, []int{2, 3, 4, 5, 6, 7, 8, 9, 10}...)
	fmt.Println(c)

}

func sum(x ...int) int {
	v := 0
	for _, u := range x {
		v += u
	}
	return v
}
func evensum(f func(x ...int) int, y ...int) int {
	var xi []int
	for _, j := range y {
		if j%2 == 0 {
			xi = append(xi, j)
		}
	}
	total := f(xi...)
	return total

}

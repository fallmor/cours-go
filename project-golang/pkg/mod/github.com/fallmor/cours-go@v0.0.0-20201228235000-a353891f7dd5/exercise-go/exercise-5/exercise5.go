package main

import "fmt"

func main() {
	x := []int{12, 3, 4, 5}
	fmt.Println(x)
	y := []int{120, 8, 7, 9}
	x = append(x, y...)
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}
}

package main

import "fmt"

func main() {
	x := make([]int, 6, 7)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[5] = 7
	x = append(x, 12)
	x = append(x, 1)
	x = append(x, 52)
	x = append(x, 22)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 1, 3, 4, 5, 6)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

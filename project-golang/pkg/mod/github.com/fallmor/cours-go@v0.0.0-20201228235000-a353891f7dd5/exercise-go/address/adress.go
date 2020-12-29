package main

import "fmt"

func main() {
	x := 10
	fmt.Println(&x)
	fmt.Printf("%T", &x)
}

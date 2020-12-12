package main

import "fmt"

type mor int

var x mor

func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)
}

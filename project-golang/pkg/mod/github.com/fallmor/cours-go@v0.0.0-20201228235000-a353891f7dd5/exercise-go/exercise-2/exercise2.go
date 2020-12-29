package main

import "fmt"

func main() {
	s := "hello, Mr Mor"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	sb := []byte(s)
	fmt.Println(sb)
	fmt.Printf("%T\n", sb)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#b\n", s[i])
	}
	for v, i := range s {
		fmt.Printf("At position %d we have %#x\n", v, i)
	}
}

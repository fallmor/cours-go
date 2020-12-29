package main

import "fmt"

func main() {
	var nom string
	fmt.Println("Bonjour entrez votre nom: ")
	fmt.Scan(&nom)
	fmt.Println("votre bom est:", nom)
	mafonc()
}
func mafonc() {
	var age int
	fmt.Println("Entrez votre age: ")
	fmt.Scan(&age)
	for age := 0; age < 100; age++ {
		if age%2 == 0 {
			fmt.Println("votre est: ", age, "vous etes pair")
		} else {
			fmt.Println("votre est: ", age, "vous etes impair")
		}
	}
}

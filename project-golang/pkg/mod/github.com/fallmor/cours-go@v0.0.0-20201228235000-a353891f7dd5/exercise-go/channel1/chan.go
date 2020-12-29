package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)
	receive(even, odd, quit)
	fmt.Print("about to exit")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Printed from chan even \t", v)
		case v := <-o:
			fmt.Println("Printed from chan odd \t", v)
		case v := <-q:
			fmt.Println("Printed from chan quit \t", v)
			return
		}

	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

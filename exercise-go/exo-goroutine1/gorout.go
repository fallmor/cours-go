package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Begin CPU", runtime.NumCPU())
	fmt.Println("Begin Goroutines", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("I run first")
		wg.Done()
	}()
	go func() {
		fmt.Println("I run second")
		wg.Done()
	}()
	fmt.Println("Mid CPU", runtime.NumCPU())
	fmt.Println("Mid Goroutines", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("I exited")
	fmt.Println("End CPU", runtime.NumCPU())
	fmt.Println("End Goroutines", runtime.NumGoroutine())

}

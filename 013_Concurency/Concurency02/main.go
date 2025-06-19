package main

import (
	"fmt"
	"runtime"
	"sync"
)

func foo() {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println("Foo : ",i)
	}
}

func bar() {
	for i := 1; i <= 5; i++ {
		fmt.Println("bar : ",i)
	}
}

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t" ,runtime.GOARCH)
	fmt.Println("CPUs\t",runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	
	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t",runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}	
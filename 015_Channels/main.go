package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 42
	number := <-c
	if number == 42{
		fmt.Println("True")
	}

}
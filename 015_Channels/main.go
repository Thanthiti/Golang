package main

import (
	"fmt"
	"time"
)
func start(){
		c := make(chan int,2)
	// [[],[]]
	
	c <- 10
	// [[10],[]]

	number := <-c
	// [[],[]] number received 10 from channel 
	
	fmt.Println(number)

}

func main() {
	ch := make(chan int)
	go func()  {
		time.Sleep(time.Second * 2)
		fmt.Println("Test 1")
		ch <- 10
		}()

		fmt.Println("Test 2")
		fmt.Println("Test 3")
		
		v := <- ch
		fmt.Println("Test 4")
	fmt.Println(v)
}
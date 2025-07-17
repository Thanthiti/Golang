package main

import (
	"fmt"
	"time"
)

func HelloSet1() {
	fmt.Println("Hello 1")
	fmt.Println("Hello 2")
}
func HelloSet2() {
	fmt.Println("Hello 3")
	fmt.Println("Hello 4")
}
func main() {
	go HelloSet1()
	HelloSet2()
	time.Sleep(time.Second)
}
package main

import "fmt"

func increment() func(int,int) int {
	x := 0
	return func(a,b int) int {
		x++
		fmt.Print(a)
		fmt.Print(b)
		return x
	}

}

func main() {
	number := increment()
	fmt.Println(number(1,2))
	fmt.Println(number(3,4))
	fmt.Println(number(5,6))

}
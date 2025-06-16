package main

import "fmt"

type com interface{
	comparable 
}

func findEqualNumber[T com](a, b T) string {
	fmt.Println("asd")
	return fmt.Sprintln("result is ",(a==b))
}

func main() {
	a :=findEqualNumber(12,15)
	fmt.Println(a)
}
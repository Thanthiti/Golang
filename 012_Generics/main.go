package main

import (
	"fmt"
)


func addNumber[T int | float64](a , b T) T {
	return a-b
}

func main() {
	fmt.Println("Value is = ",addNumber(1.9,3.8))
}
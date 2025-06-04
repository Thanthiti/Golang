package main

import "fmt"

func domath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}
func subtract(a, b int) int {
	return a - b
}

func main() {

	plus := domath(1, 2, add)
	minus := domath(5, 3, subtract)

	fmt.Println(plus)
	fmt.Println(minus)
}
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
	plus := domath(3, 4, add)
	fmt.Println(plus)

}
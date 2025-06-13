package main

import "fmt"

func a(n *int) {
	*n += 2
	b(n)
}
func b(n *int) {
	*n += 2
	c(n)
}
func c(n *int) {
	*n += 2
}
func main() {
	number := 1
	fmt.Println("Initial : ",number)
	a(&number)
	fmt.Println("After : ",number)
}
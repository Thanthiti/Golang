package main

import "fmt"

func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x - factor
	}
}
func main() {
	func() {
		fmt.Println("this Annonymous FN")
	}()

	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	fmt.Println(double(5)) // 10
	fmt.Println(triple(5)) // 15

	double = makeMultiplier(5)
	triple = makeMultiplier(5)
	fmt.Println(double(3)) // 10
	fmt.Println(triple(2)) // 15
}
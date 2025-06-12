package main

import "fmt"

func main() {
	fmt.Println("Hello")
	x := 42
	fmt.Println(x)
	fmt.Println(&x)
	y := &x
	fmt.Println(y)
	fmt.Println(&y)
	fmt.Println(*y)
	fmt.Println("after assign new value")
	*y = 43
	fmt.Println(y)
	fmt.Println(*y)
	fmt.Println(&y)

	fmt.Println(&x)
	fmt.Println(*&x)

}
package main

import (
	"fmt"
)

var globol = 10

const (
	_ = iota
    A 
    B        
    C        
	D
	E
	F
)

func printVariable(num int){
	fmt.Println("this num : ",num)
}

func main() {
	a := 42
	m := 43
	// var h int = 40
	// fmt.Print(h)
	fmt.Println(a)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	fmt.Println(m)
	const Z = iota
	fmt.Println(Z)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f,)

	b, d, c,n := 0, 1, 2,4
	fmt.Println(b, c, d,n)

	var g int
	fmt.Println(g)
	g = 43
	fmt.Println(g)

	var h int = 44
	fmt.Println(h)
}

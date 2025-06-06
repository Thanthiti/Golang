package main

import (
	"fmt"
)

type Retangle struct{
	width , height float64

}

func main() {
	fmt.Println("Hello World!")
	rec := Retangle{3,4}
	fmt.Println(rec.height)
	
}
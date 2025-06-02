package main

import "fmt"

func main() {

	defer foo()
	 bar()
	defer bark()
}

func foo(){
	fmt.Println("Foo")
}
func bark(){
	fmt.Println("Bark")
}
func bar(){
	fmt.Println("Bar")
}


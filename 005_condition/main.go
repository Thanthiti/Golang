package main

import (
	"fmt"
	"math/rand"
)
func main(){
	if z := 2 * rand.Intn(5); z > 6{
		fmt.Printf("The value of z is %v\n",z)
	} else {
		fmt.Printf("The value of z %v less than 6 ",z)

	}
}
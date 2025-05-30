package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	fmt.Println("Hello")
	number := rand.Intn(6) + 5
	fmt.Println(number)

	Slices := []int{4,1,5,2,3,3 }
	fmt.Println("Before sort : ",Slices)
	sort.Ints(Slices)
	fmt.Println("After sort : ",Slices)

}



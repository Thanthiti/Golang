package main

import "fmt"

func doubleArray(arr *[5]int) {
	for i:= 0 ;i< len(arr);i++ {
		arr[i] *= 2
	}
}
func doubleSlice(arr []int) {
	for i:= 0 ;i< len(arr);i++ {
		arr[i] *= 2
	}
}

func main() {
	// Array
	fmt.Println("\nArray--------")
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Original: ", a)
	doubleArray(&a)
	fmt.Println("Modified: ", a)
	
	// Slice
	fmt.Println("\nSlice--------")
	b := []int{1, 2, 3, 4, 5}
	fmt.Println("Original: ", b)
	doubleSlice(b)
	fmt.Println("Modified: ", b)
}

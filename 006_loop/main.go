package main

import (
	"fmt"
)

func main(){

	// first loop
	// for i := 1 ; i <= 5 ; i++{
	// 	fmt.Printf("Loop 1 , the value of x is %v\n",i)
	// }
	// fmt.Println("Stop first loop!!")

	// // second loop
	// x := 1	
	// for x <= 5{
	// 	fmt.Printf("Loop 2 , the value of x is %v\n",x)
	// 	// if x <= 10
	// 	x++
	// }
	// fmt.Println("Stop second loop!!")

	// //third loop 
	// x = 1
	// for{
	// 	fmt.Printf("Loop 3 , the value of x is %v\n",x)
	// 	if x >= 5{
	// 		fmt.Println("Stop third loop!!")
	// 		break
	// 	}
	// 	x++
	// }

	// for range loop and Array
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}
	index := 0
	for index > len(xi){
		fmt.Println("ranging over a slice",index)
		index++
	}

	m := map[int]string{
		1:     "42",
		2: "32",
	}
	for k, v := range m {
		fmt.Println("ranging over a map", k, v)
	}

	slice := []int{}
	fmt.Println(slice)
	for i := 0; i <=10;i++ {
		slice = append(slice, i)
	}
	fmt.Println(slice)

	a := [10]int{}
	fmt.Println(a)
	for i := 0; i < 10;i++ {

		a[i] = i
	}
	fmt.Println(a)


}
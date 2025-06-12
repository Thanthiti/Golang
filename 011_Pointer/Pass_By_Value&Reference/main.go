package main

import "fmt"

func ChangePointer(n *int,number int) {
	*n += 5
	number += 2
}

func SliceReference(Arr[] int){
	for i,_ := range Arr{
		Arr[i] += 10
	}
}
func mapReference(data map[string]int,s string){
	data[s] = 23
}

func main() {
	a := 1
	n := 2
	fmt.Printf("Before call function : %v \t %v\n",a,n)
	ChangePointer(&a,n)
	fmt.Printf("After call function : %v \t %v\n",a,n)
	
	slices := []int{1,2,3,4}	
	fmt.Println("Before call function : ",slices)
	SliceReference(slices)
	fmt.Println("After call function : ",slices)

	name := make(map[string]int)
	name["palm"] = 22
	fmt.Println("Before call function : ",name)
	mapReference(name,"palm")
	fmt.Println("After call function : ",name)


}
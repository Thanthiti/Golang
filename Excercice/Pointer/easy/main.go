package main

import "fmt"

func swap(x ,y *int){
	temp := *y
	*y = *x
	*x = temp
}

func main(){
	a := 10
	b := 20
	
	fmt.Printf("Befor swap : a = %v, b = %v \n",a,b)
	swap(&a,&b)
	fmt.Printf("after swap : a = %v, b = %v \n",a,b)
}
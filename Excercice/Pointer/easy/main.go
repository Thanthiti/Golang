package main

import "fmt"

func swap(x ,y *int){
	temp := *y
	*y = *x
	*x = temp
}

func main(){
	// a := 10
	// b := 20
	
	// fmt.Printf("Befor swap : a = %v, b = %v \n",a,b)
	// swap(&a,&b)
	// fmt.Printf("after swap : a = %v, b = %v \n",a,b)

	x := 10
	p := &x
	

	fmt.Println("ค่าของ x:", x)          // 10
	fmt.Println("ที่อยู่ของ x:", &x)     // เช่น 0xc000018098
	
	fmt.Println("ค่าของ pointer p:", p) // เช่น 0xc000018098
	fmt.Println("ค่าที่ pointer p ชี้:", *p) // 10
	
	*p = 20
	fmt.Println("ค่าของ x:", x)          // 10
	fmt.Println("ที่อยู่ของ x:", &x)     // เช่น 0xc000018098
	fmt.Println("ค่าของ pointer p:", p) // เช่น 0xc000018098
	fmt.Println("ค่าที่ pointer p ชี้:", *p) // 10

}
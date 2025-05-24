package main

import "fmt"


func SayHello(){
	fmt.Println("Hello World")
}

func SayMyName(firstname, lastname string,age int ) {
	fmt.Printf("My name's %s  %s now i'm %d\n",firstname,lastname,age)
}

func MathFunc(n1,n2 int) int{
	return n1+n2
}
func tstcout(){
	x := 10
	if x == 11 {fmt.Printf("asd")}
}


func main(){
	fmt.Println("asd")
	SayHello()
	SayMyName("Thanthiti","Pinmolee",22)
	fmt.Println("Result of func Mathfunc is " , MathFunc(10,20),"")


}
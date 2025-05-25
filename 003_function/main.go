package main

import "fmt"

const (
	n1 = 10
	n2 = 30
)

func SayHello(){
	fmt.Println("Hello World")
}

func SayMyName(firstname, lastname string,age int ) {
	fmt.Printf("My name's %s  %s now i'm %d\n",firstname,lastname,age)
}

func MathFunc(n1,n2 int) int{
	return n1+n2
}

func Return2param(x,y string) (string,int){
	return y,22
}

func NameReturn(sum int) (x,y int){
	x = sum + 5
	y = sum - 5
	return 
}

func main(){
	SayHello()

	SayMyName("Thanthiti","Pinmolee",22)

	fmt.Println("Result of func Mathfunc is " , MathFunc(10,20),"")

	a, b := Return2param("pin","than")
	fmt.Printf("my name is %s and %d\n",a,b)

	c,d := NameReturn(10)
	fmt.Println(c,d)	

	n :=10
	fmt.Println("this n before : ",n)
	n = 20
	fmt.Println("this n after : ",n)
	fmt.Print("this const ",n1+20,n2)
	



}
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
	return fmt.Sprint("My last name is ",x," ",y),22
}

func NameReturn(sum int) (x,y int){
	x = sum + 5
	y = sum - 5
	return 
}
//  Variadic Function **** 
func Sum(Number ...int){
	Sum := 0
	for _,n := range Number{
		Sum+=n
	}
	fmt.Println(Number[1])
	fmt.Println("Sum of Number is : ",Sum)
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
	fmt.Println("this const ",n1+20,n2)

	fmt.Println("--------------")
	t,v := na("test",12)
	fmt.Println(t)
	fmt.Println(v)
	Sum(1,2,3,4)
	Slice := []int{1,2,3,4,5}
	Sum(Slice...)
	

}
func na(name string, age int)(string ,int){
	return fmt.Sprint(name," is thi old "),age
}
package main

import "fmt"


type person struct{
	first string
	last string
	age int
}

func (p person) Call(){
	fmt.Println("Welcome Person : " , p.first)
}
func (p person) Call1(){
	fmt.Println("Welcome Home : " , p.first)
}
type people struct{
	person
	status bool
}
func (pp people) CallPP(){
	fmt.Println("There Are Too many People : ",pp.first )
}

type human interface{
	Call1()
	Call()
	
}
func SaySomeThings(h human){
	h.Call()
	h.Call1()
}
func main(){

	p1 := person{
		first: "Test",
	}
	p1.Call()
	fmt.Println("----------------------")
	SaySomeThings(p1)
	fmt.Println("----------------------")
	ShowStruct()
	EmbeddedStruct()
	
}

func ShowStruct() {

	fmt.Println("hello World!")
	p1 := person{
		first: "palm",
		last: "pa",
		age: 12,
	}
	fmt.Println(p1.first)
}

func EmbeddedStruct(){

	// p1 := person{
	// 	first: "palm",
	// 	last: "pa",
	// 	age: 12,
	// }

	// group2 := people{
	// 	person: p1,
	// 	status: true,
	// }

	group1 := people{
		person: person{
			first: "2palm",
			last: "pa1",
			age: 32,	
		},
		status: true,
	}
	// fmt.Println(group2)
	// fmt.Println(group2.first)
	fmt.Println(group1)
	fmt.Println(group1.first)
	group1.CallPP()
	fmt.Println("----------------------")
	SaySomeThings(group1)
}
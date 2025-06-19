package main

import (
	"fmt"
	"sort"
	// "strings"
)

type Person struct{
	First string
	Age int
}

type WeightAvG struct{
	First string
	Weight int
}
func NormalSort(){
	Slices := []int{4, 1, 5, 2, 3, 3}
	fmt.Println("Before sort : ", Slices)
	sort.Ints(Slices)
	fmt.Println("After sort : ", Slices)

	func(){
		sort.Slice(Slices,func(i, j int) bool {
			return Slices[i] > Slices[j]
		})
		}()
	fmt.Println("After sort most to least : ", Slices)


	slicesString := []string{"Plam","Beer","Dream","Arm","arm"}
	fmt.Println("Before sort : ", slicesString)
	sort.Strings(slicesString)
	fmt.Println("After sort : ", slicesString)
	
	sort.Slice(slicesString, func(i, j int) bool {
		return (slicesString[i]) > (slicesString[j])
	})
	fmt.Println("After sort : ", slicesString)
}
func CustomSort(){
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}
	people := []Person{p1,p2,p3,p4}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}	
type ByAge []Person

func (a ByAge) Swap(i, j int){ 
	fmt.Println("1")
	a[i], a[j] = a[j], a[i] 
}
func (a ByAge) Len() int{ 
	fmt.Println("2")
	return len(a) 
}
func (a ByAge) Less(i, j int) bool{ 
	fmt.Println("3")
	return a[i].Age < a[j].Age 
}

type Cal []WeightAvG

func main() {
	// NormalSort()
	CustomSort()

	// p1 := Person{"B", 53}
	// p2 := Person{"P", 54}
	// p3 := Person{"Q", 64}
	// p4 := Person{"M", 56}

	// people := []Person{p1,p2,p3,p4}

	
}
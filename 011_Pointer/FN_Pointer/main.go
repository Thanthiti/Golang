package main

import "fmt"

type yougin interface{
	bark()
	rename()
}

type dog struct {
	name string
}

func (d dog) bark() {
	fmt.Println("Hello ", d.name)
}
func (d *dog) rename() {
	d.name = "Hong hong"
	fmt.Println("Hello i'm running", d.name)
}

func youngRun(y yougin){
	y.rename()
	y.bark()
}

func main() {
	d1 := dog{"Beer"}
	d2 := &dog{"palm"}
	
	fmt.Println("\n This D1 ")
	d1.bark()
	d1.rename()
	fmt.Println("\nThis D1 ",d1.name)
	
	fmt.Println("\n This D2 ")
	d2.bark()
	d2.rename()
	fmt.Println("\nThis D2 ",d2.name)
	
	fmt.Println("\n--------------------")

	// youngRun(d1)
	youngRun(d2)

}
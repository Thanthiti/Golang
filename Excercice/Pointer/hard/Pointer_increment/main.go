package main

import "fmt"

type Counter struct {
	value int
}

func (v *Counter) Increment() {
	v.value += 1
}

func (v *Counter) Reset() {
	fmt.Println("Reset...")
	v.value = 0
}
func main() {
	value := Counter{0}
	fmt.Printf("value : %v\n",value.value)
	value.Increment()
	value.Increment()
	value.Increment()

	value.Increment()
	fmt.Printf("value : %v\n",value.value)
	value.Reset()
	fmt.Printf("value : %v\n",value.value)
}
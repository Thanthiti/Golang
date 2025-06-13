package main

import "fmt"

type Person struct {
	name string
	age  int
}

func haveBirthday(p *Person) {
	p.age+= 1
}
func main() {
	person := Person{
		name : "Palm",
		age: 21,
	}
	fmt.Printf("Before: %s is %d years old\n", person.name, person.age)

	haveBirthday(&person)

	fmt.Printf("After birthday: %s is %d years old\n", person.name, person.age)
}
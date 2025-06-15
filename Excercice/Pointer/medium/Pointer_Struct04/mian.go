package main

import "fmt"

type Employee struct {
	name string
	age  int
}
type Manager struct {
	EmployeeInfo *Employee
}

type Department struct {
	Head *Manager
}

func IncreaseAge(de *Department) {
	de.Head.EmployeeInfo.age += 10
}

func main() {
	Employee := Employee{"palm", 22}
	manager := Manager{&Employee}
	depaetment := Department{&manager}
	fmt.Printf("Before : %v  , %v", Employee.name,Employee.age)
	IncreaseAge(&depaetment)
	fmt.Printf("Before : %v  , %v", Employee.name,Employee.age)
}
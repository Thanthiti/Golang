package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Team struct {
	Leader *Person
}

func UpdateLeader(t Team,newName string) {
	t.Leader.name = newName
}
func main() {
	person := Person{"Palm", 22}
	leader := Team{&person}
	newName := "Thanthiti"
	fmt.Println("Before Change Leader : ", person.name)
	leader.Leader.name = newName
	// UpdateLeader(leader,newName)
	fmt.Print("After Change Leader : ", person.name)

}
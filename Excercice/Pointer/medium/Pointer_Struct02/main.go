package main

import "fmt"

type Boss struct {
	name string
}

type Company struct {
	boss *Boss
}

func renameBoss(c *Company, newName string) {
	c.boss.name = newName
}

func main() {
	comp := Company{boss: &Boss{name: "John"}}
	fmt.Println("Old Boss:", comp.boss.name)
	renameBoss(&comp, "Jane")
	fmt.Println("New Boss:", comp.boss.name)
}

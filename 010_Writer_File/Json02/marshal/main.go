package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		"palm", "pin", 12,
	}
	p2 := person{
		"beer", "org", 22,
	}
	people := []person{p1, p2}

  bs, _ := json.Marshal(people)
  s, _ := json.Marshal(p1)
	fmt.Println(people)
  fmt.Println(string(bs))
  fmt.Println(string(s))


}
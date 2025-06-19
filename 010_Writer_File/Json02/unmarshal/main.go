package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"palm","Last":"pin","Age":12},{"First":"beer","Last":"org","Age":22}]`
	bs := []byte(s)
	fmt.Println(string(bs))


	// people := []person{}
	var people []person

	err := json.Unmarshal(bs,&people)
	if err != nil{
		fmt.Println("Error")
	}

	for _,n := range people{
		fmt.Println(n)
	}

}
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	person := person{
		Name:  "palm",
		Age:   30,
		Email: "palm@mail.com",
	}
	file, err := os.Create("010_Writer_File/Json/data.json")
	if err != nil{
		log.Fatalf("Eror %s",err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent(""," ")
	err = encoder.Encode(person)
	fmt.Println(err)
	if  err != nil{
		log.Fatalf("Error encoding JSON: %v", err)
	}
	log.Println("SuccessFully!")
}
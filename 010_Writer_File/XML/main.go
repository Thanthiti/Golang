package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type person struct {
	Name  string `xml:"name"`
	Age   int    `xml:"age"`
	Email string `xml:"email"`
}

func main() {
	person := person{
		Name:  "palm",
		Age:   30,
		Email: "palm@mail.com",
	}
	file, err := os.Create("010_Writer_File/xml/data.xml")
	if err != nil{
		log.Fatalf("Eror %s",err)
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	err = encoder.Encode(person)
	fmt.Println(err)
	if  err != nil{
		log.Fatalf("Error encoding xml: %v", err)
	}
	log.Println("SuccessFully!")
}
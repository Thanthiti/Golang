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

func readJson(fileName string)(person,error){
	file,err := os.Open(fileName)
	if err != nil{
		return person{},fmt.Errorf("failed to openee JSON file: %w", err)
		
	}
	defer file.Close()
	
	var person person
	deCoder := json.NewDecoder(file)
	if err := deCoder.Decode(&person); err != nil{
		return person,fmt.Errorf("failed to open JSON file: %w", err)
	}
	return person,nil
}

func writeJson(fileName string, person person) error{
	file,err := os.Create(fileName)
	if err != nil{
		return fmt.Errorf("failed to create JSON file: %w", err)
	}
	defer file.Close()

	enCoder := json.NewEncoder(file) 
	enCoder.SetIndent(""," ")
	err = enCoder.Encode(person)
	if err != nil{
		return fmt.Errorf("failed to encode JSON: %w", err)
	}
	return nil
}

func main() {
	person := person{
		Name:  "palm",
		Age:   30,
		Email: "palm@mail.com",
	}
	// Encoder Json 
	jsonFile := "010_Writer_File/Json/data.json"
	if err := writeJson(jsonFile,person); err != nil{
		log.Fatalf("Error writing JSON: %v", err)
	}
	
	// Decoder Json
	data,err := readJson(jsonFile)
	if err != nil{
		log.Fatalf("Error Reading JSON: %v", err)
	}
	log.Printf("Read from JSON: %+v\n", data)
	fmt.Println(data.Age)
	

}

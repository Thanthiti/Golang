package main

import (
	// "io"
	"log"
	"os"
)

// type person struct {
// 	first string
// }

// func (p person) writeout(w io.Writer) {
// 	_,err := w.Write([] byte(p.first))
// }
func main() {
	f,err := os.Create("010_Writer_File/Write/output.txt")
	if err != nil{
		log.Fatalf("Error %s",err)
	}
	defer f.Close()
	
	// 
	s := []byte("Hello Wordl!")
	_,err = f.Write(s)


	text := []string{
		"Line 1: Hello, Go!",
		"Line 2: Writing multiple lines.",
		"Line 3: This overwrites the file each time.",
	}
	for _,line := range text{
		_,err = f.WriteString(line+"\n")
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	}
	if err != nil{
		log.Fatalf("Error %s",err)
	}
	
}
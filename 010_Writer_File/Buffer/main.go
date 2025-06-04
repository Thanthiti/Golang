package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("Hello ")
	b.WriteString("World!\n")
	b.WriteString("It's Thursday\n")
	fmt.Println(b.String())
	
	fmt.Println("------------")
	b.Reset()
	b.WriteString("Helllppp!\n")
	fmt.Println(b.String())

	b.Write([]byte("Happy Happy"))
	fmt.Println(b.String())
}
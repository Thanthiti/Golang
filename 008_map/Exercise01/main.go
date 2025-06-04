
// Count words and sort them from most to least and show which word counts more.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Asd")
	// file, err := os.Open("great-gatsby.txt")
	file, err := os.Open("great-gatsbyTest.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	Worlds := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		Worlds = append(Worlds, strings.Split(line, " ")...)
	}
	fmt.Println(len(Worlds))
	fmt.Println(Worlds)
	for i,v := range Worlds{
		fmt.Printf("%v \t %v \n",i,v)
	}
}
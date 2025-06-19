package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  fmt.Print("Enter Number : ")
  reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadString('\n')
  input = strings.TrimSpace(input)
  
  num, _ := strconv.Atoi(input) 
  fmt.Printf("You entered int : %t\n", num)
 
  
  fmt.Printf("The converted integer is: %t", num) 

 }
 
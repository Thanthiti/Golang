package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  // Output Example
  fmt.Println("Hello borntoDev Lab")
  
  // Input Example
  reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadString('\n')
  input = strings.TrimSpace(input)
  
  num, err := strconv.Atoi(input) // Convert string to integer
  fmt.Println("You entered int:", num)
  if err != nil {
	  fmt.Println("Error converting string to integer:", err)
	  return
  }
  for i:=1; i<= num ;i++{
	for n := 1 ; n <= i;n++{
		fmt.Print("*")
	}
	fmt.Println("")
}
  fmt.Println("The converted integer is:", num) // Output: 12345

 }
 
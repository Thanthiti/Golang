package main

import (
  "fmt"
  "sync"
  "time"
)

var m sync.Mutex

var n = 10

func p() {
  m.Lock()
  fmt.Println("LOCK")
  
  fmt.Println(n)

  time.Sleep(1 * time.Second)
  m.Unlock()
 
  fmt.Println("UNLOCK")
}

func main() {
  fmt.Println("FIRST")
 
  go p()
 
  fmt.Println("SECOND")
 
  p()
 
  fmt.Println("THIRD")
  time.Sleep(time.Second*2 )
 
  fmt.Println("DONE")
}
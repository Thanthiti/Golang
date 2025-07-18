package main

import (
  "fmt"
  "sync"
)
func main() {
  var once sync.Once
  var wg sync.WaitGroup

  initialize := func() {
    fmt.Println("Initializing only once")
  }

  doWork := func(workerId int) {
    defer wg.Done()
    fmt.Printf("Worker %d started\n", workerId)
    once.Do(initialize) // This will only be executed once
    fmt.Printf("Worker %d done\n", workerId)
  }

  numWorkers := 5
  wg.Add(numWorkers)

  // Launch several goroutines
  for i := 1; i <= numWorkers; i++ {
    go doWork(i)
  }

  // Wait for all goroutines to complete
  wg.Wait()
  fmt.Println("All workers completed")
}
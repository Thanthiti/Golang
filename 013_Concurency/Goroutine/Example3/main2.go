package main

import (
  "fmt"
  "math/rand"
  "sync"
  "time"
)

func worker(id int, wg *sync.WaitGroup) {
  defer wg.Done() // Decrement the counter when the goroutine completes

  fmt.Printf("Worker %d starting\n", id)

  // Simulate some work by sleeping
  sleepDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
  time.Sleep(sleepDuration)

  fmt.Printf("Worker %d done\n", id)
}

func main() {
  var wg sync.WaitGroup

  // Launch several goroutines and increment the WaitGroup counter for each
  for i := 1; i <= 5; i++ {
	  wg.Add(1)
    go worker(i, &wg)
  }

  wg.Wait() // Block until the WaitGroup counter goes back to 0; all workers are done

  fmt.Println("All workers completed")
}
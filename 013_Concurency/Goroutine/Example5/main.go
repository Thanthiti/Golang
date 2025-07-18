package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a new condition variable
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	// A shared resource
	ready := false

	// A goroutine that waits for a condition
	go func() {
		fmt.Println("Goroutine: Waiting for the condition...")

		mutex.Lock()

		// if goroutine FN into wait state mutex wii be go wait state to
		// And mutex.lock's not effective.
		for !ready {
			cond.Wait() // Wait for the condition
		}
		fmt.Println("Goroutine: Condition met, proceeding...")
		mutex.Unlock()

	}()

	// Simulate some work (e.g., loading resources)
	time.Sleep(2 * time.Second)

	// Signal the condition
	mutex.Lock()
	ready = true
	cond.Signal() // Signal one waiting goroutine
	mutex.Unlock()
	fmt.Println("Push signal !")

	// Give some time for the goroutine to complete
	time.Sleep(1 * time.Second)
	fmt.Println("Main: Work is done.")
}

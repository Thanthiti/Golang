package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Cpus",runtime.NumCPU())
	fmt.Println("Goroutines",runtime.NumGoroutine())

	var  counter int64
	const gs int = 100
	var wg sync.WaitGroup
	// var mu sync.Mutex

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			//  Mutex
			// mu.Lock()
			// v := counter
			// runtime.Gosched()
			// v++
			// counter = v 		
			// mu.Unlock()

			// Atomic
			atomic.AddInt64(&counter,1)	
			runtime.Gosched()
			fmt.Println("Counter \t", atomic.LoadInt64(&counter))
			}()
	}
	wg.Wait()
	fmt.Println("Goroutines",runtime.NumGoroutine())
	fmt.Println("Counter : ",counter)
}

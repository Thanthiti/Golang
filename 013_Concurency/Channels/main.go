package main

import (
	"fmt"
)

func buffered() {
	c := make(chan int, 2)
	// [[],[]]

	c <- 10
	// [[10],[]]

	number := <-c
	// [[],[]] number received 10 from channel

	fmt.Println(number)

}

func Unbuffered() {
	ch := make(chan int)
	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

func SelectGoroutine(){
		ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 10
		close(ch1)
	}()
	go func() {
		ch2 <- 20
		close(ch2)
	}()

	closedCheckch1, closedCheckch2 := false, false
	for {
		if closedCheckch1 && closedCheckch2 {
			break
		}
		select {
		case v, ok := <-ch1:
			if !ok {
				closedCheckch1 = true
				continue
			}
			fmt.Println("Channel 1 ", v)
			
		case v, ok := <-ch2:
			if !ok {
				closedCheckch2 = true
				continue
			}
			fmt.Println("Channel 2 ", v)
		}
	}
}

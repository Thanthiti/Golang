package main
import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello")
}

func sayHi() {
    fmt.Println("Hi")
}

func main() {
    go sayHello() 
    go sayHi()   

    time.Sleep(1 * time.Second) // รอให้ goroutine ทำงานก่อนจบโปรแกรม
}

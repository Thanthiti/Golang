package main

import (
  "fmt"
  "time"
)

func main() {
  // สร้าง channel เพื่อส่งข้อความ
  ch := make(chan string)

  // สร้าง goroutine เพื่อส่งข้อความไปยัง channel
  go func() {
    for i := 0; i < 10; i++ {
      ch <- fmt.Sprintf("Hello, world! %d", i)
      time.Sleep(1 * time.Second)
    }
  }()

  // สร้าง goroutine เพื่อรับข้อความจาก channel
  go func() {
    for {
      msg := <-ch
      fmt.Println(msg)
    }
  }()

  // รอให้ goroutines ทำงานเสร็จสิ้น
  time.Sleep(3 * time.Second)
}
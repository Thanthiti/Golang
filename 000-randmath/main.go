package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	fmt.Println("Hello")
	number := rand.Intn(6) + 5
	fmt.Println(number)

}

// package main

// import (
//     "fmt"
//     "math/rand"
//     "time"
// )

// func main() {
//     rand.Seed(time.Now().UnixNano()) // ตั้ง seed เพื่อไม่ให้ได้เลขซ้ำ
//     number := rand.Intn(6) + 5       // rand.Intn(6) จะให้ค่า 0 ถึง 5 → บวก 5 → ได้ 5 ถึง 10
//     fmt.Println("สุ่มได้เลข:", number)
// }

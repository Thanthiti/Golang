package main

import "fmt"

type User struct {
    Name string
    Age  int
}

// รับ struct แบบ value → ทำงานกับสำเนา
func ChangeName(u User) {
    u.Name = "Changed"
}

func main() {
    user := User{Name: "Original", Age: 20}

    ChangeName(user)

    fmt.Println(user.Name) // ผลลัพธ์: Original → ไม่เปลี่ยน!
}


//  รับ pointer → แก้ของจริง
// func ChangeName(u *User) {
//     u.Name = "Changed"
// }

// func main() {
//     user := User{Name: "Original", Age: 20}

//     ChangeName(&user)

//     fmt.Println(user.Name) // ผลลัพธ์: Changed → เปลี่ยนจริง!
// }
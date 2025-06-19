package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "mySecret123"

	// 🧂 Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hashed password:", string(hashedPassword))

	// ✅ เปรียบเทียบรหัสผ่านที่ user พิมพ์มา กับ hash ที่เก็บไว้
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte("mySecret123"))
	if err != nil {
		fmt.Println("❌ Password is incorrect!")
	} else {
		fmt.Println("✅ Password is correct!")
	}
}

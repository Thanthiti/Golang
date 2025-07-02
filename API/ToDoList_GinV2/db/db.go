package db

// import (
// 	"log"

// 	// "mymodule/API/server03.2_ToDoList_Gin/model"

// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func InitDB() {
// 	var err error
// 	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect database: ", err)
// 	}

// 	// ❌ อย่า AutoMigrate ถ้าใช้ migration จริง
// 	DB.AutoMigrate(&model.Todo{})
// }

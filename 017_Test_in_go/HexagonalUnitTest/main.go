package main

import (
	"fmt"
	"log"
	"mymodule/017_Test_in_go/HexagonalUnitTest/adapters"
	"mymodule/017_Test_in_go/HexagonalUnitTest/core"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func InitializeDB() *gorm.DB{
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{	})

	// Check Connected database
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&core.Order{})
	return db
}

func main() {
	app := fiber.New()
	db := InitializeDB()

	// NewGormOrderRepository return OrderRepository interface 
	orderRepo := adapters.NewGormOrderRepository(db)
	// NewOrderService return OrderService interface
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order/",orderHandler.CreateOrder)

	// Migeate schema
	app.Listen(":8080")
}
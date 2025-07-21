package main

import (
	"mymodule/018_Hexagonal/adapters"
	"mymodule/018_Hexagonal/core"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()


	//Initialize the database 
	db,err := gorm.Open(sqlite.Open("orders.db"),&gorm.Config{})
	if err != nil{
		panic("failed to connect database")
	}
	// NewGormOrderRepository return OrderRepository interface 
	orderRepo := adapters.NewGormOrderRepository(db)
	// NewOrderService return OrderService interface
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order/",orderHandler.CreateOrder)


	// Migeate schema
	db.AutoMigrate(&core.Order{})
	app.Listen(":8080")
}
package main

import (
	"log"
	"mymodule/019_Clearn_Architecture/adapters"
	"mymodule/019_Clearn_Architecture/entities"
	"mymodule/019_Clearn_Architecture/usecases"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
  app := fiber.New()

  db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
  if err != nil {
    log.Fatalf("failed to connect database: %v", err)
  }

  if err := db.AutoMigrate(&entities.Order{}); err != nil {
    log.Fatalf("failed to migrate database: %v", err)
  }

  orderRepo := adapters.NewGormOrderRepository(db)
  orderService := usecases.NewOrderService(orderRepo)
  orderHandler := adapters.NewHttpOrderHandler(orderService)

  app.Post("/order", orderHandler.CreateOrder)

  log.Fatal(app.Listen(":8000"))
}
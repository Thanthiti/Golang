package adapters

import (
	"mymodule/019_Clean_Architecture/entities"
	"mymodule/019_Clean_Architecture/usecases"

	"github.com/gofiber/fiber/v2"
)

type HttpOrderHandler struct {
  orderUseCase usecases.OrderUseCase
}

func NewHttpOrderHandler(service usecases.OrderUseCase) *HttpOrderHandler {
	return &HttpOrderHandler{orderUseCase: service}
}


func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
  var order entities.Order
  if err := c.BodyParser(&order); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
  }

  if err := h.orderUseCase.CreateOrder(order); err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
  }

  return c.Status(fiber.StatusCreated).JSON(order)
}
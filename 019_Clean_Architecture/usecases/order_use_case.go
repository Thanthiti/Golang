package usecases

import "mymodule/019_Clean_Architecture/entities"

type OrderUseCase interface {
  CreateOrder(order entities.Order) error
}

type OrderService struct {
  repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderUseCase {
  return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order entities.Order) error {
  return s.repo.Save(order)
}
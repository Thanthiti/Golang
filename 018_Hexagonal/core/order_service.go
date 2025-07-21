package core

import "errors"

type OrderService interface {
	CreateOrder(order Order) error
}

type OrderServiceImpl struct {
	repo OrderRepository
}

func newOrderService(repo OrderRepository) OrderService {
	return &OrderServiceImpl{repo: repo}
}

func (s *OrderServiceImpl) CreateOrder(Order Order) error {
	if Order.Total <= 0 {
		return errors.New("Total must be positive")
	}

	if err := s.repo.Save(Order); err != nil{
		return err
	}
	return nil
}
package core

import (
	"errors"
)

// Primary port for an input requirement of order service
type OrderService interface{
	CreateOrder(order Order) error
}

// Service implements Primary and Secondary ports to be able to connect them.
type OrderServiceImpl struct{
	repo OrderRepository
}
func (s *OrderServiceImpl) CreateOrder(order Order) error {
	// Business logic
	if order.Total <= 0 {
		return errors.New("Total must be positive")
	}

	if err := s.repo.SaveOrder(order); err != nil {
		return err
	}

	return nil
}

func NewOrderService(repo OrderRepository) OrderService{
	return &OrderServiceImpl{repo: repo}
}

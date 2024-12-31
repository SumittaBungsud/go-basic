package usecases

import (
	"errors"
	"go-clean/entities"
)

// Input port
type OrderUseCase interface{
	CreateOrder(order entities.Order) error
}

// Service (implement input and output ports)
type OrderService struct{
	repo OrderRepository
}
func (s *OrderService) CreateOrder(order entities.Order) error{
	if order.Total <= 0 {
		return errors.New("total must be more than 0")
	}
	return s.repo.Save(order)
}

// function to create the service instance
func NewOrderService(repo OrderRepository) OrderUseCase{
	return &OrderService{repo: repo}
}
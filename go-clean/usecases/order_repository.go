package usecases

import "go-clean/entities"

// Output port
type OrderRepository interface{
	Save(order entities.Order) error
}
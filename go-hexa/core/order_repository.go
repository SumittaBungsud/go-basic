package core

// Secondary port for an output requirement of order insertion in database
type OrderRepository interface{
	SaveOrder(order Order) error
}
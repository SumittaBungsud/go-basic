package adapters

import (
	"go-clean/entities"
	"go-clean/usecases"

	"gorm.io/gorm"
)

// Output (order repository) adapter connects to gorm's db
type GormOrderRepository struct{
	db *gorm.DB // External service
}
func (r *GormOrderRepository) Save(order entities.Order) error{ // implement output port
	return r.db.Create(&order).Error
}

// function to create gorm's DB adapter instance
func NewGormOrderRepository(db *gorm.DB) usecases.OrderRepository{
	return &GormOrderRepository{db: db}
}
package adapters

import (
	"go-hexa/core"

	"gorm.io/gorm"
)

// Secondary adapter (output)
type GormOrderRepository struct{
	db *gorm.DB
}
func (r *GormOrderRepository) SaveOrder(order core.Order) error{
	if result := r.db.Create(&order); result.Error != nil {
		return result.Error
	}

	return nil
}

func NewGormOrderRepository(db *gorm.DB) core.OrderRepository{
	return &GormOrderRepository{db: db}
}
package adapters

import (
	"mymodule/019_Clearn_Architecture/entities"
	usecase "mymodule/019_Clearn_Architecture/usecases"

	"gorm.io/gorm"
)

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) usecase.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Save(order entities.Order) error{
	
	return r.db.Create(&order).Error
}
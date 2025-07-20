package usecases

import "mymodule/019_Clearn_Architecture/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}
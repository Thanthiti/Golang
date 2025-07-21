package usecases

import "mymodule/019_Clean_Architecture/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}
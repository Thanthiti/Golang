package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockOrderRepo struct {
	saveFunc func(Order Order) error
}

func (m *MockOrderRepo) Save(order Order) error {
	return m.saveFunc(order)
}
func TestCreateOrder(t *testing.T){
	// Success case
	t.Run("Success",func(t *testing.T) {

		repo := &MockOrderRepo{
			saveFunc: func(Order Order) error {
				return nil
			},
		}
		service := NewOrderService(repo)
		
		err := service.CreateOrder(Order{Total: 100})
		assert.NoError(t,err)
	})
	t.Run("Total less than 0",func(t *testing.T) {

		repo := &MockOrderRepo{
			saveFunc: func(Order Order) error {
				return nil
			},
		}
		service := NewOrderService(repo)
		
		err := service.CreateOrder(Order{Total: -10})
		assert.Error(t,err)
		assert.Equal(t,"Total must be positive",err.Error())
	})
}

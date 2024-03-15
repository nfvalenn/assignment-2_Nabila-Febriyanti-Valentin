package repository

import "assignment-2/model"

type IOrderRepository interface {
	Create(newOrder model.Order) (model.Order, error)
	GetAll() ([]model.Order, error)
	Delete(id int) error
	Update(updatedOrder model.Order) (model.Order, error)
}
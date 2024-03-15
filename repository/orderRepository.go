package repository

import (
	"assignment-2/model"

	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository  {
	return &orderRepository{
		db: db,
	}
}

func (o *orderRepository) Create(newOrder model.Order) (model.Order, error)  {
	tx := o.db.Create(&newOrder)
	return newOrder, tx.Error
}

func (o *orderRepository) GetAll() ([]model.Order, error)  {
	var orders = []model.Order{}

	tx := o.db.Find(&orders)
	return orders, tx.Error
}

func (o *orderRepository) Delete(id int) error  {
	tx := o.db.Unscoped().Delete(&model.Order{}, "id = ?", id)
	return tx.Error
}

func (o *orderRepository) Update(updatedOrder model.Order) (model.Order, error)  {
	tx := o.db.Save(&updatedOrder)
	return updatedOrder, tx.Error
}


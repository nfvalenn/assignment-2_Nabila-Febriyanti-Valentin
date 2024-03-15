package repository

import "assignment-2/model"

type IItemRepository interface {
	Create(newItem model.Items) (model.Items, error)
	GetAll() ([]model.Items, error)
	Delete(id int) error
	Update(updatedItem model.Items) (model.Items, error)
}
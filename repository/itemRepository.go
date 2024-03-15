package repository

import (
	"assignment-2/model"

	"gorm.io/gorm"
)

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *itemRepository  {
	return &itemRepository {
		db: db,
	}
}

func (i *itemRepository) Create(newItem model.Items) (model.Items, error)  {
	tx := i.db.Create(&newItem)
	return newItem, tx.Error
}

func (i *itemRepository) GetAll() ([]model.Items, error)  {
	var items = []model.Items{}

	tx := i.db.Find(&items)
	return items, tx.Error
}

func (i *itemRepository) Delete(id int) error {
	tx := i.db.Unscoped().Delete(&model.Items{}, "id = ?", id)
	return tx.Error
}

func (i *itemRepository) Update(updatedItem model.Items) (model.Items, error)  {
	tx := i.db.Save(&updatedItem)
	return updatedItem, tx.Error
}
package model

import "gorm.io/gorm"

type Items struct {
	gorm.Model
	ID			uint	`json:"item_id" gorm:"primaryKey"`
	ItemCode    string	`json:"item_code"`
    Description string	`json:"description"`
    Quantity    int		`json:"quantity"`
    OrderID		uint	`json:"order_id"`
}


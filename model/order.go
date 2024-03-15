package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID		   	 int		`json:"order_id" gorm:"primaryKey"`
	CustomerName string		`json:"customer_name"`
	OrderedAT    *time.Time	`json:"ordered_at"`
	Items		 []Items	`gorm:"foreignKey:OrderID"`
}
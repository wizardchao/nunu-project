package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
}

// func (m *Order) TableName() string {
// 	return "order"
// }

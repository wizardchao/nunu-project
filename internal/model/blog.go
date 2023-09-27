package model

import (
	"time"
)

type Blog struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Content   string `gorm:"not null"`
	Deleted   uint   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (m *Blog) TableName() string {
	return "blog"
}

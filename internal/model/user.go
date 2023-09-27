package model

import (
	"nunu-project/pkg/helper/localTime"
)

type User struct {
	Id        uint                `json:"id" gorm:"primarykey"`
	UserId    string              `json:"userId" gorm:"unique;not null"`
	Username  string              `json:"username" gorm:"unique;not null"`
	Nickname  string              `json:"nickname" gorm:"not null"`
	Password  string              `json:"-" gorm:"not null"`
	Email     string              `json:"email" gorm:"not null"`
	CreatedAt localTime.LocalTime `json:"createdAt"`
	UpdatedAt localTime.LocalTime `json:"updatedAt"`
	deletedAt localTime.LocalTime `gorm:"index"`
}

func (u *User) TableName() string {
	return "users"
}

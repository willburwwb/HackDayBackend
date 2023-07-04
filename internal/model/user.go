package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Phone        string `gorm:"not null;unique" json:"phone"`
	PasswordHash string `gorm:"column:passwordHash"`
}

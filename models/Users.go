package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Firstname string `gorm:"not null`
	LastName  string `gorm:"not null`
	Email     string `gorm:"not null; unique_index`
	Tasks     []Tasks
}

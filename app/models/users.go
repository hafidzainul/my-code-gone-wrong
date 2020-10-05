package models

import (
	"gorm.io/gorm"
)

// User Models.
type User struct {
	gorm.Model
	Fullname string
	Email    string
	Password string
	OfficeID int
	Office   Office `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Todos    []Todos
}

// UserTodos Models.
type UserTodos struct {
	Fullname    string
	Email       string
	Name        string
	Description string
}

package models

import (
	"gorm.io/gorm"
)

// Todos Models
type Todos struct {
	gorm.Model
	Name        string
	Description string
	UserID      int
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

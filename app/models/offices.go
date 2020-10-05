package models

import "gorm.io/gorm"

// Office Models.
type Office struct {
	gorm.Model
	Name    string
	Address string
}

// OfficeUser Models.
type OfficeUser struct {
	Name     string
	Address  string
	Fullname string
	Email    string
}

// OfficeTodos Models.
type OfficeTodos struct {
	Name        string
	Address     string
	Jobs        string
	Description string
}

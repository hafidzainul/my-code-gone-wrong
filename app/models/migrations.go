package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations comment.
func Migrations(db *gorm.DB) {
	var check bool

	check = db.Migrator().HasTable(&Office{})
	if !check {
		db.Migrator().CreateTable(&Office{})
		fmt.Println("Create Table Office")
	}

	check = db.Migrator().HasTable(&User{})
	if !check {
		db.Migrator().CreateTable(&User{})
		fmt.Println("Create Table User")
	}

	check = db.Migrator().HasTable(&Todos{})
	if !check {
		db.Migrator().CreateTable(&Todos{})
		fmt.Println("Create Table User")
	}
}

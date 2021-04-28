package model

import (
	"github.com/jinzhu/gorm"
)

type Language struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:book_languages;"`
}

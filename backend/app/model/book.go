package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Price      int        `validate:"required,max=100"`
	Title      string     `validate:"required"`
	Content    string     `validate:"required"`
	BookDetail BookDetail `validate:"omitempty,nostructlevel"`
	Comments   []Comment
	Languages  []Language `gorm:"many2many:book_languages;"`
}

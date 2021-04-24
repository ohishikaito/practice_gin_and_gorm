package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Price      int    `validate:"required,max=100"`
	Title      string `validate:"required"`
	Content    string `validate:"required"`
	BookDetail BookDetail
	Comments   []Comment
}

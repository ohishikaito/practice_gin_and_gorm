package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Price   int
	Title   string
	Content string
}

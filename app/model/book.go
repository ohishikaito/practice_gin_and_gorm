package model

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title   string
	Content string
}

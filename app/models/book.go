package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title   string
	Content string
}

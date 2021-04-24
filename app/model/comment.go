package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	UserID uint
	BookID uint
	Text   string `validate:"required"`
}

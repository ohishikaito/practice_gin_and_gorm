package model

import (
	"github.com/jinzhu/gorm"
)

type BookDetail struct {
	gorm.Model
	BookID      uint   `gorm:"foreignKey"`
	Description string `validate:"required"`
}

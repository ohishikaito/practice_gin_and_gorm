package model

import (
	"github.com/jinzhu/gorm"
)

type BookDetail struct {
	gorm.Model
	// BookIDがnilでもcreateできちゃうけど、commentsを作ってから。後回し！
	BookID      uint   `gorm:"foreignKey"`
	Description string `validate:"required"`
}

package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	// ID        uint `gorm:"primary_key" json:"id"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	gorm.Model
	Email string `validate:"required"`
	// パスワードは管理しないことにしました〜
	// Password  string `json:"password"`
	LastName  string
	FirstName string
	Comments  []Comment
}

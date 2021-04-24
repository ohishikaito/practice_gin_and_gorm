package model

import "time"

type User struct {
	// 手動でID振ったら、自動的にincrementされるのかなあ？
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// gorm.Model
	Email    string `json:"email" validate:"required"`
	Password string `json:"password"`
}

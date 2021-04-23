package model

import "time"

type User struct {
	// 手動でID振ったら、自動的にincrementされるのかなあ？
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// gorm.Model
	email    string
	password string
}

package model

import "github.com/jinzhu/gorm"

// jsonとformはいらん気がする！リクエストの形式次第だけど〜
type Book struct {
	gorm.Model
	Title   string
	Content string
}

package model

import (
	"github.com/jinzhu/gorm"
)

type BookDetail struct {
	gorm.Model
	// TODO: BookIDがnilでもcreateできちゃうけど、commentsを作ってから。後回し！
	// BookID      uint   `gorm:"foreignKey"`
	BookID int
	// Book   Book `gorm:"foreignkey:BookID,constraint"`
	// Book        Book   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Description string `validate:"required"`
}

// ↓のparamsを送ると、bookと一緒に保存できる。外部キー制約は貼れないけどw
// {
//     "description": "hoge",
//     "book": {
//         "price": 1,
//         "title": "title",
//         "content": "content"
//     }
// }

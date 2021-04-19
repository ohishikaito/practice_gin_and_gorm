package models

// import (
// 	"go_myapp/db"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/jinzhu/gorm"
// )

// type Book struct {
// 	gorm.Model
// 	ID      uint   `json:"id" gorm:"primary_key"`
// 	Title   string `json:"title"`
// 	Content string `json:"contect"`
// }

// // type BookService struct{}

// // // jsonで受け取る用のstruct
// // // update時にどうゆう挙動になるか要検証。
// // // idは含める必要がありそうなので入れたけど、created_atとupdated_atは送らなくても良いのかな？
// // // gorm.Modelのstructに共通化出来たら楽そう！
// // type BookJsonRequest struct {
// // 	Id      int    `json:"id"`
// // 	Title   string `json:"title"`
// // 	Contect string `json:"contect"`
// // }

// func GetBookList() []Book {
// 	db, _ := db.SqlConnect()
// 	books := make([]Book, 0)
// 	db.Find(&books)
// 	return books
// }

// func CreateBook(book *Book) (err error) {
// 	db, _ := db.SqlConnect()
// 	if err := db.Create(&book).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// // func (BookService) SetBook(book *model.Book) error {
// //     _, err := DbEngine.Insert(book)
// //     if err!= nil{
// //          return  err
// //     }
// //     return nil
// // }

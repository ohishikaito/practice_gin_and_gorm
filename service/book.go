package service

import (
	"go_myapp/app/models"
	"go_myapp/db"
)

type BookService struct{}

func (BookService) GetBookList() []models.Book {
	db, _ := db.SqlConnect()
	books := make([]models.Book, 0)
	db.Find(&books)
	return books
}

func (BookService) CreateBook(book *models.Book) (err error) {
	db, _ := db.SqlConnect()
	if err := db.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

// func (BookService) SetBook(book *model.Book) error {
//     _, err := DbEngine.Insert(book)
//     if err!= nil{
//          return  err
//     }
//     return nil
// }

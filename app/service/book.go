package service

import (
	"go_myapp/app/model"
)

type BookService struct{}

func (BookService) GetBookIndex() (books []model.Book) {
	books = make([]model.Book, 0)
	if err := db.Find(&books).Error; err != nil {
		panic(err.Error())
	}
	return
}

func (BookService) SetBook(book *model.Book) error {
	if err := db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

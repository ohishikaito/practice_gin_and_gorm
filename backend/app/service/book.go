package service

import (
	"app/app/model"
)

type BookService struct{}

func (BookService) GetBookIndex() ([]model.Book, error) {
	var books []model.Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (BookService) CreateBook(book *model.Book) error {
	if err := validate.Struct(book); err != nil {
		return err
	}
	if err := db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (BookService) ShowBook(bookID int) (*model.Book, error) {
	book := &model.Book{}
	if err := db.Where("id = ?", bookID).Preload("Comments.User").Preload("BookDetail").Preload("Languages").Take(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (BookService) UpdateBook(book *model.Book) (*model.Book, error) {
	if err := validate.Struct(book); err != nil {
		return nil, err
	}
	if err := db.Model(&model.Book{}).Where("id = ?", book.ID).Update(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (BookService) DeleteBook(bookID int) error {
	book := &model.Book{}
	if err := db.Where("id = ?", bookID).Take(book).Delete(book).Error; err != nil {
		return err
	}
	return nil
}

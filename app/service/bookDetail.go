package service

import (
	"go_myapp/app/model"
)

type BookDetailService struct{}

func (BookDetailService) CreateBookDetail(bookDetail *model.BookDetail) error {
	if err := validate.Struct(bookDetail); err != nil {
		return err
	}
	if err := db.Create(bookDetail).Error; err != nil {
		return err
	}
	return nil
}

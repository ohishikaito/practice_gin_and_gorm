package service

import (
	"go_myapp/app/model"
)

type BookDetailService struct{}

func (BookDetailService) GetBookDetailIndex() (model.BookDetail, error) {
	var bookDetail model.BookDetail
	if err := db.Find(&bookDetail).Error; err != nil {
		// nilにすると何故かエラーになる。配列じゃないからが原因説濃厚。空のstructを返す
		return model.BookDetail{}, err
	}
	return bookDetail, nil
}

func (BookDetailService) CreateBookDetail(bookDetail *model.BookDetail) error {
	if err := validate.Struct(bookDetail); err != nil {
		return err
	}
	if err := db.Create(bookDetail).Error; err != nil {
		return err
	}
	return nil
}

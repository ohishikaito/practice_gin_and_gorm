package service

import (
	"go_myapp/app/model"
)

type CommentService struct{}

func (CommentService) GetCommentIndex(bookID int) (*model.Book, error) {
	book := &model.Book{}
	if err := db.Where("id = ?", bookID).Preload("Comments").Find(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

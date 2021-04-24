package service

import (
	"go_myapp/app/model"
)

type CommentService struct{}

func (CommentService) CreateComment(comment *model.Comment) error {
	if err := validate.Struct(comment); err != nil {
		return err
	}
	if err := db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

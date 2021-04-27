package service

import (
	"fmt"
	"go_myapp/app/model"

	"github.com/jinzhu/gorm"
)

type CommentService struct {
}

func (CommentService) CreateComment(comment *model.Comment) error {
	if err := validate.Struct(comment); err != nil {
		return err
	}
	// レスポンスにUser含めたくないんだけど、nullにもできないからとりあえず含めるようにした！
	if err := db.Create(comment).Take(&comment.User).Error; err != nil {
		return err
	}
	return nil
}

func (CommentService) GetBookComments(bookID int) (*[]model.Comment, error) {
	comments := &[]model.Comment{}
	if err := db.Scopes(CommentsByBookID(bookID), PreloadUser).Find(comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func CommentsByBookID(bookID int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("book_id = ?", bookID)
	}
}

func PreloadUser(db *gorm.DB) *gorm.DB {
	return db.Preload("User")
}

func (CommentService) UserCommentBooks(userID int) (*[]model.Book, error) {
	books := &[]model.Book{}
	query := fmt.Sprintln(`
		SELECT DISTINCT b.* FROM users AS u
			INNER JOIN comments AS c ON c.user_id = u.id
			INNER JOIN books AS b ON b.id = c.book_id
		WHERE c.user_id = ?
	`)
	err := db.Raw(query, userID).Scan(books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

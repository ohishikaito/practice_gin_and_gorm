package service

import (
	"go_myapp/app/model"
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

func (BookService) FindBookById(bookID int) (*model.Book, error) {
	book := &model.Book{}
	// fmt.Println(db.Where("id = ?", bookID))
	// fmt.Println(db.Where("id = ?", bookID).Take(book))
	// fmt.Println(db.Where("id = ?", bookID).Take(book).Association("bookDetail"))
	// fmt.Println(book)
	if err := db.Where("id = ?", bookID).Take(book).Association("bookDetail").Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (BookService) UpdateBook(book *model.Book, data model.Book) error {
	// update時にvalidateすると、全てのカラムが必須になってしまうのでコメントアウト
	// 値を差し替えるべきなのかなあ
	// if err := validate.Struct(data); err != nil {
	// 	return err
	// }
	if err := db.Model(book).Update(data).Error; err != nil {
		return err
	}
	return nil
}

func (BookService) DeleteBook(book *model.Book) error {
	if err := db.Delete(book).Error; err != nil {
		return err
	}
	return nil
}

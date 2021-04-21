package service

import (
	"go_myapp/app/model"
)

type BookService struct{}

// serviceでエラー吐いたら、コントローラーにエラー返すようにしたいなあ〜、そっちのが分かりやすそう
func (BookService) GetBookIndex() (books []model.Book) {
	books = make([]model.Book, 0)
	if err := db.Find(&books).Error; err != nil {
		panic(err.Error())
	}
	return
}

// クリエイトしたら、作成したインスタンスも返すように
func (BookService) CreateBook(book *model.Book) error {
	if err := db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (BookService) FindBookById(bookID int) (*model.Book, error) {
	book := &model.Book{}
	if err := db.Where("id = ?", bookID).Take(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (BookService) UpdateBook(book *model.Book) error {
	if err := db.Update(book).Error; err != nil {
		return err
	}
	return nil
}

// func (BookService) UpdateBook(newBook *model.Book) error {
//     _, err := DbEngine.Id(newBook.Id).Update(newBook)
//     if err != nil {
//         return err
//     }
//     return nil
// }

func (BookService) DeleteBook(book *model.Book) error {
	if err := db.Delete(book).Error; err != nil {
		return err
	}
	return nil
}

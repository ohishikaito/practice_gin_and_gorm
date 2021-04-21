package controller

import (
	"go_myapp/app/model"
	"go_myapp/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var bookService service.BookService

func BookIndex(c *gin.Context) {
	books := bookService.GetBookIndex()
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func BookCreate(c *gin.Context) {
	book := model.Book{}
	// BindJSONとShouldBindJSONどっちにすべきなんだろうか？
	if err := c.BindJSON(&book); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	if err := bookService.CreateBook(&book); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": book,
	})
}

func BookShow(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	book, err := bookService.FindBookById(bookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

// updateはむずいから一旦保留
func BookUpdate(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	_, err := bookService.FindBookById(bookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"data": err.Error(),
		})
		return
	}
}

// func BookUpdate(c *gin.Context){
//     book := model.Book{}
//     err := c.Bind(&book)
//     if err != nil{
//         c.String(http.StatusBadRequest, "Bad request")
//         return
//     }
//     bookService :=service.BookService{}
//     err = bookService.UpdateBook(&book)
//     if err != nil{
//         c.String(http.StatusInternalServerError, "Server Error")
//         return
//     }
//     c.JSON(http.StatusCreated, gin.H{
//         "status": "ok",
//     })
// }

func BookDelete(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	book, err := bookService.FindBookById(bookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"data": err.Error(),
		})
		return
	}
	bookService.DeleteBook(book)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

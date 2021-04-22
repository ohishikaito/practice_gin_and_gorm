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
	books, err := bookService.GetBookIndex()
	if err != nil {
		c.String(http.StatusInternalServerError, ""+err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func BookCreate(c *gin.Context) {
	book := model.Book{}
	c.BindJSON(&book)
	if err := bookService.CreateBook(&book); err != nil {
		c.String(http.StatusUnprocessableEntity, ""+err.Error())
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
		c.String(http.StatusNotFound, ""+err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func BookUpdate(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	book, err := bookService.FindBookById(bookID)
	if err != nil {
		c.String(http.StatusNotFound, ""+err.Error())
		return
	}
	data := model.Book{}
	c.BindJSON(&data)
	if book.ID != data.ID {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	if err := bookService.UpdateBook(book, data); err != nil {
		c.String(http.StatusUnprocessableEntity, ""+err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func BookDelete(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	book, err := bookService.FindBookById(bookID)
	if err != nil {
		c.String(http.StatusNotFound, ""+err.Error())
		return
	}
	bookService.DeleteBook(book)
	c.String(http.StatusOK, "success")
}

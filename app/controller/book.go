package controller

import (
	"errors"
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
		app.ErrorResponse(c, http.StatusNotFound, err)
		return
	}
	app.JSONResponse(c, http.StatusOK, books)
}

func BookCreate(c *gin.Context) {
	book := model.Book{}
	c.BindJSON(&book)
	if err := bookService.CreateBook(&book); err != nil {
		app.ErrorResponse(c, http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": book,
	})
}

func BookShow(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	book, err := bookService.ShowBook(bookID)
	if err != nil {
		app.ErrorResponse(c, http.StatusNotFound, err)
		return
	}
	app.JSONResponse(c, http.StatusOK, book)
}

func BookUpdate(c *gin.Context) {
	paramID, _ := strconv.Atoi(c.Param("id"))
	data := &model.Book{}
	c.BindJSON(data)
	// これなかったら、違うIDにリクエスト送って書き換えれる。でも必要なんですかねー？
	if data.ID != uint(paramID) {
		err := errors.New("Bad Request")
		app.ErrorResponse(c, http.StatusNotFound, err)
		return
	}
	book, err := bookService.UpdateBook(data)
	if err != nil {
		app.ErrorResponse(c, http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func BookDelete(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	if err := bookService.DeleteBook(bookID); err != nil {
		app.ErrorResponse(c, http.StatusNotFound, err)
		return
	}
	c.String(http.StatusOK, "deleted!")
}

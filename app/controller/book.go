package controller

import (
	"go_myapp/app/model"
	"go_myapp/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BookIndex(c *gin.Context) {
	bookService := service.BookService{}
	books := bookService.GetBookIndex()
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   books,
	})
}

func BookCreate(c *gin.Context) {
	book := model.Book{}
	// BindJSONとShouldBindJSONどっちにすべきなんだろうか？
	// if err := c.ShouldBindJSON(&book); err != nil {
	if err := c.BindJSON(&book); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService := service.BookService{}
	if err := bookService.SetBook(&book); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

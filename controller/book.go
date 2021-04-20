package controller

import (
	"go_myapp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BookIndex(c *gin.Context) {
	bookService := service.BookService{}
	Books := bookService.GetBookIndex()
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   Books,
	})
}

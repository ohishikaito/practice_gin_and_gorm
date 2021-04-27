package controller

import (
	"go_myapp/app/model"
	"go_myapp/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var bookDetailService service.BookDetailService

func BookDetailCreate(c *gin.Context) {
	bookDetail := model.BookDetail{}
	c.BindJSON(&bookDetail)
	if err := bookDetailService.CreateBookDetail(&bookDetail); err != nil {
		app.ErrorResponse(c, http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": bookDetail,
	})
}

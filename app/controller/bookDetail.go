package controller

import (
	"go_myapp/app/model"
	"go_myapp/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var bookDetailService service.BookDetailService

func BookDetailIndex(c *gin.Context) {
	bookDetail, err := bookDetailService.GetBookDetailIndex()
	if err != nil {
		c.String(http.StatusInternalServerError, ""+err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": bookDetail,
	})
}

func BookDetailCreate(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	bookDetail := model.BookDetail{
		BookID: bookID,
	}
	c.BindJSON(&bookDetail)
	if err := bookDetailService.CreateBookDetail(&bookDetail); err != nil {
		c.String(http.StatusUnprocessableEntity, ""+err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": bookDetail,
	})
}
package controller

import (
	"go_myapp/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var commentService service.CommentService

func CommentIndex(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	book, err := commentService.GetCommentIndex(bookID)
	if err != nil {
		c.String(http.StatusInternalServerError, ""+err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

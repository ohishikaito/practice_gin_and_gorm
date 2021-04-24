package controller

import (
	"go_myapp/app/model"
	"go_myapp/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var commentService service.CommentService

func CommentCreate(c *gin.Context) {
	comment := &model.Comment{}
	c.BindJSON(comment)
	if err := commentService.CreateComment(comment); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
		return
	}
	app.JSONResponse(c, http.StatusCreated, comment)
}

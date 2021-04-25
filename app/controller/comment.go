package controller

import (
	"go_myapp/app/model"
	"go_myapp/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var commentService service.CommentService

func CommentCreate(c *gin.Context) {
	comment := &model.Comment{}
	c.BindJSON(comment)
	if err := commentService.CreateComment(comment); err != nil {
		app.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
		return
	}
	app.JSONResponse(c, http.StatusCreated, comment)
}

// コメントしたユーザー一覧
// コメント一覧を取得して、そこからユーザーを取得した方がいいのではwwwwwww
// 後でスコープ化する
func CommentIndex(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	comments, err := commentService.GetBookComments(bookID)
	if err != nil {
		app.ErrorResponse(c, http.StatusInternalServerError, err)
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
		return
	}
	app.JSONResponse(c, http.StatusOK, comments)
}

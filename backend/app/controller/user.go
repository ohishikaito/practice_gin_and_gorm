package controller

import (
	"app/app/model"
	"app/app/service"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userService service.UserService

func (app *App) UserIndexHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := userService.GetUserIndex()
	type data struct {
		Data []model.User `json:"data"`
	}
	body := data{
		Data: users,
	}
	app.successResponse(w, body, http.StatusOK)
}

func (app *App) userCreateHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	if err := userService.CreateUser(&user); err != nil {
		app.errorResponse(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	type data struct {
		Data model.User `json:"data"`
	}
	body := data{
		Data: user,
	}
	app.successResponse(w, body, http.StatusCreated)
}

func UserCreate(c *gin.Context) {
	user := model.User{}
	c.BindJSON(&user)
	if err := userService.CreateUser(&user); err != nil {
		app.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, user)
}

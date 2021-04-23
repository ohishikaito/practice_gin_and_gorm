package controller

import (
	"encoding/json"
	"go_myapp/app/model"
	"go_myapp/app/service"
	"net/http"
)

var userService service.UserService

func (app *App) UserIndexHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := userService.GetUserIndex()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	type data struct {
		Users []model.User `json:"data"`
	}
	json.NewEncoder(w).Encode(data{users})
}

func (app *App) userCreateHandler(w http.ResponseWriter, r *http.Request) {
}

package controller

import (
	"github.com/gorilla/mux"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func HttpRoute(app *App) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", app.UserIndexHandler).Methods("GET")
	r.HandleFunc("/users", app.userCreateHandler).Methods("POST")
	return r
}

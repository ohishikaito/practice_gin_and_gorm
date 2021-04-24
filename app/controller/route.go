package controller

import (
	"encoding/json"
	"go_myapp/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (app *App) successResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func (app *App) errorResponse(w http.ResponseWriter, Response string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	type JSONResponse struct {
		Response string `json:"response"`
	}
	jsonError, err := json.Marshal(JSONResponse{Response: Response})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)
}

func HttpRoute(app *App) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", app.UserIndexHandler).Methods("GET")
	r.HandleFunc("/users", app.userCreateHandler).Methods("POST")
	return r
}

func Route() {
	engine := gin.Default()
	engine.Use(middleware.TestMiddleware())
	bookEngine := engine.Group("/books")
	{
		bookEngine.GET("/", BookIndex)
		bookEngine.POST("/", BookCreate)
		bookEngine.GET("/:id", BookShow)
		bookEngine.PUT("/:id/", BookUpdate)
		bookEngine.DELETE("/:id", BookDelete)
	}
	engine.Run()
}

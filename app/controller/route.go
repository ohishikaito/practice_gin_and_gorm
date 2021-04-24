package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct{}

var app *App

func (app *App) JSONResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"data": data,
	})
}

package controller

import (
	"github.com/gin-gonic/gin"
)

type App struct{}

var app *App

func (app *App) JSONResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"data": data,
	})
}

func (app *App) ErrorResponse(c *gin.Context, statusCode int, err error) {
	c.AbortWithStatusJSON(statusCode, gin.H{
		"Message": string(err.Error()),
	})
}

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		// c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin")
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3000/")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, token, accept, origin, Cache-Control, X-Requested-With, XMLHttpReques")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, PATCH, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

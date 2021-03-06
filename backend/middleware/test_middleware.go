package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// // リクエストを受けた時
		fmt.Println("request")
		c.Next()
		// // レスポンスを返す時
		fmt.Println("respones")
		if err := c.Errors; err != nil {
			// // 実装中
			fmt.Println("[errorMiddleware]", err)
		}
	}
}

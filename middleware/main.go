package middleware

import (
	"github.com/gin-gonic/gin"
)

func TestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// リクエストを受けた時
		c.Next()
		// レスポンスを返す時
		// if err := c.Errors; err != nil {
		// 実装中
		// fmt.Println("[errorMiddleware]", err)
		// }
	}
}

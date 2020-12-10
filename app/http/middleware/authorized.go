package middleware

import (
	"github.com/gin-gonic/gin"
)

//中间件
func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.Set("example", "token不存在")
		} else {
			c.Set("example", "token已存在")
		}
	}
}

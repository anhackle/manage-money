package middlewares

import (
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponseExternal(c, response.ErrTokenInvalid, "")
			c.Abort()
			return
		}
		c.Next()
	}
}

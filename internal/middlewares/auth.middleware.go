package middlewares

import (
	"context"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

var (
	cookieName = "access-token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieValue, err := c.Cookie(cookieName)
		if err != nil {
			response.ErrorResponseNoLogin(c, response.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		email, err := global.Rdb.Get(context.Background(), cookieValue).Result()
		if err != nil {
			response.ErrorResponseNoLogin(c, response.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Set("email", email)

		c.Next()

	}
}

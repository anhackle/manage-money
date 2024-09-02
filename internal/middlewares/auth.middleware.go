package middlewares

import (
	"context"
	"strings"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

var (
	headerName = "Authorization"
)

func AuthMiddleware() gin.HandlerFunc {
	//TODO: Use DI for AuthMiddleware
	return func(c *gin.Context) {
		headerValue := c.GetHeader(headerName)
		if headerValue == "" {
			response.ErrorResponseNoLogin(c, response.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		arrayHeaderValues := strings.Split(headerValue, " ")
		if len(arrayHeaderValues) != 2 || arrayHeaderValues[0] != "Bearer" {
			response.ErrorResponseNoLogin(c, response.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		accessToken := arrayHeaderValues[1]
		userID, err := global.Rdb.Get(context.Background(), accessToken).Int()
		if err != nil {
			response.ErrorResponseNoLogin(c, response.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Set("userID", userID)

		c.Next()

	}
}

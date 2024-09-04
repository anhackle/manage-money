package middlewares

import (
	"strings"

	service "github.com/anle/codebase/internal/services"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

var (
	headerName = "Authorization"
)

type AuthMiddleware struct {
	authService service.IAuthService
}

func (am *AuthMiddleware) Authentication() gin.HandlerFunc {
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
		userID, err := am.authService.Authentication(accessToken)
		if err != nil {
			response.ErrorResponseNoLogin(c, response.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}

func NewAuthMiddleware(authService service.IAuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

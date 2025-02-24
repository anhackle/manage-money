package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleResult(c *gin.Context, result int, data any) {
	switch result {
	case ErrCodeSuccess:
		SuccessResponse(c, result, data)

	case ErrCodeInternal:
		ErrorResponseInternal(c, result, nil)

	case ErrCodeUserHasExists:
		ErrorResponseExternal(c, result, nil)

	case ErrCodeLoginFail:
		ErrorResponseExternal(c, result, nil)
	}
}

// success
func SuccessResponse(c *gin.Context, code int, data any) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponseExternal(c *gin.Context, code int, data any) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

func ErrorResponseInternal(c *gin.Context, code int, data any) {
	c.JSON(http.StatusInternalServerError, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

func ErrorResponseNoLogin(c *gin.Context, code int, data any) {
	c.JSON(http.StatusUnauthorized, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

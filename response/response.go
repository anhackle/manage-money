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

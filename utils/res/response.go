package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

const (
	SUCCESS = 0
	ERROR   = -1
)

func Result(code int, msg string, data any, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func Ok(msg string, data any, c *gin.Context) {
	Result(SUCCESS, msg, data, c)
}

// OkWith 只返回最基本的成功
func OkWith(c *gin.Context) {
	Result(SUCCESS, "success", map[string]string{}, c)
}

func OkWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, msg, map[string]string{}, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, "success", data, c)
}

func Fail(msg string, data any, c *gin.Context) {
	Result(ERROR, msg, data, c)
}

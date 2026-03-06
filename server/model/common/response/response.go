package response

import (
	"errors"
	"hab/code"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 响应结构体
type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
}

// Result 响应结果
func Result(code code.Code, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    int(code),
		Data:    data,
		Message: msg,
	})
}

func Ok(c *gin.Context) {
	Result(code.CodeSuccess, map[string]interface{}{}, "ok", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(code.CodeSuccess, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(code.CodeSuccess, data, "success", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(code.CodeSuccess, data, message, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(code.CodeError, map[string]interface{}{}, message, c)
}

func FailWithMessageData(message string, data interface{}, c *gin.Context) {
	Result(code.CodeError, data, message, c)
}

func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		7,
		nil,
		message,
	})
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(code.CodeError, data, message, c)
}

func FailWithErr(err error, c *gin.Context) {
	var codeMsg code.CodeMsg
	if !errors.As(err, &codeMsg) {
		codeMsg = code.CodeMsg{Code: 7, Msg: err.Error()}
	}
	Result(codeMsg.Code, nil, codeMsg.Msg, c)
}

func FailWithErrData(err error, data interface{}, c *gin.Context) {
	var codeMsg code.CodeMsg
	if !errors.As(err, &codeMsg) {
		codeMsg = code.CodeMsg{Code: 7, Msg: err.Error()}
	}
	Result(codeMsg.Code, data, codeMsg.Msg, c)
}

func FailWithCodeMessage(code code.Code, message string, c *gin.Context) {
	Result(code, nil, message, c)
}

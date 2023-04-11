package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 9
	SUCCESS = 0
)

func Result(httpCode int, code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(httpCode, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(c *gin.Context) {
	Result(http.StatusOK, SUCCESS, map[string]interface{}{}, "操作成功！", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(http.StatusOK, SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(http.StatusOK, SUCCESS, data, "查询成功", c)
}

func Fail(c *gin.Context) {
	Result(http.StatusOK, ERROR, map[string]interface{}{}, "操作失败！", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(http.StatusOK, ERROR, map[string]interface{}{}, message, c)
}

func FailWithDataAndMsg(data interface{}, message string, c *gin.Context) {
	Result(http.StatusOK, ERROR, data, message, c)
}

func FailWithJwtFailed(data interface{}, message string, c *gin.Context) {
	//Jwt鉴权错误，返回400
	Result(http.StatusBadRequest, ERROR, data, message, c)
}

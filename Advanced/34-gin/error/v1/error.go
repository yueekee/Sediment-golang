package error

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS_CODE = 200   //成功的状态码
	FAIL_CODE    = 30000 //失败的状态码
)

// ResponseModel ...
type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseModelBase ...
type ResponseModelBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ret := ResponseModel{Code: SUCCESS_CODE, Message: "ok", Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// ResFail 响应失败
func ResFail(c *gin.Context, msg string) {
	ret := ResponseModelBase{Code: FAIL_CODE, Message: msg}
	ResJSON(c, http.StatusOK, &ret)
}

// ResJSON 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	c.JSON(status, v)
	c.Abort()
}


// Package common
// @Time  : 2021/10/12 上午10:44
// @Author: Jtyoui@qq.com
// @note  : 响应数据类型
package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 所有的响应数据格式
type response struct {
	Code CodeType    `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// Result 通用返回格式
func result(code CodeType, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, response{Code: code, Data: data, Msg: code.getError().Error()})
}

// Err 返回错误信息
func Err(code CodeType, c *gin.Context) {
	result(code, gin.H{}, c)
}

// Success 返回成功信息
func Success(data interface{}, c *gin.Context) {
	result(SuccessCode, data, c)
}

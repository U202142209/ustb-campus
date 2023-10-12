/*
@Author ;@我不是大佬
@Email  ;2869210303@qq.com
@Qq     ;2869210303
@wx     ;safeseaa
@github ;https://github.com/U202142209
@blog   ;https://blog.csdn.net/V123456789987654
@File   ;response.go
@Time   ;2023/10/8 14:17
*/
package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Body struct {
	Code int         `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// Response setting gin.JSON
func Response(c *gin.Context, httpCode int, msg string, data interface{}) {
	if httpCode == http.StatusOK && msg == "" {
		msg = "接口请求成功"
	} else if httpCode == http.StatusBadRequest && msg == "" {
		msg = "接口参数有误"
	} else if httpCode == http.StatusInternalServerError && msg == "" {
		msg = "服务器内部出错"
	}
	c.JSON(httpCode, Body{
		Code: httpCode,
		Msg:  msg,
		Data: data,
	})
}

// 但会json 数据，自带错误处理
func JsonResponse(c *gin.Context, data interface{}, err error) {
	if err != nil {
		fmt.Println("\n\n程序出错了，", err)
		Response(c, http.StatusInternalServerError, "", gin.H{"error": err.Error()})
		return
	}
	Response(c, http.StatusOK, "", data)
}

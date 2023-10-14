/*
@Author ;@我不是大佬
@Email  ;2869210303@qq.com
@Qq     ;2869210303
@wx     ;safeseaa
@github ;https://github.com/U202142209
@blog   ;https://blog.csdn.net/V123456789987654
@File   ;quanzi.go
@Time   ;2023/10/8 14:14
*/
package router

import (
	"BackEnd/spyder"
	"BackEnd/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GettaskbyTypeQuanziHandler(c *gin.Context) {
	// 获取请求的参数
	length := c.Query("length")
	radioGroup := c.Query("radioGroup")
	Type := c.Query("type")
	campus := c.Query("campus")
	util.CheckQuanziQuery(&length, &radioGroup, &Type, &campus) // 检查参数
	search := c.Query("search")
	var (
		data interface{}
		err  error
	)
	if search == "" {
		// 判断是否需要搜索帖子
		data, err = spyder.GettaskbyTypeQuanziSpyder(length, radioGroup, Type, campus)
	} else {
		// 参数不为空，搜索帖子
		data, err = spyder.GettaskbySearchQuanziSpyder(search, length, campus)
	}
	util.JsonResponse(c, data, err)
}

// 帖子详情
func GettaskbyIdQuanziHandler(c *gin.Context) {
	pk := c.Query("pk")
	if pk == "" {
		util.Response(c, http.StatusBadRequest, "", "")
	}
	data, err := spyder.GettaskbyIdQuanziSpyder(pk)
	util.JsonResponse(c, data, err)
}

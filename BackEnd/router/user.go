/*
@Author ;@我不是大佬
@Email  ;2869210303@qq.com
@Qq     ;2869210303
@wx     ;safeseaa
@github ;https://github.com/U202142209
@blog   ;https://blog.csdn.net/V123456789987654
@File   ;user.go
@Time   ;2023/10/8 14:14
*/
package router

import (
	"BackEnd/spyder"
	"BackEnd/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserDataHandler(c *gin.Context) {
	url := c.Param("url") // 获取请求的参数
	var openid string = c.Query("openid")
	var length string = c.Query("length")
	if openid == "" {
		// o7FW15Xeu8-ZRHmmDwfH0ezT9hII
		util.Response(c, http.StatusBadRequest, "", "")
		return
	}
	_, length = util.CheckLength(&length)
	data, err := spyder.GetUserDataSpyder(url, openid, length)
	util.JsonResponse(c, data, err)
}

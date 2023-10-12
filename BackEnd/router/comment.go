/*
@Author ;@我不是大佬
@Email  ;2869210303@qq.com
@Qq     ;2869210303
@wx     ;safeseaa
@github ;https://github.com/U202142209
@blog   ;https://blog.csdn.net/V123456789987654
@File   ;comment.go
@Time   ;2023/10/8 14:14
*/
package router

import (
	"BackEnd/spyder"
	"BackEnd/util"
	"github.com/gin-gonic/gin"
)

// 获取帖子的评论
func GetCommentByTypeQuanziHandler(c *gin.Context) {
	length := c.Query("length")
	pk := c.Query("pk")
	Ttype := c.Query("Ttype")
	data, err := spyder.GetCommentByTypeQuanziSpyder(length, pk, Ttype)
	util.JsonResponse(c, data, err)
}

// 删除评论
func DeleteCommentQuanziHandler(c *gin.Context) {
	pk := c.Param("pk")
	data, err := spyder.DeleteCommentQuanziSpyder(pk)
	util.JsonResponse(c, data, err)
}

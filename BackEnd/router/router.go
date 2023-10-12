/*
@Author ;@我不是大佬
@Email  ;2869210303@qq.com
@Qq     ;2869210303
@wx     ;safeseaa
@github ;https://github.com/U202142209
@blog   ;https://blog.csdn.net/V123456789987654
@File   ;router.go
@Time   ;2023/10/8 13:57
*/
package router

import (
	"BackEnd/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 处理跨域请求,支持options访问
	r.Use(middleware.Cors())

	r.GET("/", IndexHandler)

	rapi := r.Group("/api")
	rapi.GET("/user/:url", GetUserDataHandler) // 获取用户的数据，发布的帖子、评论、回复通知

	rapi.GET("/post/gettaskbyTypeQuanzi/", GettaskbyTypeQuanziHandler) // 获取首页的帖子
	rapi.GET("/post/detail/", GettaskbyIdQuanziHandler)
	// 获取帖子的评论
	rapi.GET("/comment/getCommentByTypeQuanzi/", GetCommentByTypeQuanziHandler)
	// 删除评论
	rapi.GET("/comment/deleteCommentQuanzi/:pk", DeleteCommentQuanziHandler)
	// 搜索帖子
	return r
}

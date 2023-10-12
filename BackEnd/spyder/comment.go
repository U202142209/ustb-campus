/*
@Author ;@我不是大佬
@Email  ;2869210303@qq.com
@Qq     ;2869210303
@wx     ;safeseaa
@github ;https://github.com/U202142209
@blog   ;https://blog.csdn.net/V123456789987654
@File   ;comment.go
@Time   ;2023/10/8 14:38
*/
package spyder

import (
	"BackEnd/config"
	"BackEnd/util"
	"fmt"
)

// 获取帖子的评论
func GetCommentByTypeQuanziSpyder(length, pk, Type string) (interface{}, error) {
	requestUrl := fmt.Sprintf("/getCommentByTypeQuanzi?length=%s&pk=%s&type=%s", length, pk, Type)
	fmt.Println("\n\n" + requestUrl + "\n\n")
	return util.Request(config.RemoteServerAddress_02 + requestUrl)
}

// 删除评论
func DeleteCommentQuanziSpyder(pk string) (interface{}, error) {
	requestUrl := fmt.Sprintf("/deleteCommentQuanzi?pk=%s", pk)
	fmt.Println("\n\n" + requestUrl + "\n\n")
	return util.Request(config.RemoteServerAddress_02 + requestUrl)
}

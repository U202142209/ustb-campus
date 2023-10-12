/*
@Author ;@我不是大佬
@Email  ;2869210303@qq.com
@Qq     ;2869210303
@wx     ;safeseaa
@github ;https://github.com/U202142209
@blog   ;https://blog.csdn.net/V123456789987654
@File   ;post.go
@Time   ;2023/10/8 14:04
*/
package spyder

import (
	"BackEnd/config"
	"BackEnd/util"
	"fmt"
	"net/url"
)

func GettaskbyTypeQuanziSpyder(length, radioGroup, Type, campus string) (interface{}, error) {
	requestUrl := fmt.Sprintf("/gettaskbyTypeQuanzi?length=%s&radioGroup=%s&type=%s&campus=%s", length, radioGroup, Type, campus)
	// fmt.Println("\n\n" + requestUrl + "\n\n")
	return util.Request(config.RemoteServerAddress_02 + requestUrl)
}

// 搜索帖子
func GettaskbySearchQuanziSpyder(search, length, campus string) (interface{}, error) {
	requestUrl := fmt.Sprintf("/gettaskbySearchQuanzi?search=%s&length=%s&campus=%s", url.QueryEscape(search), length, campus)
	fmt.Println("\n\n" + config.RemoteServerAddress_02 + requestUrl + "\n\n")
	return util.Request(config.RemoteServerAddress_02 + requestUrl)
}

// 帖子详情
func GettaskbyIdQuanziSpyder(pk string) (interface{}, error) {
	requestUrl := fmt.Sprintf("/gettaskbyIdQuanzi?pk=%s", pk)
	// fmt.Println("\n\n" + requestUrl + "\n\n")
	return util.Request(config.RemoteServerAddress_02 + requestUrl)
}

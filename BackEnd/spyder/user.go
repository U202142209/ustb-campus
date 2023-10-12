/*
@Author ;@我不是大佬
@Email  ;2869210303@qq.com
@Qq     ;2869210303
@wx     ;safeseaa
@github ;https://github.com/U202142209
@blog   ;https://blog.csdn.net/V123456789987654
@File   ;user.go
@Time   ;2023/10/8 14:39
*/
package spyder

import (
	"BackEnd/config"
	"BackEnd/util"
	"fmt"
)

func GetUserDataSpyder(path, openid, length string) (interface{}, error) {
	requestUrl := fmt.Sprintf("/%s?openid=%s&length=%s&applyTo=%s", path, openid, length, openid)
	fmt.Println("\n\n" + requestUrl + "\n\n")
	return util.Request(config.RemoteServerAddress_02 + requestUrl)
}

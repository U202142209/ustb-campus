/*
@Author ;@我不是大佬
@Email  ;2869210303@qq.com
@Qq     ;2869210303
@wx     ;safeseaa
@github ;https://github.com/U202142209
@blog   ;https://blog.csdn.net/V123456789987654
@File   ;request.go
@Time   ;2023/10/8 16:55
*/
package util

import (
	"fmt"
	"github.com/goccy/go-json"
	"net/http"
)

func Request(requestUrl string) (interface{}, error) {
	resp, err := http.DefaultClient.Get(requestUrl)
	if err != nil {
		fmt.Println("\n\n在发请求的过程中出错了，", err)
		return nil, err
	}
	defer resp.Body.Close()
	var data interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("\n\n在解析json数据的过程中出错了，", err)
		fmt.Println("\n请求体信息：", resp.Body)
		return nil, err
	}
	return data, nil
}

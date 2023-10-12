/*
@Author ;@我不是大佬
@Email  ;2869210303@qq.com
@Qq     ;2869210303
@wx     ;safeseaa
@github ;https://github.com/U202142209
@blog   ;https://blog.csdn.net/V123456789987654
@File   ;main.go
@Time   ;2023/10/8 13:36
*/
package main

import (
	"BackEnd/router"
	"fmt"
)

func main() {
	var ServerPost string = ":8000"
	fmt.Println("localhost" + ServerPost)
	r := router.SetupRouter()
	err := r.Run(ServerPost)
	if err != nil {
		fmt.Println("出戳了：")
		panic(err)
		return
	}
}

/*
@Author ;@我不是大佬
@Email  ;2869210303@qq.com
@Qq     ;2869210303
@wx     ;safeseaa
@github ;https://github.com/U202142209
@blog   ;https://blog.csdn.net/V123456789987654
@File   ;user.go
@Time   ;2023/10/8 13:39
*/

package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id         uint64 `xorm:"pk autoincr"`   // 自增 id
	Username   string `xorm:"varchar(64)"`   // 用户名
	Openid     string `xorm:"varchar(64)"`   // 生成随机的openid
	LoginCount uint64 `xorm:"int default 1"` // 登录次数
	Password   string `xorm:"varchar(100)"`  // 登录密码

	Createtime    time.Time `json:"createtime"`    // 创建时间
	LastLoginTime time.Time `json:"lastlogintime"` // 上次登录时间
}

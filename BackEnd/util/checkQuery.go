/*
@Author ;@我不是大佬
@Email  ;2869210303@qq.com
@Qq     ;2869210303
@wx     ;safeseaa
@github ;https://github.com/U202142209
@blog   ;https://blog.csdn.net/V123456789987654
@File   ;checkQuery.go
@Time   ;2023/10/8 14:50
*/
package util

import (
	"strconv"
)

func CheckQuanziQuery(length, radioGroup, Type, campus *string) bool {
	// 如果参数不合法，则设置默认值
	// 1.查看全部 新发
	var isValid bool = true
	isValid, *length = CheckLength(length)
	if !CheckradioGroup(radioGroup) {
		isValid = false
	}
	if !CheckType(Type) {
		isValid = false
	}
	if campus == nil || *campus == "" {
		*campus = "1"
		isValid = false
	}
	return isValid
}

/*
新发、新回、最热、精选
*/
func CheckType(Type *string) bool {
	if Type != nil {
		switch *Type {
		case "0", "1", "2", "3":
			return true
		}
	}
	*Type = "0"
	return false
}

/*
查看全部:	%5B%22radio4%22%2C%22radio40%22%2C%22radio41%22%2C%22radio42%22%2C%22radio43%22%5D
吐槽		：	%5B%22radio4%22%2C%22radio40%22%5D
表白		：	%5B%22radio41%22%5D
心愿		：	%5B%22radio42%22%5D
知乎		：	%5B%22radio43%22%5D
*/
func CheckradioGroup(radioGroup *string) bool {
	if radioGroup != nil {
		switch *radioGroup {
		case "%5B%22radio4%22%2C%22radio40%22%2C%22radio41%22%2C%22radio42%22%2C%22radio43%22%5D", "%5B%22radio4%22%2C%22radio40%22%5D", "%5B%22radio41%22%5D", "%5B%22radio42%22%5D", "%5B%22radio43%22%5D":
			return true
		}
	}
	*radioGroup = "%5B%22radio4%22%2C%22radio40%22%2C%22radio41%22%2C%22radio42%22%2C%22radio43%22%5D"
	return false
}

/*
判断时候为10的倍数，否则返回10
*/
func CheckLength(lengthStr *string) (bool, string) {
	if lengthStr == nil || *lengthStr == "" {
		return false, "10"
	}
	length, err := strconv.Atoi(*lengthStr)
	if err == nil {
		// 判断length是否为10的倍数
		if length%10 == 0 {
			return true, *lengthStr
		}
	}
	return false, "10"
}

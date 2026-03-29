package utils

import (
	"crypto/md5"
	"fmt"
)

// 给字符串生成md5
// @params str 需要加密的字符串
// @params salt interface{} 加密的盐
// @return str 返回md5码
func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	if len(salt) > 0 {
		saltStr := fmt.Sprint(salt...)
		str = str + saltStr
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

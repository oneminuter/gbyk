package util

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	num = "0123456789"
	str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	tokenKey = []byte("XXXXXXXXXXXXXXXX")
)

/*获取定长随机字符
l 长度,要生成的字符串长度
t 类型,number 纯数字类型, string 数字字母组合类型
*/
func GetRandomString(length int, rtype string) string {
	res := ""
	if rtype == "number" {
		res = num
	} else if rtype == "string" {
		res = str
	} else {
		return ""
	}

	bytes := []byte(res)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/*
将手机号中间4为用*代替
*/
func ReplaceMobile(mobile string) string {
	rs := []rune(mobile)
	return fmt.Sprintf("%s****%s", string(rs[0:3]), string(rs[7:]))
}

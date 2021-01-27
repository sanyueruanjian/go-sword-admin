package utils

import (
	"regexp"
	"time"
)

// 获取当前时间
func NowTime() string {
	return time.Unix(time.Now().Unix(), 0,).Format("2006-01-02 15:04:05")
}

// 校验手机号
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}


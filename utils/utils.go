package utils

import (
	"encoding/json"
	"regexp"
	"strconv"
	"time"

	"project/app/admin/models/bo"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}

// Assert 条件断言
// 当断言条件为 假 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
func Assert(condition bool, msg string, code ...int) {
	if !condition {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}

// HasError 错误断言
// 当 error 不为 nil 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
// 若 msg 为空，则默认为 error 中的内容
func HasError(err error, msg string, code ...int) {
	if err != nil {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		if msg == "" {
			msg = err.Error()
		}
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}

func OrderJson(orders string) (orderData []bo.Order, err error) {
	if orders == "" {
		orders = `[{"column": "create_time", "asc": "false"}]`
	}
	err = json.Unmarshal([]byte(orders), &orderData)
	return
}

func GetOrderRule(orderData []bo.Order) (orderRule string) {
	orderRule = ""
	if orderData == nil || len(orderData) == 0 {
		return
	}
	for index, v := range orderData {
		switch index {
		case len(orderData) - 1:
			if v.Asc == "true" {
				orderRule += v.Column + " asc"
			} else {
				orderRule += v.Column + " desc"
			}
		default:
			if v.Asc == "true" {
				orderRule += v.Column + " asc, "
			} else {
				orderRule += v.Column + " desc, "
			}
		}
	}
	return
}

func TimeToString(time time.Time) string {
	var timeLayoutStr = "2006-01-02 15:04:05"
	return time.Format(timeLayoutStr)
}

func UnixTimeToString(t int64) string {
	tTime := time.Unix(t/1000, 0)
	return TimeToString(tTime)
}

func StringToTime(t string) time.Time {
	var timeLayoutStr = "2006-01-02 15:04:05"
	tm, err := time.Parse(timeLayoutStr, t)
	if err != nil {
		zap.L().Error("StringToTime failed", zap.Error(err))
		return time.Time{}
	}
	return tm
}

func PagesCount(count int, size int) int {
	return (count-1)/size + 1
}

func BlurryCache(blurry string, sourceStr string) bool {
	matched, _ := regexp.MatchString(blurry, sourceStr)
	return matched
}

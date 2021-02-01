package utils

import (
	"encoding/json"
	"strconv"
	"time"
)

func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 时间戳转时间
func UnixToTime(e string) (datatime time.Time, err error) {
	data, err := strconv.ParseInt(e, 10, 64)
	datatime = time.Unix(data/1000, 0)
	return
}

// 时间转时间戳
func TimeToUnix(e time.Time) int64 {
	timeUnix, _ := time.Parse("2006-01-02 15:04:05", e.Format("2006-01-02 15:04:05"))
	return timeUnix.UnixNano() / 1e6
}

func GetCurrentTimeUnix() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

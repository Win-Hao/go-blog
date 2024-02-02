package utils

import (
	"strconv"
	"time"
)

func GetDay() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func Int(str string) (int, error) {
	num, err := strconv.Atoi(str)
	return num, err
}
func Str(num int) string {
	str := strconv.Itoa(num)
	return str
}

func GetUnix() int {
	now := time.Now()
	return int(now.Unix())
}

package util

import (
	"strconv"
)

// StrToInt 字符串转整型
func StrToInt(value string) int {
	val, _ := strconv.Atoi(value)
	return val
}

// IntToStr 整型转字符串
func IntToStr(value int) string {
	return strconv.Itoa(value)
}

// StrToDouble 字符串转浮点
func StrToDouble(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

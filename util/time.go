package util

import (
	"time"
)

const (
	// SecondTimeLayout 秒的时间布局
	SecondTimeLayout = "2006-01-02 15:04:05"
	// MinuteTimeLayout 分的时间布局
	MinuteTimeLayout = "2006-01-02 15:04"
	// StdDateLayout 日期的时间布局
	StdDateLayout = "2006-01-02"
	// DuaDateLayout 日期的时间布局
	DuaDateLayout = "20060102"
)

// FormatTime 格式化时间字符串
func FormatTime(t time.Time, layout string) string {
	if !t.IsZero() {
		res := t.Format(layout)
		if res == "0001-01-01 00:00:00" || res == "0001-01-01" {
			return ""
		}
		return res
	}
	return ""
}

func FromLayout(t string, layout string) (time.Time, error) {
	return time.Parse(layout, t)
}

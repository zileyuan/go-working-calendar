package service

import (
	"strconv"
	"time"

	"github.com/zileyuan/go-working-calendar/config"
	"github.com/zileyuan/go-working-calendar/response"
	"github.com/zileyuan/go-working-calendar/util"
)

func Holiday(date string) (bool, response.StatusType) {
	t, err := util.FromLayout(date, util.DuaDateLayout)
	if err == nil {
		if util.Contains(config.CalendarData.Maintain, util.IntToStr(t.Year())) {
			if util.Contains(config.CalendarData.Working, date) {
				return false, response.StatusOK
			} else if util.Contains(config.CalendarData.Holiday, date) {
				return true, response.StatusOK
			} else if t.Weekday() == time.Sunday || t.Weekday() == time.Saturday {
				return true, response.StatusOK
			} else {
				return false, response.StatusOK
			}
		} else {
			return false, response.StatusMaintainError
		}
	} else {
		return false, response.StatusParamsError
	}

}

func Count(from, to string) (int, response.StatusType) {
	t1, err1 := util.FromLayout(from, util.DuaDateLayout)
	t2, err2 := util.FromLayout(to, util.DuaDateLayout)
	if err1 == nil && err2 == nil && !t2.Before(t1) {
		tf := t1
		length := 0
		for {
			if t2.Before(tf) {
				break
			}
			res, status := Holiday(util.FormatTime(tf, util.DuaDateLayout))
			if status != response.StatusOK {
				return 0, status
			}
			if !res {
				length++
			}
			tf = tf.AddDate(0, 0, 1)
		}
		return length, response.StatusOK
	} else {
		return 0, response.StatusParamsError
	}
}

func Calc(from, amount string) (string, response.StatusType) {
	t, err1 := util.FromLayout(from, util.DuaDateLayout)
	a, err2 := strconv.Atoi(amount)
	if err1 == nil && err2 == nil && a != 0 {
		tf := t
		if a > 0 {
			length := a - 1
			for {
				if length == 0 {
					break
				}
				res, status := Holiday(util.FormatTime(tf, util.DuaDateLayout))
				if status != response.StatusOK {
					return "", status
				}
				if !res {
					length--
				}
				tf = tf.AddDate(0, 0, 1)
			}
			return util.FormatTime(tf, util.DuaDateLayout), response.StatusOK
		} else {
			length := a
			for {
				if length == 0 {
					break
				}
				res, status := Holiday(util.FormatTime(tf, util.DuaDateLayout))
				if status != response.StatusOK {
					return "", status
				}
				if !res {
					length++
				}
				tf = tf.AddDate(0, 0, -1)
			}
			return util.FormatTime(tf, util.DuaDateLayout), response.StatusOK
		}
	} else {
		return "", response.StatusParamsError
	}
}

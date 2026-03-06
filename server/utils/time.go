package utils

import (
	"fmt"
	"strconv"
	"time"
)

func NowTime() *time.Time {
	now := time.Now()
	return &now
}

func NowTimeAdd(d time.Duration) *time.Time {
	now := time.Now()
	if d != 0 {
		now = now.Add(d)
	}
	return &now
}

// NowTimeRange 获取当天的时间范围
func NowTimeRange() []interface{} {
	startOfDay := time.Now().Truncate(24 * time.Hour) // 当天 00:00:00
	endOfDay := startOfDay.Add(24 * time.Hour)        // 次日 00:00:00
	return []interface{}{startOfDay, endOfDay}
}

// IntDate 获取日期的整数格式 20201201
func IntDate(t time.Time) uint {
	i, _ := strconv.Atoi(t.Format("20060102"))
	return uint(i)
}

// DateInt converts an integer in the format (YYYYMMDD) to a time.Time.
func DateInt(i uint) time.Time {
	dateStr := fmt.Sprintf("%08d", i)
	t, _ := time.Parse("20060102", dateStr)
	return t
}

// DaysInMonth 获取某个月份的天数
func DaysInMonth(t time.Time) int {
	// 获取当前年和月
	year, month := t.Year(), t.Month()
	// 获取下一个月的第 1 天
	nextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	// 计算当前月份的天数
	return nextMonth.AddDate(0, 0, -1).Day()
}

// GetTimeWeek 计算当周的第一天（星期一）
func GetTimeWeek(t time.Time) time.Time {
	// 1. 获取当前时间是星期几
	weekday := t.Weekday()

	// 2. 将星期天视为一周的第 7 天
	if weekday == time.Sunday {
		weekday = 7
	}

	// 3. 计算星期一到现在的偏移量
	offset := int(weekday - time.Monday)

	// 4. 获取当周的第一天
	weekStart := t.AddDate(0, 0, -offset)

	// 5. 将时间设置为 0 点（即当天的开始时间）
	weekStart = time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, weekStart.Location())
	return weekStart
}

// GetTimeMonth 计算当月的第一天（1 号）
func GetTimeMonth(t time.Time) time.Time {
	monthStart := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	return monthStart
}

// GetTimeDay 计算当天的起始时间（0 点）
func GetTimeDay(t time.Time) time.Time {
	dayStart := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return dayStart
}

// GetTimeHour 计算当前小时的起始时间（0 分）
func GetTimeHour(t time.Time) time.Time {
	hourStart := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	return hourStart
}

func TimeOffset(t *time.Time) string {
	// 获取当前时间
	now := time.Now()

	// 获取当前时区
	loc := now.Location()

	// 获取当前时区相对于 UTC 的偏移量（以秒为单位）
	_, offset := now.In(loc).Zone()

	newTime := t.Add(time.Duration(offset) * time.Second)
	return newTime.Format("2006-01-02 15:04:05")
}

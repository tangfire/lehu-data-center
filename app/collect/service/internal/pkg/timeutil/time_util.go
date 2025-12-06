// internal/pkg/timeutil/time_util.go
package timeutil

import (
	"time"
)

var (
	// 时间格式常量
	DateFormat         = "2006-01-02"
	DateFormatZhCN     = "01月02日"
	MonthFormat        = "2006-01"
	YearFormat         = "2006"
	DateTimeFormat     = "2006-01-02 15:04:05"
	DateFormatMMDDZhCN = "01月02日"
)

// 获取指定日期的开始时间（00:00:00）
func ObtainStartDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

// 获取指定日期的结束时间（23:59:59.999999999）
func ObtainEndDay(date time.Time) time.Time {
	start := ObtainStartDay(date)
	return start.AddDate(0, 0, 1).Add(-time.Nanosecond)
}

// 获取指定日期所在周的开始时间（周一 00:00:00）
func ObtainStartWeek(date time.Time) time.Time {
	// Go的Weekday周日=0，周一=1
	offset := int(date.Weekday()) - 1 // 计算距离周一的偏移
	if offset < 0 {
		offset = 6 // 如果是周日，偏移6天
	}
	startOfDay := ObtainStartDay(date)
	return startOfDay.AddDate(0, 0, -offset)
}

// 获取指定日期所在周的结束时间（周日 23:59:59.999999999）
func ObtainEndWeek(date time.Time) time.Time {
	startOfWeek := ObtainStartWeek(date)
	return startOfWeek.AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// 获取指定日期所在月的开始时间
func ObtainStartMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
}

// 获取指定日期所在月的结束时间
func ObtainEndMonth(date time.Time) time.Time {
	startOfMonth := ObtainStartMonth(date)
	nextMonth := startOfMonth.AddDate(0, 1, 0)
	return nextMonth.Add(-time.Nanosecond)
}

// 获取指定日期所在年的开始时间
func ObtainStartYear(date time.Time) time.Time {
	return time.Date(date.Year(), 1, 1, 0, 0, 0, 0, date.Location())
}

// 获取指定日期所在年的结束时间
func ObtainEndYear(date time.Time) time.Time {
	startOfYear := ObtainStartYear(date)
	nextYear := startOfYear.AddDate(1, 0, 0)
	return nextYear.Add(-time.Nanosecond)
}

// Format 格式化时间
func Format(t time.Time, layout string) string {
	return t.Format(layout)
}

// AddDay 增加天数
func AddDay(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, days)
}

// AddWeek 增加周数
func AddWeek(date time.Time, weeks int) time.Time {
	return date.AddDate(0, 0, weeks*7)
}

// Now 获取当前时间
func Now() time.Time {
	return time.Now()
}

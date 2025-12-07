// internal/pkg/datetime/datetime_func.go
package funcs

import (
	"errors"
	"fmt"
	"lehu-data-center/app/collect/service/internal/enums"
	"lehu-data-center/app/collect/service/internal/pkg/timeutil"
	"lehu-data-center/app/collect/service/internal/types/transfers"
	"time"
)

var (
	// 常量
	Strike = "-"

	// SQL参数常量
	StartTimeParam    = "startTime"
	EndTimeParam      = "endTime"
	DatePattern       = timeutil.DateFormat
	DateFormatPattern = timeutil.DateTimeFormat
)

// CreateRequestTime 根据日期类型和日期创建请求时间范围
func CreateRequestTime(dateType enums.DateType, date time.Time) (*transfers.RequestTime, error) {
	requestTime := &transfers.RequestTime{
		GatherDate: date,
		DateType:   dateType,
	}

	var startTime, endTime time.Time

	switch dateType {
	case enums.DateTypeDay:
		startTime = timeutil.ObtainStartDay(date)
		endTime = timeutil.ObtainEndDay(date)
	case enums.DateTypeWeek:
		startTime = timeutil.ObtainStartWeek(date)
		endTime = timeutil.ObtainEndWeek(date)
	case enums.DateTypeMonth:
		startTime = timeutil.ObtainStartMonth(date)
		endTime = timeutil.ObtainEndMonth(date)
	case enums.DateTypeYear:
		startTime = timeutil.ObtainStartYear(date)
		endTime = timeutil.ObtainEndYear(date)
	default:
		return nil, errors.New(fmt.Sprintf("不支持的日期类型: %v", dateType))
	}

	requestTime.StartTime = &startTime
	requestTime.EndTime = &endTime

	return requestTime, nil
}

// CreateTimeParam 创建时间参数Map
func CreateTimeParam(startTime, endTime time.Time, periodFormat string) map[string]interface{} {
	return map[string]interface{}{
		StartTimeParam: timeutil.Format(startTime, periodFormat),
		EndTimeParam:   timeutil.Format(endTime, periodFormat),
	}
}

// UpgradeDateType 升级时间类型
func UpgradeDateType(dateType enums.DateType) enums.DateType {
	switch dateType {
	case enums.DateTypeDay:
		return enums.DateTypeWeek
	case enums.DateTypeWeek:
		return enums.DateTypeMonth
	case enums.DateTypeMonth:
		return enums.DateTypeYear
	default:
		return 0 // 返回零值，表示不支持
	}
}

// CreateCountTime 创建统计时间字符串
func CreateCountTime(dateType enums.DateType, startTime, endTime time.Time) string {
	switch dateType {
	case enums.DateTypeDay:
		return timeutil.Format(startTime, DatePattern)
	case enums.DateTypeWeek, enums.DateTypeMonth, enums.DateTypeYear:
		return timeutil.Format(startTime, DatePattern) + Strike + timeutil.Format(endTime, DatePattern)
	default:
		return ""
	}
}

// AssemblyCountTime 组装统计时间字符串
func AssemblyCountTime(dateType enums.DateType, startTime, endTime time.Time) string {
	// 逻辑和CreateCountTime相同
	return CreateCountTime(dateType, startTime, endTime)
}

// GetEntryName 生成条目名称
func GetEntryName(dateType enums.DateType, startTime, endTime time.Time) string {
	if dateType == 0 || startTime.IsZero() || endTime.IsZero() {
		return ""
	}

	switch dateType {
	case enums.DateTypeDay:
		return timeutil.Format(startTime, timeutil.DateFormatMMDDZhCN)
	case enums.DateTypeWeek:
		return timeutil.Format(startTime, "01月02日") + "-" + timeutil.Format(endTime, "02日")
	case enums.DateTypeMonth:
		return timeutil.Format(startTime, "06年01月")
	case enums.DateTypeYear:
		year := startTime.Year()
		yearStr := fmt.Sprintf("%d", year)
		if len(yearStr) > 2 {
			yearStr = yearStr[2:] // 取后两位
		}
		return yearStr + "年"
	default:
		return ""
	}
}

// GetYearString 获取年份字符串
func GetYearString(date time.Time) string {
	year := date.Year()
	yearStr := fmt.Sprintf("%d", year)
	if len(yearStr) > 2 {
		yearStr = yearStr[2:] // 取后两位
	}
	return yearStr
}

// ParseDateStr 解析日期字符串
func ParseDateStr(dateStr string) (time.Time, error) {
	// 尝试多种格式
	layouts := []string{
		"2006-01-02",
		"2006/01/02",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		time.RFC3339,
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, dateStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, errors.New("无法解析日期字符串: " + dateStr)
}

// IsSameDay 判断两个时间是否在同一天
func IsSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// GetWeekRange 获取周的日期范围
func GetWeekRange(date time.Time) (start, end time.Time) {
	start = timeutil.ObtainStartWeek(date)
	end = timeutil.ObtainEndWeek(date)
	return
}

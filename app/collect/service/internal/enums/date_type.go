// internal/pkg/enums/date_type.go
package enums

// DateType 日期类型
type DateType int

const (
	DateTypeDay   DateType = 1 // 天
	DateTypeWeek  DateType = 2 // 周
	DateTypeMonth DateType = 3 // 月
	DateTypeYear  DateType = 4 // 年
)

// Code 返回枚举的代码值
func (d DateType) Code() int {
	return int(d)
}

// Value 返回枚举的字符串值
func (d DateType) Value() string {
	switch d {
	case DateTypeDay:
		return "day"
	case DateTypeWeek:
		return "week"
	case DateTypeMonth:
		return "month"
	case DateTypeYear:
		return "year"
	default:
		return ""
	}
}

// Msg 返回枚举的描述信息
func (d DateType) Msg() string {
	switch d {
	case DateTypeDay:
		return "天"
	case DateTypeWeek:
		return "周"
	case DateTypeMonth:
		return "月"
	case DateTypeYear:
		return "年"
	default:
		return ""
	}
}

// GetDateTypeMsg 根据code获取描述信息
func GetDateTypeMsg(code int) string {
	return GetDateTypeByCode(code).Msg()
}

// GetDateTypeByCode 根据code获取枚举
func GetDateTypeByCode(code int) DateType {
	switch code {
	case 1:
		return DateTypeDay
	case 2:
		return DateTypeWeek
	case 3:
		return DateTypeMonth
	case 4:
		return DateTypeYear
	default:
		return 0
	}
}

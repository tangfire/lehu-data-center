package enums

type DateType int

const (
	DateTypeDay   DateType = 1 // 天
	DateTypeWeek  DateType = 2 // 周
	DateTypeMonth DateType = 3 // 月
	DateTypeYear  DateType = 4 // 年
)

func (dt DateType) String() string {
	switch dt {
	case DateTypeDay:
		return "DAY"
	case DateTypeWeek:
		return "WEEK"
	case DateTypeMonth:
		return "MONTH"
	case DateTypeYear:
		return "YEAR"
	default:
		return "UNKNOWN"
	}
}

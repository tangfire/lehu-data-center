// internal/pkg/enums/metric_type.go
package enums

// MetricType 通用状态枚举
type MetricType int

const (
	MetricTypeGather    MetricType = 1 // 采集
	MetricTypeCalculate MetricType = 2 // 计算
)

// Code 返回枚举的代码值
func (m MetricType) Code() int {
	return int(m)
}

// Msg 返回枚举的描述信息
func (m MetricType) Msg() string {
	switch m {
	case MetricTypeGather:
		return "采集"
	case MetricTypeCalculate:
		return "计算"
	default:
		return ""
	}
}

// GetMetricTypeMsg 根据code获取描述信息
func GetMetricTypeMsg(code int) string {
	return GetMetricTypeByCode(code).Msg()
}

// GetMetricTypeByCode 根据code获取枚举
func GetMetricTypeByCode(code int) MetricType {
	switch code {
	case 1:
		return MetricTypeGather
	case 2:
		return MetricTypeCalculate
	default:
		return 0
	}
}

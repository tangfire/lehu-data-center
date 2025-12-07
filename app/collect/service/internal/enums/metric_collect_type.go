// internal/pkg/enums/metric_collect_type.go
package enums

// MetricCollectType 指标收集的类型
type MetricCollectType int

const (
	MetricCollectTypeBase    MetricCollectType = 1 // 基础型
	MetricCollectTypeCompute MetricCollectType = 2 // 计算型
)

// Code 返回枚举的代码值
func (m MetricCollectType) Code() int {
	return int(m)
}

// Msg 返回枚举的描述信息
func (m MetricCollectType) Msg() string {
	switch m {
	case MetricCollectTypeBase:
		return "基础型"
	case MetricCollectTypeCompute:
		return "计算型"
	default:
		return ""
	}
}

// GetMetricCollectTypeMsg 根据code获取描述信息
func GetMetricCollectTypeMsg(code int) string {
	return GetMetricCollectTypeByCode(code).Msg()
}

// GetMetricCollectTypeByCode 根据code获取枚举
func GetMetricCollectTypeByCode(code int) MetricCollectType {
	switch code {
	case 1:
		return MetricCollectTypeBase
	case 2:
		return MetricCollectTypeCompute
	default:
		return 0
	}
}

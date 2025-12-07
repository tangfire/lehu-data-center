// internal/pkg/enums/collect_type.go
package enums

// CollectType 数据采集类型
type CollectType int

const (
	CollectTypeSQL   CollectType = 1 // sql采集
	CollectTypeHTTP  CollectType = 2 // http采集
	CollectTypeRedis CollectType = 3 // redis采集
)

// Code 返回枚举的代码值
func (c CollectType) Code() int {
	return int(c)
}

// Msg 返回枚举的描述信息
func (c CollectType) Msg() string {
	switch c {
	case CollectTypeSQL:
		return "sql采集"
	case CollectTypeHTTP:
		return "http采集"
	case CollectTypeRedis:
		return "redis采集"
	default:
		return ""
	}
}

// GetCollectTypeMsg 根据code获取描述信息
func GetCollectTypeMsg(code int) string {
	return GetCollectTypeByCode(code).Msg()
}

// GetCollectTypeByCode 根据code获取枚举
func GetCollectTypeByCode(code int) CollectType {
	switch code {
	case 1:
		return CollectTypeSQL
	case 2:
		return CollectTypeHTTP
	case 3:
		return CollectTypeRedis
	default:
		return 0
	}
}

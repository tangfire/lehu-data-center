// internal/pkg/enums/type_level.go
package enums

// TypeLevel 当前指标是否和维度有关
type TypeLevel int

const (
	TypeLevelYes TypeLevel = 1 // 是
	TypeLevelNo  TypeLevel = 0 // 否
)

// Code 返回枚举的代码值
func (t TypeLevel) Code() int {
	return int(t)
}

// Msg 返回枚举的描述信息
func (t TypeLevel) Msg() string {
	switch t {
	case TypeLevelYes:
		return "是"
	case TypeLevelNo:
		return "否"
	default:
		return ""
	}
}

// GetTypeLevelMsg 根据code获取描述信息
func GetTypeLevelMsg(code int) string {
	return GetTypeLevelByCode(code).Msg()
}

// GetTypeLevelByCode 根据code获取枚举
func GetTypeLevelByCode(code int) TypeLevel {
	switch code {
	case 1:
		return TypeLevelYes
	case 0:
		return TypeLevelNo
	default:
		return 0
	}
}

// internal/pkg/enums/default_show.go
package enums

// DefaultShow 通用状态枚举
type DefaultShow int

const (
	DefaultShowShow DefaultShow = 1 // 显示
	DefaultShowHide DefaultShow = 0 // 隐藏
)

// Code 返回枚举的代码值
func (d DefaultShow) Code() int {
	return int(d)
}

// Msg 返回枚举的描述信息
func (d DefaultShow) Msg() string {
	switch d {
	case DefaultShowShow:
		return "显示"
	case DefaultShowHide:
		return "隐藏"
	default:
		return ""
	}
}

// GetDefaultShowMsg 根据code获取描述信息
func GetDefaultShowMsg(code int) string {
	return GetDefaultShowByCode(code).Msg()
}

// GetDefaultShowByCode 根据code获取枚举
func GetDefaultShowByCode(code int) DefaultShow {
	switch code {
	case 1:
		return DefaultShowShow
	case 0:
		return DefaultShowHide
	default:
		return 0
	}
}

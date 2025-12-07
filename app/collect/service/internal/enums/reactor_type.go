// internal/pkg/enums/reactor_type.go
package enums

// ReactorType 视频反应
type ReactorType int

const (
	ReactorTypeNormal  ReactorType = 0 // 无反应
	ReactorTypeLike    ReactorType = 1 // 点赞
	ReactorTypeDislike ReactorType = 2 // 点踩
)

// Code 返回枚举的代码值
func (r ReactorType) Code() int {
	return int(r)
}

// Msg 返回枚举的描述信息
func (r ReactorType) Msg() string {
	switch r {
	case ReactorTypeNormal:
		return "无反应"
	case ReactorTypeLike:
		return "点赞"
	case ReactorTypeDislike:
		return "点踩"
	default:
		return ""
	}
}

// GetReactorTypeMsg 根据code获取描述信息
func GetReactorTypeMsg(code int) string {
	return GetReactorTypeByCode(code).Msg()
}

// GetReactorTypeByCode 根据code获取枚举
func GetReactorTypeByCode(code int) ReactorType {
	switch code {
	case 0:
		return ReactorTypeNormal
	case 1:
		return ReactorTypeLike
	case 2:
		return ReactorTypeDislike
	default:
		return 0
	}
}

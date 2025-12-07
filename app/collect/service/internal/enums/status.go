// internal/pkg/enums/status.go
package enums

// Status 状态
type Status int

const (
	StatusRun  Status = 1 // 正常
	StatusStop Status = 0 // 禁用
)

// Code 返回枚举的代码值
func (s Status) Code() int {
	return int(s)
}

// Msg 返回枚举的描述信息
func (s Status) Msg() string {
	switch s {
	case StatusRun:
		return "正常"
	case StatusStop:
		return "禁用"
	default:
		return ""
	}
}

// GetStatusMsg 根据code获取描述信息
func GetStatusMsg(code int) string {
	return GetStatusByCode(code).Msg()
}

// GetStatusByCode 根据code获取枚举
func GetStatusByCode(code int) Status {
	switch code {
	case 1:
		return StatusRun
	case 0:
		return StatusStop
	default:
		return 0
	}
}

// internal/pkg/enums/show_status.go
package enums

// ShowStatus 显示状态枚举
type ShowStatus int

const (
	ShowStatusYes ShowStatus = 1 // 是
	ShowStatusNo  ShowStatus = 0 // 否
)

// Code 返回枚举的代码值
func (s ShowStatus) Code() int {
	return int(s)
}

// Msg 返回枚举的描述信息
func (s ShowStatus) Msg() string {
	switch s {
	case ShowStatusYes:
		return "是"
	case ShowStatusNo:
		return "否"
	default:
		return ""
	}
}

// GetShowStatusMsg 根据code获取描述信息
func GetShowStatusMsg(code int) string {
	return GetShowStatusByCode(code).Msg()
}

// GetShowStatusByCode 根据code获取枚举
func GetShowStatusByCode(code int) ShowStatus {
	switch code {
	case 1:
		return ShowStatusYes
	case 0:
		return ShowStatusNo
	default:
		return 0
	}
}

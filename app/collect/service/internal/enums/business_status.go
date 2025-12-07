// internal/pkg/enums/business_status.go
package enums

// BusinessStatus 通用状态枚举
type BusinessStatus int

const (
	BusinessStatusYes BusinessStatus = 1 // 是
	BusinessStatusNo  BusinessStatus = 0 // 否
)

// Code 返回枚举的代码值
func (b BusinessStatus) Code() int {
	return int(b)
}

// Msg 返回枚举的描述信息
func (b BusinessStatus) Msg() string {
	switch b {
	case BusinessStatusYes:
		return "是"
	case BusinessStatusNo:
		return "否"
	default:
		return ""
	}
}

// GetBusinessStatusMsg 根据code获取描述信息
func GetBusinessStatusMsg(code int) string {
	return GetBusinessStatusByCode(code).Msg()
}

// GetBusinessStatusByCode 根据code获取枚举
func GetBusinessStatusByCode(code int) BusinessStatus {
	switch code {
	case 1:
		return BusinessStatusYes
	case 0:
		return BusinessStatusNo
	default:
		return 0
	}
}

// internal/pkg/enums/phase.go
package enums

// Phase 当前指标是否和统计周期有关
type Phase int

const (
	PhaseYes Phase = 1 // 是
	PhaseNo  Phase = 0 // 否
)

// Code 返回枚举的代码值
func (p Phase) Code() int {
	return int(p)
}

// Msg 返回枚举的描述信息
func (p Phase) Msg() string {
	switch p {
	case PhaseYes:
		return "是"
	case PhaseNo:
		return "否"
	default:
		return ""
	}
}

// GetPhaseMsg 根据code获取描述信息
func GetPhaseMsg(code int) string {
	return GetPhaseByCode(code).Msg()
}

// GetPhaseByCode 根据code获取枚举
func GetPhaseByCode(code int) Phase {
	switch code {
	case 1:
		return PhaseYes
	case 0:
		return PhaseNo
	default:
		return 0
	}
}

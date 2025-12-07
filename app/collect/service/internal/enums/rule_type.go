// internal/pkg/enums/rule_type.go
package enums

// RuleType 规则类型
type RuleType int

const (
	RuleTypeGather RuleType = 1 // 采集
	RuleTypeQuery  RuleType = 2 // 查询
)

// Code 返回枚举的代码值
func (r RuleType) Code() int {
	return int(r)
}

// Msg 返回枚举的描述信息
func (r RuleType) Msg() string {
	switch r {
	case RuleTypeGather:
		return "采集"
	case RuleTypeQuery:
		return "查询"
	default:
		return ""
	}
}

// GetRuleTypeMsg 根据code获取描述信息
func GetRuleTypeMsg(code int) string {
	return GetRuleTypeByCode(code).Msg()
}

// GetRuleTypeByCode 根据code获取枚举
func GetRuleTypeByCode(code int) RuleType {
	switch code {
	case 1:
		return RuleTypeGather
	case 2:
		return RuleTypeQuery
	default:
		return 0
	}
}

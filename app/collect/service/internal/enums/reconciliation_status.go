// internal/pkg/enums/reconciliation_status.go
package enums

// ReconciliationStatus 对账状态
type ReconciliationStatus int

const (
	ReconciliationStatusNo      ReconciliationStatus = 1  // 未对账
	ReconciliationStatusFail    ReconciliationStatus = -1 // 对账完成有问题
	ReconciliationStatusSuccess ReconciliationStatus = 2  // 对账完成没有问题
	ReconciliationStatusFinish  ReconciliationStatus = 3  // 对账有问题处理完毕
)

// Code 返回枚举的代码值
func (r ReconciliationStatus) Code() int {
	return int(r)
}

// Msg 返回枚举的描述信息
func (r ReconciliationStatus) Msg() string {
	switch r {
	case ReconciliationStatusNo:
		return "未对账"
	case ReconciliationStatusFail:
		return "对账完成有问题"
	case ReconciliationStatusSuccess:
		return "对账完成没有问题"
	case ReconciliationStatusFinish:
		return "对账有问题处理完毕"
	default:
		return ""
	}
}

// GetReconciliationStatusMsg 根据code获取描述信息
func GetReconciliationStatusMsg(code int) string {
	return GetReconciliationStatusByCode(code).Msg()
}

// GetReconciliationStatusByCode 根据code获取枚举
func GetReconciliationStatusByCode(code int) ReconciliationStatus {
	switch code {
	case 1:
		return ReconciliationStatusNo
	case -1:
		return ReconciliationStatusFail
	case 2:
		return ReconciliationStatusSuccess
	case 3:
		return ReconciliationStatusFinish
	default:
		return 0
	}
}

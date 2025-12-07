// internal/pkg/enums/function_type.go
package enums

// FunctionType 指标统计函数枚举
type FunctionType int

const (
	FunctionTypeSum       FunctionType = 1 // sum - 累加
	FunctionTypeDeduction FunctionType = 2 // deduction - 相减
	FunctionTypeMultiply  FunctionType = 3 // multiply - 相乘
	FunctionTypeRatio     FunctionType = 4 // ratio - 相除
	FunctionTypeAvg       FunctionType = 5 // avg - 平均数
)

// Code 返回枚举的代码值
func (f FunctionType) Code() int {
	return int(f)
}

// Value 返回枚举的字符串值
func (f FunctionType) Value() string {
	switch f {
	case FunctionTypeSum:
		return "sum"
	case FunctionTypeDeduction:
		return "deduction"
	case FunctionTypeMultiply:
		return "multiply"
	case FunctionTypeRatio:
		return "ratio"
	case FunctionTypeAvg:
		return "avg"
	default:
		return ""
	}
}

// Msg 返回枚举的描述信息
func (f FunctionType) Msg() string {
	switch f {
	case FunctionTypeSum:
		return "累加"
	case FunctionTypeDeduction:
		return "相减"
	case FunctionTypeMultiply:
		return "相乘"
	case FunctionTypeRatio:
		return "相除"
	case FunctionTypeAvg:
		return "平均数"
	default:
		return ""
	}
}

// GetFunctionTypeMsg 根据code获取描述信息
func GetFunctionTypeMsg(code int) string {
	return GetFunctionTypeByCode(code).Msg()
}

// GetFunctionTypeByCode 根据code获取枚举
func GetFunctionTypeByCode(code int) FunctionType {
	switch code {
	case 1:
		return FunctionTypeSum
	case 2:
		return FunctionTypeDeduction
	case 3:
		return FunctionTypeMultiply
	case 4:
		return FunctionTypeRatio
	case 5:
		return FunctionTypeAvg
	default:
		return 0
	}
}

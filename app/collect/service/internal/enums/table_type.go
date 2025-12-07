// internal/pkg/enums/table_type.go
package enums

// TableType 查询的类型
type TableType int

const (
	TableTypeTable TableType = 1 // 报表
	TableTypeChart TableType = 2 // 图表
)

// Code 返回枚举的代码值
func (t TableType) Code() int {
	return int(t)
}

// Msg 返回枚举的描述信息
func (t TableType) Msg() string {
	switch t {
	case TableTypeTable:
		return "报表"
	case TableTypeChart:
		return "图表"
	default:
		return ""
	}
}

// GetTableTypeMsg 根据code获取描述信息
func GetTableTypeMsg(code int) string {
	return GetTableTypeByCode(code).Msg()
}

// GetTableTypeByCode 根据code获取枚举
func GetTableTypeByCode(code int) TableType {
	switch code {
	case 1:
		return TableTypeTable
	case 2:
		return TableTypeChart
	default:
		return 0
	}
}

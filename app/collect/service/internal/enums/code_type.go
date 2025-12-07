// internal/pkg/enums/code_type.go
package enums

import (
	"time"
)

// CodeType 数据详细分类
type CodeType int

const (
	CodeTypeInteger   CodeType = 1 // INTEGER - 整型数字
	CodeTypeLong      CodeType = 2 // LONG - 长整类型
	CodeTypeFloat     CodeType = 3 // FLOAT - 浮点类型
	CodeTypeDouble    CodeType = 4 // DOUBLE - 双精度类型
	CodeTypeBoolean   CodeType = 5 // BOOLEAN - 布尔值
	CodeTypeString    CodeType = 6 // STRING - 字符串
	CodeTypeText      CodeType = 7 // TEXT - 字符串
	CodeTypeDate      CodeType = 8 // DATETIME - 日期类型
	CodeTypeTimestamp CodeType = 9 // TIMESTAMP - 时间戳
)

// Code 返回枚举的代码值
func (c CodeType) Code() int {
	return int(c)
}

// CodeTypeName 返回枚举的代码类型名称
func (c CodeType) CodeTypeName() string {
	switch c {
	case CodeTypeInteger:
		return "INTEGER"
	case CodeTypeLong:
		return "LONG"
	case CodeTypeFloat:
		return "FLOAT"
	case CodeTypeDouble:
		return "DOUBLE"
	case CodeTypeBoolean:
		return "BOOLEAN"
	case CodeTypeString:
		return "STRING"
	case CodeTypeText:
		return "TEXT"
	case CodeTypeDate:
		return "DATETIME"
	case CodeTypeTimestamp:
		return "TIMESTAMP"
	default:
		return ""
	}
}

// ColumnType 返回列类型
func (c CodeType) ColumnType() string {
	switch c {
	case CodeTypeInteger:
		return "int(64)"
	case CodeTypeLong:
		return "bigint(64)"
	case CodeTypeFloat:
		return "float"
	case CodeTypeDouble:
		return "double"
	case CodeTypeBoolean:
		return "bit(1)"
	case CodeTypeString:
		return "varchar(256)"
	case CodeTypeText:
		return "text"
	case CodeTypeDate:
		return "datetime"
	case CodeTypeTimestamp:
		return "timestamp"
	default:
		return ""
	}
}

// CodeClassType 返回代码类类型
func (c CodeType) CodeClassType() string {
	switch c {
	case CodeTypeInteger:
		return "int"
	case CodeTypeLong:
		return "int64"
	case CodeTypeFloat:
		return "float32"
	case CodeTypeDouble:
		return "float64"
	case CodeTypeBoolean:
		return "bool"
	case CodeTypeString, CodeTypeText:
		return "string"
	case CodeTypeDate, CodeTypeTimestamp:
		return "time.Time"
	default:
		return ""
	}
}

// DbDefaultValue 返回数据库默认值
func (c CodeType) DbDefaultValue() string {
	switch c {
	case CodeTypeInteger, CodeTypeLong:
		return "0"
	case CodeTypeFloat, CodeTypeDouble:
		return "0.0"
	case CodeTypeBoolean:
		return "0"
	case CodeTypeString, CodeTypeText:
		return ""
	case CodeTypeDate:
		return "CURRENT_TIMESTAMP"
	case CodeTypeTimestamp:
		return "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"
	default:
		return ""
	}
}

// Msg 返回枚举的描述信息
func (c CodeType) Msg() string {
	switch c {
	case CodeTypeInteger:
		return "整型数字"
	case CodeTypeLong:
		return "长整类型"
	case CodeTypeFloat:
		return "浮点类型"
	case CodeTypeDouble:
		return "双精度类型"
	case CodeTypeBoolean:
		return "布尔值"
	case CodeTypeString:
		return "字符串"
	case CodeTypeText:
		return "字符串"
	case CodeTypeDate:
		return "日期类型"
	case CodeTypeTimestamp:
		return "时间戳"
	default:
		return ""
	}
}

// GetCodeTypeByCode 根据code获取枚举
func GetCodeTypeByCode(code int) CodeType {
	switch code {
	case 1:
		return CodeTypeInteger
	case 2:
		return CodeTypeLong
	case 3:
		return CodeTypeFloat
	case 4:
		return CodeTypeDouble
	case 5:
		return CodeTypeBoolean
	case 6:
		return CodeTypeString
	case 7:
		return CodeTypeText
	case 8:
		return CodeTypeDate
	case 9:
		return CodeTypeTimestamp
	default:
		return 0
	}
}

// GetCodeTypeName 根据code获取代码类型名称
func GetCodeTypeName(code int) string {
	return GetCodeTypeByCode(code).CodeTypeName()
}

// GetDefaultValue 根据code获取默认值
func (c CodeType) GetDefaultValue() interface{} {
	switch c {
	case CodeTypeInteger:
		return 0
	case CodeTypeLong:
		return int64(0)
	case CodeTypeFloat:
		return float32(0.0)
	case CodeTypeDouble:
		return 0.0
	case CodeTypeBoolean:
		return false
	case CodeTypeString, CodeTypeText:
		return ""
	case CodeTypeDate, CodeTypeTimestamp:
		return time.Time{}
	default:
		return nil
	}
}

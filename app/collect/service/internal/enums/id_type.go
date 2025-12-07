// internal/pkg/enums/id_type.go
package enums

// IdType 证件类型
type IdType int

const (
	IdTypeIdentity                 IdType = 1 // 身份证
	IdTypeHKMTWPermit              IdType = 2 // 港澳台居民居住证
	IdTypeHKMPass                  IdType = 3 // 港澳居民来往内地通行证
	IdTypeTWPass                   IdType = 4 // 台湾居民来往内地通行证
	IdTypePassport                 IdType = 5 // 护照
	IdTypeForeignerResidencePermit IdType = 6 // 外国人永久居住证
)

// Code 返回枚举的代码值
func (i IdType) Code() int {
	return int(i)
}

// Msg 返回枚举的描述信息
func (i IdType) Msg() string {
	switch i {
	case IdTypeIdentity:
		return "身份证"
	case IdTypeHKMTWPermit:
		return "港澳台居民居住证"
	case IdTypeHKMPass:
		return "港澳居民来往内地通行证"
	case IdTypeTWPass:
		return "台湾居民来往内地通行证"
	case IdTypePassport:
		return "护照"
	case IdTypeForeignerResidencePermit:
		return "外国人永久居住证"
	default:
		return ""
	}
}

// GetIdTypeMsg 根据code获取描述信息
func GetIdTypeMsg(code int) string {
	return GetIdTypeByCode(code).Msg()
}

// GetIdTypeByCode 根据code获取枚举
func GetIdTypeByCode(code int) IdType {
	switch code {
	case 1:
		return IdTypeIdentity
	case 2:
		return IdTypeHKMTWPermit
	case 3:
		return IdTypeHKMPass
	case 4:
		return IdTypeTWPass
	case 5:
		return IdTypePassport
	case 6:
		return IdTypeForeignerResidencePermit
	default:
		return 0
	}
}

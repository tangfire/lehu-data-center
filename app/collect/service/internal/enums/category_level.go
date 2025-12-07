// internal/pkg/enums/category_level.go
package enums

// CategoryLevel 维度分类
type CategoryLevel int

const (
	CategoryLevelOne CategoryLevel = 1 // 一级
	CategoryLevelTwo CategoryLevel = 2 // 二级
)

// Code 返回枚举的代码值
func (c CategoryLevel) Code() int {
	return int(c)
}

// Msg 返回枚举的描述信息
func (c CategoryLevel) Msg() string {
	switch c {
	case CategoryLevelOne:
		return "一级"
	case CategoryLevelTwo:
		return "二级"
	default:
		return ""
	}
}

// GetCategoryLevelMsg 根据code获取描述信息
func GetCategoryLevelMsg(code int) string {
	return GetCategoryLevelByCode(code).Msg()
}

// GetCategoryLevelByCode 根据code获取枚举
func GetCategoryLevelByCode(code int) CategoryLevel {
	switch code {
	case 1:
		return CategoryLevelOne
	case 2:
		return CategoryLevelTwo
	default:
		return 0
	}
}

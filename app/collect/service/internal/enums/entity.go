// internal/pkg/enums/entity.go
package enums

// Entity 表名
type Entity int

const (
	EntityVideoReact Entity = 1 // 用户观看视频反应表
	EntityVideo      Entity = 2 // 视频表
	EntityVideoType  Entity = 3 // 视频类型表
)

// Code 返回枚举的代码值
func (t Entity) Code() int {
	return int(t)
}

// Value 返回枚举的字符串值
func (t Entity) Value() string {
	switch t {
	case EntityVideoReact:
		return "d_video_react"
	case EntityVideo:
		return "d_video"
	case EntityVideoType:
		return "d_video_type"
	default:
		return ""
	}
}

// Msg 返回枚举的描述信息
func (t Entity) Msg() string {
	switch t {
	case EntityVideoReact:
		return "用户观看视频反应表"
	case EntityVideo:
		return "视频表"
	case EntityVideoType:
		return "视频类型表"
	default:
		return ""
	}
}

// GetEntityMsg 根据code获取描述信息
func GetEntityMsg(code int) string {
	return GetEntityByCode(code).Msg()
}

// GetEntityByCode 根据code获取枚举
func GetEntityByCode(code int) Entity {
	switch code {
	case 1:
		return EntityVideoReact
	case 2:
		return EntityVideo
	case 3:
		return EntityVideoType
	default:
		return 0
	}
}

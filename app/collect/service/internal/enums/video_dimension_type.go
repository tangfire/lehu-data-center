// internal/pkg/enums/video_dimension_type.go
package enums

// VideoDimensionType 视频维度类型
type VideoDimensionType int

const (
	// VideoDimensionTypeParentVideoType 父级视频分类
	VideoDimensionTypeParentVideoType VideoDimensionType = 1
	// VideoDimensionTypeVideoType 视频分类
	VideoDimensionTypeVideoType VideoDimensionType = 2
	// VideoDimensionTypeVideo 视频本身
	VideoDimensionTypeVideo VideoDimensionType = 3
)

// Code 返回枚举的代码值
func (v VideoDimensionType) Code() int {
	return int(v)
}

// Value 返回枚举的字符串值
func (v VideoDimensionType) Value() string {
	switch v {
	case VideoDimensionTypeVideo:
		return "video"
	case VideoDimensionTypeVideoType:
		return "video_type"
	case VideoDimensionTypeParentVideoType:
		return "parent_video_type"
	default:
		return ""
	}
}

// Msg 返回枚举的描述信息
func (v VideoDimensionType) Msg() string {
	switch v {
	case VideoDimensionTypeVideo:
		return "视频本身"
	case VideoDimensionTypeVideoType:
		return "视频分类"
	case VideoDimensionTypeParentVideoType:
		return "父级视频分类"
	default:
		return ""
	}
}

// GetByCode 根据code获取枚举
func GetByCode(code int) VideoDimensionType {
	switch code {
	case 3:
		return VideoDimensionTypeVideo
	case 2:
		return VideoDimensionTypeVideoType
	case 1:
		return VideoDimensionTypeParentVideoType
	default:
		return 0
	}
}

// GetByValue 根据value获取枚举
func GetByValue(value string) VideoDimensionType {
	switch value {
	case "video":
		return VideoDimensionTypeVideo
	case "video_type":
		return VideoDimensionTypeVideoType
	case "parent_video_type":
		return VideoDimensionTypeParentVideoType
	default:
		return 0
	}
}

// GetMsgByCode 根据code获取描述信息
func GetMsgByCode(code int) string {
	switch code {
	case 3:
		return "视频本身"
	case 2:
		return "视频分类"
	case 1:
		return "父级视频分类"
	default:
		return ""
	}
}

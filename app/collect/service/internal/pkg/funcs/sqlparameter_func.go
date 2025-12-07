// internal/pkg/funcs/sql_parameter_func.go
package funcs

import (
	"lehu-data-center/app/collect/service/internal/enums"
	"time"
)

const (
	StartTime         = "startTime"
	EndTime           = "endTime"
	StatsTime         = "stats_time"
	VideoId           = "video_id"
	VideoIds          = "video_ids"
	VideoTypeId       = "video_type_id"
	ParentVideoTypeId = "parent_video_type_id"
	RuleId            = "rule_id"
)

// SelectDimensionParameter 选择维度参数
func SelectDimensionParameter(startDate, endDate time.Time) map[string]interface{} {
	return map[string]interface{}{
		StartTime: startDate,
		EndTime:   endDate,
	}
}

// GetVideoDbDimensionColumn 对应维度的数据库字段
func GetVideoDbDimensionColumn(videoDimensionType enums.VideoDimensionType) string {
	switch videoDimensionType {
	case enums.VideoDimensionTypeVideo:
		return VideoId
	case enums.VideoDimensionTypeVideoType:
		return VideoTypeId
	case enums.VideoDimensionTypeParentVideoType:
		return ParentVideoTypeId
	default:
		return ""
	}
}

// CreateDimensionCondition 创建维度条件
func CreateDimensionCondition(videoDimensionType enums.VideoDimensionType) string {
	switch videoDimensionType {
	case enums.VideoDimensionTypeVideo:
		return VideoId + " IN (:" + VideoIds + ")"
	case enums.VideoDimensionTypeVideoType:
		return VideoTypeId + " IN (:" + VideoTypeId + ")"
	case enums.VideoDimensionTypeParentVideoType:
		return ParentVideoTypeId + " =:" + ParentVideoTypeId
	default:
		return ""
	}
}

// Offset 计算分页偏移量
func Offset(pageSize, pageNum int) int {
	if pageNum < 1 {
		pageNum = 1
	}
	return (pageNum - 1) * pageSize
}

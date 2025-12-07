package funcs

import (
	"lehu-data-center/app/collect/service/internal/enums"
	"lehu-data-center/app/collect/service/internal/types/transfers"
)

// UpGradeVideoDimension 升级视频维度
func UpGradeVideoDimension(videoDimensionType enums.VideoDimensionType) enums.VideoDimensionType {
	switch videoDimensionType {
	case enums.VideoDimensionTypeVideo:
		return enums.VideoDimensionTypeVideoType
	case enums.VideoDimensionTypeVideoType:
		return enums.VideoDimensionTypeParentVideoType
	default:
		return 0 // 返回零值
	}
}

// DownGradeVideoDimension 降级视频维度
func DownGradeVideoDimension(videoDimensionType enums.VideoDimensionType) enums.VideoDimensionType {
	switch videoDimensionType {
	case enums.VideoDimensionTypeParentVideoType:
		return enums.VideoDimensionTypeVideoType
	case enums.VideoDimensionTypeVideoType:
		return enums.VideoDimensionTypeVideo
	default:
		return 0 // 返回零值
	}
}

// GetParamName 获取参数名
func GetParamName(queryDataTransfers *transfers.QueryDataTransfers) string {
	if queryDataTransfers == nil || queryDataTransfers.QueryDataExecuteTransfers == nil {
		return ""
	}

	queryDataExecuteTransfers := queryDataTransfers.QueryDataExecuteTransfers

	if len(queryDataExecuteTransfers.VideoIdList) > 0 {
		return VideoId
	}

	if len(queryDataExecuteTransfers.VideoTypeIdList) > 0 {
		return VideoTypeId
	}

	if len(queryDataExecuteTransfers.ParentVideoTypeIdList) > 0 {
		return ParentVideoTypeId
	}

	return ""
}

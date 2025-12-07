// transfers/total_param_transfers.go
package transfers

import (
	"lehu-data-center/app/collect/service/internal/entity"
	"lehu-data-center/app/collect/service/internal/pkg/funcs"
)

// TotalParamTransfers 总的参数传输对象
type TotalParamTransfers struct {
	MetricList     []*entity.Metric         `json:"metric_list"`
	ParamTransfers *ParamTransfers          `json:"param_transfers"`
	ResultDataList []map[string]interface{} `json:"result_data_list"`
}

// NewTotalParamTransfers 创建总的参数传输对象
func NewTotalParamTransfers(
	metricList []*entity.Metric,
	paramTransfers *ParamTransfers,
	resultDataList []map[string]interface{},
) *TotalParamTransfers {
	return &TotalParamTransfers{
		MetricList:     metricList,
		ParamTransfers: paramTransfers,
		ResultDataList: resultDataList,
	}
}

// AssemblyParams 组装SQL参数
func (t *TotalParamTransfers) AssemblyParams() map[string]interface{} {
	params := make(map[string]interface{})

	if t.ParamTransfers == nil || t.ParamTransfers.DimensionTransfers == nil {
		return params
	}

	dimensionTransfers := t.ParamTransfers.DimensionTransfers

	// 父级的视频分类id
	if dimensionTransfers.ParentVideoTypeId > 0 {
		params["parent_video_type_id"] = dimensionTransfers.ParentVideoTypeId
	}

	// 视频分类id
	if dimensionTransfers.VideoTypeId > 0 {
		params["video_type_id"] = dimensionTransfers.VideoTypeId
	}

	// 视频id集合
	if len(dimensionTransfers.VideoIdList) > 0 {
		params["video_ids"] = dimensionTransfers.VideoIdList
	}

	// 规则id
	if t.ParamTransfers.RuleId > 0 {
		params["rule_id"] = t.ParamTransfers.RuleId
	}

	// 统计的时间
	countTime := t.AssemblyCountTime()
	if countTime != "" {
		params["stats_time"] = countTime
	}

	return params
}

// AssemblyCountTime 组装统计时间
func (t *TotalParamTransfers) AssemblyCountTime() string {
	if t.ParamTransfers == nil || t.ParamTransfers.RequestTime == nil {
		return ""
	}

	requestTime := t.ParamTransfers.RequestTime

	// 调用datetime包的函数组装时间
	return funcs.AssemblyCountTime(
		requestTime.DateType,
		*requestTime.StartTime,
		*requestTime.EndTime,
	)
}

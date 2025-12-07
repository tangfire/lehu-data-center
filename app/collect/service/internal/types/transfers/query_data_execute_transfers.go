package transfers

import "lehu-data-center/app/collect/service/internal/entity"

// QueryDataExecuteTransfers 查询数据执行传输对象
type QueryDataExecuteTransfers struct {
	ParentVideoTypeIdList []int64         `json:"parent_video_type_id_list"`
	VideoTypeIdList       []int64         `json:"video_type_id_list"`
	VideoIdList           []int64         `json:"video_id_list"`
	MetricList            []entity.Metric `json:"metric_list"`
	RuleDescribe          string          `json:"rule_describe"`
	RuleName              string          `json:"rule_name"`
	DataSave              *DataSave       `json:"data_save"`
	AccumulateColumnName  string          `json:"accumulate_column_name"`
}

// DataSave 数据保存配置
type DataSave struct {
	SaveType int    `json:"save_type"` // 保存类型：1-数据库，2-文件，3-消息队列
	Config   string `json:"config"`    // 配置信息，JSON格式
}

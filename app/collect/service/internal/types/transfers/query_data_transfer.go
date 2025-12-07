// transfers/query_data_transfers.go
package transfers

import (
	"lehu-data-center/app/collect/service/internal/enums"
	"time"
)

// QueryDataTransfers 查询数据传输对象
type QueryDataTransfers struct {
	RuleId             int64                    `json:"rule_id"`
	VideoDimensionType enums.VideoDimensionType `json:"video_dimension_type"`
	VideoTypeId        int                      `json:"video_type_id"`
	DateType           enums.DateType           `json:"date_type"`

	// 开始时间
	StartTime *time.Time `json:"start_time"`

	// 结束时间
	EndTime *time.Time `json:"end_time"`

	// 执行中需要的参数
	QueryDataExecuteTransfers *QueryDataExecuteTransfers `json:"query_data_execute_transfers"`
}

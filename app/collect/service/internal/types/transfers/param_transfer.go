// transfers/param_transfer.go
package transfers

import (
	"lehu-data-center/app/collect/service/internal/enums"
)

// ParamTransfers 参数传输对象
type ParamTransfers struct {
	RuleId             int64                    `json:"rule_id"`
	RuleType           enums.RuleType           `json:"rule_type"`
	VideoDimensionType enums.VideoDimensionType `json:"video_dimension_type"`
	DimensionTransfers *DimensionTransfer       `json:"dimension_transfers"`
	RequestTime        *RequestTime             `json:"request_time"`
	CollectType        enums.CollectType        `json:"collect_type"`
	QueryDataTransfers *QueryDataTransfers      `json:"query_data_transfers"`

	// 消息的父级链路id
	MessageParentTraceId int64 `json:"message_parent_trace_id"`

	// 消息的链路id
	MessageTraceId int64 `json:"message_trace_id"`
}

package transfers

import (
	"lehu-data-center/app/collect/service/internal/enums"
	"time"
)

// RequestTime 请求时间
type RequestTime struct {
	GatherDate time.Time      `json:"gather_date"`
	DateType   enums.DateType `json:"date_type"`
	StartTime  *time.Time     `json:"start_time"`
	EndTime    *time.Time     `json:"end_time"`
}

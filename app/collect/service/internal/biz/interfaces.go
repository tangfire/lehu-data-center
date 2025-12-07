package biz

import (
	"context"
	"lehu-data-center/app/collect/service/internal/enums"
	"lehu-data-center/app/collect/service/internal/types/transfers"
	"time"
)

// DataHandler 数据处理器接口
type DataHandler interface {
	DataHandle(ctx context.Context, totalParam *transfers.TotalParamTransfers) (*transfers.RuleHandleOutput, error)
	GetRuleType() int8
}

// GatherHandler 采集处理器接口
type GatherHandler interface {
	DoGather(ctx context.Context, totalParam *transfers.TotalParamTransfers) error
}

// SaveHandler 保存处理器接口
type SaveHandler interface {
	Save(ctx context.Context, totalParam *transfers.TotalParamTransfers) error
}

// UpgradeHandler 升级处理器接口
type UpgradeHandler interface {
	UpgradeDimension(
		ctx context.Context,
		id int64,
		messageID string,
		totalParam *transfers.TotalParamTransfers,
	) (*UpgradeDimensionDomain, error)
}

// MessageRecord 消息记录
type MessageRecord struct {
	ID        int64
	MessageID int64
	Status    enums.MessageSendStatus
	CreatedAt time.Time
}

// MessageService 消息服务接口
type MessageService interface {
	InsertMessageRecord(ctx context.Context, totalParam *transfers.TotalParamTransfers) (*MessageRecord, error)
}

// UpgradeDimensionDomain 维度升级领域对象
type UpgradeDimensionDomain struct {
	ID      int64
	Message string
	Error   error
}

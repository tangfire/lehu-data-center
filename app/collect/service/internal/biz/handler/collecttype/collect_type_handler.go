package collecttype

import (
	"context"
	"lehu-data-center/app/collect/service/internal/biz"
	"lehu-data-center/app/collect/service/internal/biz/transfers"
	"lehu-data-center/app/collect/service/internal/enums"
)

// CollectTypeHandler 采集类型处理器接口
type CollectTypeHandler interface {
	// CollectType 处理器对应的采集类型
	CollectType() enums.CollectType

	// DoCollect 执行采集
	DoCollect(ctx context.Context, dimensionGather *biz.DimensionGatherModel, requestTime *biz.RequestTime) ([]*transfers.DimensionTransfer, error)
}

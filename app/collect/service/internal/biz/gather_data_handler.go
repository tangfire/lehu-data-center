package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/enums"
	"lehu-data-center/app/collect/service/internal/types/transfers"
)

// GatherDataHandler 数据采集处理器
type GatherDataHandler struct {
	chainContext *ChainContext
	log          *log.Helper
}

func NewGatherDataHandler(chainContext *ChainContext, logger log.Logger) *GatherDataHandler {
	return &GatherDataHandler{
		chainContext: chainContext,
		log:          log.NewHelper(logger),
	}
}

func (h *GatherDataHandler) DataHandle(
	ctx context.Context,
	totalParam *transfers.TotalParamTransfers,
) (*transfers.RuleHandleOutput, error) {

	// 使用责任链进行数据处理
	if err := h.chainContext.Execute(ctx, totalParam); err != nil {
		h.log.WithContext(ctx).Errorf("Execute chain failed: %v", err)
		return transfers.NewErrorRuleHandleOutput("处理失败"), err
	}

	return transfers.NewSuccessRuleHandleOutput("处理成功", nil), nil
}

func (h *GatherDataHandler) GetRuleType() int8 {
	return int8(enums.RuleTypeGather)
}

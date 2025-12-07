package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/types/transfers"
)

// SaveChainHandler 数据保存责任链处理器
type SaveChainHandler struct {
	BaseChainHandler
	saveHandler SaveHandler
	log         *log.Helper
}

func NewSaveChainHandler(saveHandler SaveHandler, logger log.Logger) *SaveChainHandler {
	return &SaveChainHandler{
		saveHandler: saveHandler,
		log:         log.NewHelper(logger),
	}
}

func (h *SaveChainHandler) Handler(ctx context.Context, totalParam *transfers.TotalParamTransfers) error {
	if err := h.saveHandler.Save(ctx, totalParam); err != nil {
		return fmt.Errorf("save data failed: %w", err)
	}

	h.log.WithContext(ctx).Infof("Data saved successfully, ruleId: %d",
		totalParam.ParamTransfers.RuleId)
	return nil
}

func (h *SaveChainHandler) Order() int {
	return 3
}

package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/enums"
	"lehu-data-center/app/collect/service/internal/types/transfers"
	"sync"
)

// GatherChainHandler 数据采集责任链处理器
type GatherChainHandler struct {
	BaseChainHandler
	gatherContext *GatherHandlerContext
	log           *log.Helper
}

func NewGatherChainHandler(gatherContext *GatherHandlerContext, logger log.Logger) *GatherChainHandler {
	return &GatherChainHandler{
		gatherContext: gatherContext,
		log:           log.NewHelper(logger),
	}
}

func (h *GatherChainHandler) Handler(ctx context.Context, totalParam *transfers.TotalParamTransfers) error {
	param := totalParam.ParamTransfers
	if param == nil {
		return ErrInvalidParam
	}

	// 获取采集处理器
	gatherHandler, err := h.gatherContext.GetHandler(
		param.VideoDimensionType,
		param.RequestTime.DateType,
	)
	if err != nil {
		h.log.WithContext(ctx).Warnf("Get gather handler failed: %v", err)
		return nil // 返回nil继续执行后续链
	}

	// 执行采集
	if err := gatherHandler.DoGather(ctx, totalParam); err != nil {
		return fmt.Errorf("gather failed: %w", err)
	}

	return nil
}

func (h *GatherChainHandler) Order() int {
	return 1
}

// GatherHandlerContext 采集处理器上下文
type GatherHandlerContext struct {
	handlers map[string]GatherHandler
	mu       sync.RWMutex
}

func NewGatherHandlerContext() *GatherHandlerContext {
	return &GatherHandlerContext{
		handlers: make(map[string]GatherHandler),
	}
}

func (c *GatherHandlerContext) RegisterHandler(key string, handler GatherHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.handlers[key] = handler
}

func (c *GatherHandlerContext) GetHandler(videoDimensionType enums.VideoDimensionType, dateType enums.DateType) (GatherHandler, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	key := fmt.Sprintf("%d_%d", videoDimensionType, dateType)
	handler, exists := c.handlers[key]
	if !exists {
		return nil, ErrHandlerNotFound
	}

	return handler, nil
}

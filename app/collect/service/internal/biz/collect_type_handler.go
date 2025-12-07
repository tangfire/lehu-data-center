package biz

import (
	"context"
	"lehu-data-center/app/collect/service/internal/entity"
	"lehu-data-center/app/collect/service/internal/enums"
	"lehu-data-center/app/collect/service/internal/types/transfers"
	"sync"
)

type CollectTypeHandlerContext struct {
	handlers map[enums.CollectType]CollectTypeHandler
	mu       sync.RWMutex
}

func NewCollectTypeHandlerContext(handlers ...CollectTypeHandler) *CollectTypeHandlerContext {
	ctx := &CollectTypeHandlerContext{
		handlers: make(map[enums.CollectType]CollectTypeHandler),
	}

	for _, handler := range handlers {
		ctx.RegisterHandler(handler)
	}

	return ctx
}

func (c *CollectTypeHandlerContext) RegisterHandler(handler CollectTypeHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.handlers[handler.CollectType()] = handler
}

func (c *CollectTypeHandlerContext) GetHandler(collectType enums.CollectType) (CollectTypeHandler, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	handler, exists := c.handlers[collectType]
	if !exists {
		return nil, ErrCollectTypeHandlerNotFound
	}

	return handler, nil
}

// CollectTypeHandler 采集类型处理器接口
type CollectTypeHandler interface {
	// CollectType 处理器对应的采集类型
	CollectType() enums.CollectType

	// DoCollect 执行采集
	DoCollect(ctx context.Context, dimensionGather *entity.DimensionGather, requestTime *transfers.RequestTime) ([]*transfers.DimensionTransfer, error)
}

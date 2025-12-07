package biz

import (
	"context"
	"lehu-data-center/app/collect/service/internal/types/transfers"
	"sync"
)

// AbstractChainHandler 责任链抽象处理器
type AbstractChainHandler interface {
	SetNext(handler AbstractChainHandler) AbstractChainHandler
	ExecuteChain(ctx context.Context, totalParam *transfers.TotalParamTransfers) error
	Handler(ctx context.Context, totalParam *transfers.TotalParamTransfers) error
	Order() int
}

// BaseChainHandler 基础责任链处理器
type BaseChainHandler struct {
	nextHandler AbstractChainHandler
	mu          sync.RWMutex
}

func (b *BaseChainHandler) SetNext(handler AbstractChainHandler) AbstractChainHandler {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.nextHandler = handler
	return handler
}

func (b *BaseChainHandler) ExecuteChain(ctx context.Context, totalParam *transfers.TotalParamTransfers) error {
	if err := b.Handler(ctx, totalParam); err != nil {
		return err
	}

	if b.nextHandler != nil {
		return b.nextHandler.ExecuteChain(ctx, totalParam)
	}

	return nil
}

func (b *BaseChainHandler) Handler(_ context.Context, _ *transfers.TotalParamTransfers) error {
	// 基础类不实现具体逻辑，由子类实现
	return nil
}

func (b *BaseChainHandler) Order() int {
	return 0
}

// ChainContext 责任链上下文
type ChainContext struct {
	handlers []AbstractChainHandler
	first    AbstractChainHandler
	mu       sync.RWMutex
}

func NewChainContext(handlers ...AbstractChainHandler) *ChainContext {
	ctx := &ChainContext{
		handlers: make([]AbstractChainHandler, 0),
	}

	for _, handler := range handlers {
		ctx.RegisterHandler(handler)
	}

	ctx.buildChain()
	return ctx
}

func (c *ChainContext) RegisterHandler(handler AbstractChainHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.handlers = append(c.handlers, handler)
	c.buildChain()
}

func (c *ChainContext) buildChain() {
	if len(c.handlers) == 0 {
		return
	}

	// 按照Order排序
	c.sortHandlers()

	// 构建链
	first := c.handlers[0]
	current := first

	for i := 1; i < len(c.handlers); i++ {
		current = current.SetNext(c.handlers[i])
	}

	c.first = first
}

func (c *ChainContext) sortHandlers() {
	for i := 0; i < len(c.handlers)-1; i++ {
		for j := i + 1; j < len(c.handlers); j++ {
			if c.handlers[i].Order() > c.handlers[j].Order() {
				c.handlers[i], c.handlers[j] = c.handlers[j], c.handlers[i]
			}
		}
	}
}

func (c *ChainContext) GetChain() AbstractChainHandler {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.first
}

func (c *ChainContext) Execute(ctx context.Context, totalParam *transfers.TotalParamTransfers) error {
	chain := c.GetChain()
	if chain == nil {
		return ErrChainNotInitialized
	}

	return chain.ExecuteChain(ctx, totalParam)
}

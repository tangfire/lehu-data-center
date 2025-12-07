package biz

import (
	"sync"
)

type DataHandlerContext struct {
	handlers map[int8]DataHandler
	mu       sync.RWMutex
}

func NewDataHandlerContext(handlers ...DataHandler) *DataHandlerContext {
	ctx := &DataHandlerContext{
		handlers: make(map[int8]DataHandler),
	}

	for _, handler := range handlers {
		ctx.RegisterHandler(handler)
	}

	return ctx
}

func (c *DataHandlerContext) RegisterHandler(handler DataHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.handlers[handler.GetRuleType()] = handler
}

func (c *DataHandlerContext) GetHandler(ruleType int8) (DataHandler, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	handler, exists := c.handlers[ruleType]
	if !exists {
		return nil, ErrDataHandlerNotFound
	}

	return handler, nil
}

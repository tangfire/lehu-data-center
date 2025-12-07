package biz

import "errors"

var (
	// ErrCollectTypeHandlerNotFound 采集类型处理器未找到
	ErrCollectTypeHandlerNotFound = errors.New("collect type handler not found")

	// ErrRuleNotFound 规则未找到
	ErrRuleNotFound = errors.New("rule not found")

	// ErrDimensionGatherNotFound 维度收集配置未找到
	ErrDimensionGatherNotFound = errors.New("dimension gather not found")

	ErrDataHandlerNotFound = errors.New("data handler not found")

	ErrChainNotInitialized = errors.New("chain not initialized")

	ErrInvalidParam = errors.New("invalid param")

	ErrHandlerNotFound = errors.New("handler not found")
)

package biz

import (
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewDataJobUsecase,
	NewMessageConsumerRecordUsecase,
	NewMessageProducerRecordUsecase,
	NewVideoBusinessMessageConsumerRecordUsecase,
	NewVideoBusinessMessageProducerRecordUsecase,
	NewCollectUsecase,
	NewRuleUsecase,
	NewDimensionGatherUsecase,
	NewSqlCollectHandler,
	NewCollectTypeHandlerContext,
	NewMetricUsecase,
	NewDataHandlerContext)

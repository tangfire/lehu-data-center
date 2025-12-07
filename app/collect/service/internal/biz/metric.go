package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/entity"
)

type MetricRepo interface {
	GetMetricListByRuleId(ctx context.Context, ruleId int64) ([]*entity.Metric, error)
}

type MetricUsecase struct {
	metricRepo MetricRepo
	log        *log.Helper
}

func NewMetricUsecase(metricRepo MetricRepo, logger log.Logger) *MetricUsecase {
	return &MetricUsecase{
		metricRepo: metricRepo,
		log:        log.NewHelper(logger),
	}
}

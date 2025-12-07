package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/entity"
	"lehu-data-center/app/collect/service/internal/enums"
	"lehu-data-center/app/collect/service/internal/types/transfers"
)

type DimensionGatherRepo interface {
	GetDimensionGatherByRuleIdAndEntity(context.Context, int64, enums.Entity) (*entity.DimensionGather, error)
}

type DimensionGatherUsecase struct {
	repo       DimensionGatherRepo
	ruleRepo   RuleRepo
	handlerCtx *CollectTypeHandlerContext
	log        *log.Helper
}

func NewDimensionGatherUsecase(
	repo DimensionGatherRepo,
	ruleRepo RuleRepo,
	handlerCtx *CollectTypeHandlerContext,
	logger log.Logger,
) *DimensionGatherUsecase {
	return &DimensionGatherUsecase{
		repo:       repo,
		ruleRepo:   ruleRepo,
		handlerCtx: handlerCtx,
		log:        log.NewHelper(logger),
	}
}

func (uc *DimensionGatherUsecase) HandleDimensionData(
	ctx context.Context,
	ruleId int64,
	requestTime *transfers.RequestTime,
) ([]*transfers.DimensionTransfer, error) {
	// 1. 查询规则
	rule, err := uc.ruleRepo.GetRuleById(ctx, ruleId)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("HandleDimensionData|GetRuleById fail, ruleId:%v, err:%v", ruleId, err)
		return nil, fmt.Errorf("get rule failed: %w", err)
	}

	if rule == nil {
		uc.log.WithContext(ctx).Warnf("HandleDimensionData|rule not found, ruleId:%v", ruleId)
		return []*transfers.DimensionTransfer{}, nil
	}

	// 2. 获取维度收集配置
	dimensionGather, err := uc.repo.GetDimensionGatherByRuleIdAndEntity(
		ctx,
		ruleId,
		enums.EntityVideoReact,
	)
	if err != nil {
		uc.log.WithContext(ctx).Errorf(
			"HandleDimensionData|GetDimensionGatherByRuleIdAndEntity fail, ruleId:%v, err:%v",
			ruleId,
			err,
		)
		return nil, fmt.Errorf("get dimension gather failed: %w", err)
	}

	if dimensionGather == nil {
		uc.log.WithContext(ctx).Warnf(
			"HandleDimensionData|dimension gather not found, ruleId:%v",
			ruleId,
		)
		return []*transfers.DimensionTransfer{}, nil
	}

	// 3. 获取对应的采集处理器
	collectType := enums.CollectType(dimensionGather.CollectType)
	handler, err := uc.handlerCtx.GetHandler(collectType)
	if err != nil {
		uc.log.WithContext(ctx).Errorf(
			"HandleDimensionData|GetHandler fail, collectType:%v, err:%v",
			collectType,
			err,
		)
		return nil, fmt.Errorf("get collect type handler failed: %w", err)
	}

	// 4. 执行采集
	result, err := handler.DoCollect(ctx, dimensionGather, requestTime)
	if err != nil {
		uc.log.WithContext(ctx).Errorf(
			"HandleDimensionData|DoCollect fail, ruleId:%v, collectType:%v, err:%v",
			ruleId,
			collectType,
			err,
		)
		return nil, fmt.Errorf("do collect failed: %w", err)
	}

	uc.log.WithContext(ctx).Infof(
		"HandleDimensionData success, ruleId:%v, collectType:%v, resultCount:%d",
		ruleId,
		collectType,
		len(result),
	)

	return result, nil
}

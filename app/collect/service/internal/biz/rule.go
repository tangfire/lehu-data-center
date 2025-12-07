package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/entity"
	"lehu-data-center/app/collect/service/internal/types/transfers"
)

type RuleRepo interface {
	GetRuleById(context.Context, int64) (*entity.Rule, error)
}

type RuleUsecase struct {
	repo           RuleRepo
	metricRepo     MetricRepo
	dataHandlerCtx *DataHandlerContext
	log            *log.Helper
}

func NewRuleUsecase(repo RuleRepo, metricRepo MetricRepo, dataHandlerCtx *DataHandlerContext, logger log.Logger) *RuleUsecase {
	return &RuleUsecase{repo: repo,
		metricRepo:     metricRepo,
		dataHandlerCtx: dataHandlerCtx,
		log:            log.NewHelper(logger)}
}

func (h *RuleUsecase) Handle(ctx context.Context, paramTransfers *transfers.ParamTransfers) (*transfers.RuleHandleOutput, error) {
	// 1. 根据规则ID查询指标信息
	metricList, err := h.metricRepo.GetMetricListByRuleId(ctx, paramTransfers.RuleId)
	if err != nil {
		h.log.WithContext(ctx).Errorf("Handle|GetMetricListByRuleId fail,ruleId:%d,err:%v",
			paramTransfers.RuleId, err)
		return transfers.NewErrorRuleHandleOutput("查询指标列表失败"), err
	}

	if len(metricList) == 0 {
		h.log.WithContext(ctx).Warnf("Handle|No metrics found for ruleId:%d", paramTransfers.RuleId)
		return transfers.NewSuccessRuleHandleOutput("未找到指标数据", nil), nil
	}

	// 2. 构建总参数
	totalParam := transfers.NewTotalParamTransfers(
		metricList,
		paramTransfers,
		nil, // 采集任务中为空
	)

	// 3. 根据规则类型获取数据处理器
	dataHandler, err := h.dataHandlerCtx.GetHandler(int8(paramTransfers.RuleType))
	if err != nil {
		h.log.WithContext(ctx).Errorf("Handle|GetDataHandler fail,ruleType:%d,err:%v",
			paramTransfers.RuleType, err)
		return transfers.NewErrorRuleHandleOutput("获取数据处理器失败"), err
	}

	// 4. 执行数据处理
	output, err := dataHandler.DataHandle(ctx, totalParam)
	if err != nil {
		h.log.WithContext(ctx).Errorf("Handle|DataHandle fail,ruleId:%d,err:%v",
			paramTransfers.RuleId, err)
		return transfers.NewErrorRuleHandleOutput("数据处理失败"), err
	}

	h.log.WithContext(ctx).Infof("Handle completed successfully, ruleId:%d, metrics:%d",
		paramTransfers.RuleId, len(metricList))

	return output, nil
}

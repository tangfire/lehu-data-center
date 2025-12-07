package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/entity"
	"lehu-data-center/app/collect/service/internal/enums"
	"lehu-data-center/app/collect/service/internal/types/transfers"
	"strings"
)

// ComputeChainHandler 数据计算责任链处理器
type ComputeChainHandler struct {
	BaseChainHandler
	expressionHandler *ExpressionHandler
	log               *log.Helper
}

func NewComputeChainHandler(expressionHandler *ExpressionHandler, logger log.Logger) *ComputeChainHandler {
	return &ComputeChainHandler{
		expressionHandler: expressionHandler,
		log:               log.NewHelper(logger),
	}
}

func (h *ComputeChainHandler) Handler(ctx context.Context, totalParam *transfers.TotalParamTransfers) error {
	resultDataList := totalParam.ResultDataList
	if len(resultDataList) == 0 {
		h.log.WithContext(ctx).Debug("No data to compute")
		return nil
	}

	// 筛选需要计算的指标
	var computeMetrics []*entity.Metric
	for _, metric := range totalParam.MetricList {
		if metric.MetricCollectType == enums.MetricCollectTypeCompute {
			computeMetrics = append(computeMetrics, metric)
		}
	}

	if len(computeMetrics) == 0 {
		h.log.WithContext(ctx).Debug("No metrics to compute")
		return nil
	}

	// 创建指标映射
	metricMap := make(map[string]*entity.Metric)
	for _, metric := range totalParam.MetricList {
		metricMap[metric.MetricName] = metric
	}

	// 计算结果数据
	for _, resultData := range resultDataList {
		for _, computeMetric := range computeMetrics {
			if err := h.compute(ctx, metricMap, resultData, computeMetric); err != nil {
				h.log.WithContext(ctx).Errorf("Compute metric failed: %v", err)
				// 继续处理其他指标
				continue
			}
		}
	}

	return nil
}

func (h *ComputeChainHandler) compute(ctx context.Context, metricMap map[string]*entity.Metric,
	resultData map[string]interface{}, computeMetric *entity.Metric) error {

	metric, exists := metricMap[computeMetric.MetricName]
	if !exists {
		return fmt.Errorf("metric not found: %s", computeMetric.MetricName)
	}

	// 获取数据类型
	codeType, err := enums.GetCodeType(metric.CodeType)
	if err != nil {
		return fmt.Errorf("invalid code type: %w", err)
	}

	// 如果是平均值计算，需要计算数量
	if computeMetric.FunctionType == enums.FunctionTypeAvg {
		if computeMetric.Arguments != "" {
			args := strings.Split(computeMetric.Arguments, ",")
			resultData["count"] = len(args)
		}
	}

	// 获取默认值
	defaultValue := h.expressionHandler.GetDefaultValue(codeType)

	// 表达式计算
	computeValue := defaultValue
	if len(resultData) > 0 {
		computeValue, err = h.expressionHandler.HandleExpression(
			codeType.GetCodeClassType(),
			computeMetric.Expression,
			resultData,
			defaultValue,
		)
		if err != nil {
			h.log.WithContext(ctx).Errorf("Handle expression failed: %v", err)
			computeValue = defaultValue
		}
	}

	resultData[computeMetric.MetricName] = computeValue
	return nil
}

func (h *ComputeChainHandler) Order() int {
	return 2
}

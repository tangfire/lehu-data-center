package funcs

import (
	"errors"
	"fmt"
	"lehu-data-center/app/collect/service/internal/entity"
	"lehu-data-center/app/collect/service/internal/enums"
	"lehu-data-center/app/collect/service/internal/types/transfers"
	"strings"
)

// CreateGatherSql 创建采集SQL
func CreateGatherSql(rule *entity.Rule, dataSave *entity.DataSave, totalParamTransfers *transfers.TotalParamTransfers) (string, error) {
	if rule == nil || dataSave == nil || totalParamTransfers == nil {
		return "", errors.New("参数不能为空")
	}

	videoDimensionType := totalParamTransfers.ParamTransfers.VideoDimensionType

	// 过滤指标，留下基础指标并且可以高维度采集的
	var baseMetrics []*entity.Metric
	for _, metric := range totalParamTransfers.MetricList {
		// 修正：使用 MetricCollectTypeBase 而不是 MetricCollectTypeBasic
		if metric.MetricCollectType == enums.MetricCollectTypeBase {
			baseMetrics = append(baseMetrics, metric)
		}
	}

	if len(baseMetrics) == 0 {
		return "", nil
	}

	// 获取去重后的指标名称
	metricNames := make(map[string]bool)
	for _, metric := range baseMetrics {
		metricNames[metric.MetricName] = true
	}

	// 构建查询列：sum(watch_count) as watch_count
	var queryColumns []string
	for metricName := range metricNames {
		queryColumns = append(queryColumns, fmt.Sprintf("sum(%s) as %s", metricName, metricName))
	}
	queryDbColumn := strings.Join(queryColumns, ", ")

	// 统计的表名 - 注意：这里需要修复调用方式
	tableName := CreateTableName(dataSave.CreateTablePrefix, rule.RuleName,
		enums.VideoDimensionTypeVideo, enums.DateTypeDay)

	// 获取维度字段
	dbDimensionColumn := GetVideoDbDimensionColumn(videoDimensionType)
	if dbDimensionColumn == "" {
		return "", errors.New("不支持的视频维度类型")
	}

	// 创建维度条件
	dimensionCondition := CreateDimensionCondition(videoDimensionType)
	if dimensionCondition == "" {
		return "", errors.New("无法创建维度条件")
	}

	// 构建完整SQL
	sql := fmt.Sprintf("SELECT %s, %s FROM %s WHERE stats_time >= :startTime AND stats_time <= :endTime AND %s GROUP BY %s",
		dbDimensionColumn, queryDbColumn, tableName, dimensionCondition, dbDimensionColumn)

	return sql, nil
}

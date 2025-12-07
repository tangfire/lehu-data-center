package funcs

import (
	"errors"
	"fmt"
	"lehu-data-center/app/collect/service/internal/entity"
	"lehu-data-center/app/collect/service/internal/enums"
	"lehu-data-center/app/collect/service/internal/types/transfers"
	"math"
	"reflect"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

// GetDateTypeDefaultValue 获取数据类型的默认值
func GetDateTypeDefaultValue(codeType enums.CodeType) (interface{}, error) {
	if codeType == enums.CodeTypeDate || codeType == enums.CodeTypeTimestamp {
		return time.Now(), nil
	}

	codeClassType := codeType.CodeClassType()
	if codeClassType == "" {
		return nil, errors.New("不支持的数据类型")
	}

	dbDefaultValue := codeType.DbDefaultValue()

	// 根据类型解析默认值
	switch codeClassType {
	case "int", "int32":
		val, err := strconv.ParseInt(dbDefaultValue, 10, 32)
		if err != nil {
			return int32(0), nil
		}
		return int32(val), nil
	case "int64":
		val, err := strconv.ParseInt(dbDefaultValue, 10, 64)
		if err != nil {
			return int64(0), nil
		}
		return val, nil
	case "float32":
		val, err := strconv.ParseFloat(dbDefaultValue, 32)
		if err != nil {
			return float32(0), nil
		}
		return float32(val), nil
	case "float64":
		val, err := strconv.ParseFloat(dbDefaultValue, 64)
		if err != nil {
			return 0.0, nil
		}
		return val, nil
	case "string":
		return dbDefaultValue, nil
	case "decimal.Decimal":
		return decimal.NewFromFloat(0), nil
	case "time.Time":
		return time.Time{}, nil
	default:
		return nil, fmt.Errorf("不支持的数据类型: %s", codeClassType)
	}
}

// DisposalInitialDbData 处理初始数据库数据
func DisposalInitialDbData(totalParamTransfers *transfers.TotalParamTransfers) ([]map[string]interface{}, error) {
	if totalParamTransfers == nil {
		return nil, errors.New("参数不能为空")
	}

	// 获取指标的默认值
	metricDefaultValueMap := make(map[string]interface{})
	for _, metric := range totalParamTransfers.MetricList {
		codeType := enums.GetCodeTypeByCode(metric.CodeType)
		defaultValue, err := GetDateTypeDefaultValue(codeType)
		if err != nil {
			return nil, fmt.Errorf("获取指标默认值失败: %w", err)
		}
		metricDefaultValueMap[metric.MetricName] = defaultValue
	}

	videoIdList := totalParamTransfers.ParamTransfers.DimensionTransfers.VideoIdList
	videoDimensionType := totalParamTransfers.ParamTransfers.VideoDimensionType
	ruleId := totalParamTransfers.ParamTransfers.RuleId
	countTime := totalParamTransfers.AssemblyCountTime()
	videoTypeId := totalParamTransfers.ParamTransfers.DimensionTransfers.VideoTypeId
	parentVideoTypeId := totalParamTransfers.ParamTransfers.DimensionTransfers.ParentVideoTypeId

	var initialResultList []map[string]interface{}

	// 如果维度是视频本身
	if videoDimensionType.Value() == enums.VideoDimensionTypeVideo.Value() {
		for _, videoId := range videoIdList {
			initialResultMap, err := createInitialDbData(ruleId, countTime, videoId, videoTypeId, parentVideoTypeId, metricDefaultValueMap)
			if err != nil {
				return nil, err
			}
			initialResultList = append(initialResultList, initialResultMap)
		}
	} else if videoDimensionType.Value() == enums.VideoDimensionTypeVideoType.Value() {
		// 如果维度是视频分类
		initialResultMap, err := createInitialDbData(ruleId, countTime, 0, videoTypeId, parentVideoTypeId, metricDefaultValueMap)
		if err != nil {
			return nil, err
		}
		initialResultList = append(initialResultList, initialResultMap)
	} else if videoDimensionType.Value() == enums.VideoDimensionTypeParentVideoType.Value() {
		// 如果维度是父级的视频分类
		initialResultMap, err := createInitialDbData(ruleId, countTime, 0, 0, parentVideoTypeId, metricDefaultValueMap)
		if err != nil {
			return nil, err
		}
		initialResultList = append(initialResultList, initialResultMap)
	}

	return initialResultList, nil
}

func createInitialDbData(ruleId int64, countTime string, videoId, videoTypeId, parentVideoTypeId int64,
	metricDefaultValueMap map[string]interface{}) (map[string]interface{}, error) {

	result := make(map[string]interface{})

	if ruleId != 0 {
		result["rule_id"] = ruleId
	}

	if countTime != "" {
		result["stats_time"] = countTime
	}

	if videoId != 0 {
		result["video_id"] = videoId
	}

	if videoTypeId != 0 {
		result["video_type_id"] = videoTypeId
	}

	if parentVideoTypeId != 0 {
		result["parent_video_type_id"] = parentVideoTypeId
	}

	// 添加指标默认值
	for key, value := range metricDefaultValueMap {
		result[key] = value
	}

	return result, nil
}

// DisposalRelDbData 处理关联数据库数据
func DisposalRelDbData(videoDimensionType enums.VideoDimensionType,
	initialDbDataList []map[string]interface{},
	dbDataList []map[string]interface{}) error {

	if len(dbDataList) == 0 {
		return nil
	}

	dbDimensionColumn := GetVideoDbDimensionColumn(videoDimensionType)
	if dbDimensionColumn == "" {
		return errors.New("无法获取视频维度数据库字段")
	}

	for _, initialDbData := range initialDbDataList {
		for _, dbData := range dbDataList {
			initValue, ok1 := initialDbData[dbDimensionColumn]
			dbValue, ok2 := dbData[dbDimensionColumn]
			if ok1 && ok2 && reflect.DeepEqual(initValue, dbValue) {
				// 合并数据
				for key, value := range dbData {
					if key != dbDimensionColumn { // 避免覆盖维度字段
						initialDbData[key] = value
					}
				}
			}
		}
	}

	return nil
}

// Filter 过滤数据
func Filter(results []map[string]interface{}, baseMetricList []*entity.Metric) []map[string]interface{} {
	if len(results) == 0 || len(baseMetricList) == 0 {
		return []map[string]interface{}{}
	}

	// 获取指标名字
	metricNameMap := make(map[string]bool)
	for _, metric := range baseMetricList {
		metricNameMap[metric.MetricName] = true
	}

	var filteredResults []map[string]interface{}

	for _, result := range results {
		shouldInclude := false

		for metricName := range metricNameMap {
			value, exists := result[metricName]
			if !exists || value == nil {
				continue
			}

			// 尝试转换为数值并检查是否为0
			switch v := value.(type) {
			case int, int8, int16, int32, int64:
				if reflect.ValueOf(v).Int() != 0 {
					shouldInclude = true
				}
			case uint, uint8, uint16, uint32, uint64:
				if reflect.ValueOf(v).Uint() != 0 {
					shouldInclude = true
				}
			case float32:
				if math.Abs(float64(v)) > 1e-10 {
					shouldInclude = true
				}
			case float64:
				if math.Abs(v) > 1e-10 {
					shouldInclude = true
				}
			case decimal.Decimal:
				if !v.IsZero() {
					shouldInclude = true
				}
			case string:
				// 尝试解析字符串为数值
				if f, err := strconv.ParseFloat(v, 64); err == nil && math.Abs(f) > 1e-10 {
					shouldInclude = true
				}
			default:
				// 使用反射获取数值
				val := reflect.ValueOf(v)
				if val.Kind() == reflect.Ptr {
					val = val.Elem()
				}
				if val.CanFloat() {
					if val.Float() != 0 {
						shouldInclude = true
					}
				}
			}

			if shouldInclude {
				break
			}
		}

		if shouldInclude {
			filteredResults = append(filteredResults, result)
		}
	}

	return filteredResults
}

// GetRelNumberType 获取相关的数值类型
func GetRelNumberType(numberValue interface{}) (interface{}, error) {
	switch v := numberValue.(type) {
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return v, nil
	case int64:
		return v, nil
	case uint:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case float32:
		return v, nil
	case float64:
		return v, nil
	case decimal.Decimal:
		return v, nil
	default:
		return nil, fmt.Errorf("不支持的数值类型: %T", numberValue)
	}
}

// GetRelNumberClass 获取相关的数值类
func GetRelNumberClass(numberValue interface{}) (reflect.Type, error) {
	switch numberValue.(type) {
	case int:
		return reflect.TypeOf(int(0)), nil
	case int32:
		return reflect.TypeOf(int32(0)), nil
	case int64:
		return reflect.TypeOf(int64(0)), nil
	case float32:
		return reflect.TypeOf(float32(0)), nil
	case float64:
		return reflect.TypeOf(float64(0)), nil
	case decimal.Decimal:
		return reflect.TypeOf(decimal.Decimal{}), nil
	default:
		return nil, fmt.Errorf("不支持的数值类型: %T", numberValue)
	}
}

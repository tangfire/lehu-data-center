// internal/pkg/enums/base_code.go
package enums

// BaseCode 接口返回code码
type BaseCode int

const (
	BaseCodeSuccess                        BaseCode = 0    // OK
	BaseCodeSystemError                    BaseCode = -1   // 系统异常，请稍后重试
	BaseCodeParameterError                 BaseCode = -2   // 参数验证异常
	BaseCodeDataSourceCloseError           BaseCode = -3   // 数据源关闭失败
	BaseCodeUserNotLogin                   BaseCode = 1001 // 用户未登录
	BaseCodeUIDWorkIDError                 BaseCode = 5000 // uid_work_id设置失败
	BaseCodeCollectTypeNotExist            BaseCode = 5001 // 采集方式不存在
	BaseCodeMetricCollectType              BaseCode = 5002 // 指标收集类型错误
	BaseCodeCollectSourceNameEmpty         BaseCode = 5003 // 收集源头名字为空
	BaseCodeCollectTypeEmpty               BaseCode = 5004 // 收集方式为空
	BaseCodeCollectDetailEmpty             BaseCode = 5005 // 具体的收集实现为空
	BaseCodeMetricTimeFormatEmpty          BaseCode = 5006 // 统计的时间格式为空
	BaseCodeArgumentsEmpty                 BaseCode = 5007 // 指标统计的参数为空
	BaseCodeFunctionTypeEmpty              BaseCode = 5008 // 函数类型为空
	BaseCodeExpressionEmpty                BaseCode = 5009 // 表达式为空
	BaseCodeVideoDimensionTypeNotExist     BaseCode = 5010 // 视频维度不存在
	BaseCodeVideoDimensionTypeImplNotExist BaseCode = 5011 // 视频维度查询具体实现不存在
	BaseCodeRuleNotExist                   BaseCode = 511  // 规则不存在
	BaseCodeMetricNotExist                 BaseCode = 512  // 指标不存在
	BaseCodeDataSaveNotExist               BaseCode = 513  // 数据保存方式不存在
	BaseCodeDateTypeNotExist               BaseCode = 514  // 时间类型不存在
	BaseCodeMessageNotExist                BaseCode = 515  // 时间类型不存在
	BaseCodeReconciliationNotExist         BaseCode = 516  // 对账不存在
	BaseCodeCollectTypeHandlerNotExist     BaseCode = 517  // 采集类型处理器不存在
)

// Code 返回枚举的代码值
func (b BaseCode) Code() int {
	return int(b)
}

// Msg 返回枚举的描述信息
func (b BaseCode) Msg() string {
	switch b {
	case BaseCodeSuccess:
		return "OK"
	case BaseCodeSystemError:
		return "系统异常，请稍后重试"
	case BaseCodeParameterError:
		return "参数验证异常"
	case BaseCodeDataSourceCloseError:
		return "数据源关闭失败"
	case BaseCodeUserNotLogin:
		return "用户未登录"
	case BaseCodeUIDWorkIDError:
		return "uid_work_id设置失败"
	case BaseCodeCollectTypeNotExist:
		return "采集方式不存在"
	case BaseCodeMetricCollectType:
		return "指标收集类型错误"
	case BaseCodeCollectSourceNameEmpty:
		return "收集源头名字为空"
	case BaseCodeCollectTypeEmpty:
		return "收集方式为空"
	case BaseCodeCollectDetailEmpty:
		return "具体的收集实现为空"
	case BaseCodeMetricTimeFormatEmpty:
		return "统计的时间格式为空"
	case BaseCodeArgumentsEmpty:
		return "指标统计的参数为空"
	case BaseCodeFunctionTypeEmpty:
		return "函数类型为空"
	case BaseCodeExpressionEmpty:
		return "表达式为空"
	case BaseCodeVideoDimensionTypeNotExist:
		return "视频维度不存在"
	case BaseCodeVideoDimensionTypeImplNotExist:
		return "视频维度查询具体实现不存在"
	case BaseCodeRuleNotExist:
		return "规则不存在"
	case BaseCodeMetricNotExist:
		return "指标不存在"
	case BaseCodeDataSaveNotExist:
		return "数据保存方式不存在"
	case BaseCodeDateTypeNotExist:
		return "时间类型不存在"
	case BaseCodeMessageNotExist:
		return "时间类型不存在"
	case BaseCodeReconciliationNotExist:
		return "对账不存在"
	case BaseCodeCollectTypeHandlerNotExist:
		return "采集类型处理器不存在"
	default:
		return ""
	}
}

// GetBaseCodeMsg 根据code获取描述信息
func GetBaseCodeMsg(code int) string {
	return GetBaseCodeByCode(code).Msg()
}

// GetBaseCodeByCode 根据code获取枚举
func GetBaseCodeByCode(code int) BaseCode {
	switch code {
	case 0:
		return BaseCodeSuccess
	case -1:
		return BaseCodeSystemError
	case -2:
		return BaseCodeParameterError
	case -3:
		return BaseCodeDataSourceCloseError
	case 1001:
		return BaseCodeUserNotLogin
	case 5000:
		return BaseCodeUIDWorkIDError
	case 5001:
		return BaseCodeCollectTypeNotExist
	case 5002:
		return BaseCodeMetricCollectType
	case 5003:
		return BaseCodeCollectSourceNameEmpty
	case 5004:
		return BaseCodeCollectTypeEmpty
	case 5005:
		return BaseCodeCollectDetailEmpty
	case 5006:
		return BaseCodeMetricTimeFormatEmpty
	case 5007:
		return BaseCodeArgumentsEmpty
	case 5008:
		return BaseCodeFunctionTypeEmpty
	case 5009:
		return BaseCodeExpressionEmpty
	case 5010:
		return BaseCodeVideoDimensionTypeNotExist
	case 5011:
		return BaseCodeVideoDimensionTypeImplNotExist
	case 511:
		return BaseCodeRuleNotExist
	case 512:
		return BaseCodeMetricNotExist
	case 513:
		return BaseCodeDataSaveNotExist
	case 514:
		return BaseCodeDateTypeNotExist
	case 515:
		return BaseCodeMessageNotExist
	case 516:
		return BaseCodeReconciliationNotExist
	case 517:
		return BaseCodeCollectTypeHandlerNotExist
	default:
		return 0
	}
}

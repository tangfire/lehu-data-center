package model

import (
	"lehu-data-center/app/collect/service/internal/entity"
	"lehu-data-center/app/collect/service/internal/enums"
	"time"
)

type Metric struct {
	Id                int64                   `json:"id" gorm:"column:id"`                                   // id
	RuleId            int64                   `json:"rule_id" gorm:"column:rule_id"`                         // 规则id
	MetricName        string                  `json:"metric_name" gorm:"column:metric_name"`                 // 指标名字
	MetricDescribe    string                  `json:"metric_describe" gorm:"column:metric_describe"`         // 指标描述
	MetricCollectType enums.MetricCollectType `json:"metric_collect_type" gorm:"column:metric_collect_type"` // 指标收集的类型 1:基础型 2:计算型
	CollectSourceName string                  `json:"collect_source_name" gorm:"column:collect_source_name"` // 基础型指标会用到：收集的源头名字，如果是数据库，就是数据源名字
	CollectType       int8                    `json:"collect_type" gorm:"column:collect_type"`               // 基础型指标会用到：收集的方式 1：sql查询 2：调用接口   对应的枚举：CollectType
	CollectDetail     string                  `json:"collect_detail" gorm:"column:collect_detail"`           //  基础型指标会用到：具体的收集实现，如果是sql，就是sql脚本。如果是接口，就是url
	MetricTimeFormat  string                  `json:"metric_time_format" gorm:"column:metric_time_format"`   // 基础性指标会用到：统计的时间格式
	Arguments         string                  `json:"arguments" gorm:"column:arguments"`                     // 计算型指标会用到：指标统计的参数
	FunctionType      enums.FunctionType      `json:"function_type" gorm:"column:function_type"`             // 计算型指标会用到：指标统计函数，详情见FunctionType
	Expression        string                  `json:"expression" gorm:"column:expression"`                   // 计算型指标会用到：指标计算表达式
	ShowStatus        int                     `json:"show_status" gorm:"column:show_status"`                 // 指标是否显示 1:显示 0:隐藏
	CodeType          int                     `json:"code_type" gorm:"column:code_type"`                     // 指标数据对应的代码类型，详情见：CodeType
	MetricType        int                     `json:"metric_type" gorm:"column:metric_type"`                 // 类型：1抽取型 2计算型
	Sort              int                     `json:"sort" gorm:"column:sort"`                               // 排序
	Status            int                     `json:"status" gorm:"column:status"`                           // 状态 1:启用 0:禁用
	UpdateTime        time.Time               `json:"update_time" gorm:"column:update_time"`                 // 编辑时间
	CreateTime        time.Time               `json:"create_time" gorm:"column:create_time"`                 // 创建时间
}

func (d Metric) TableName() string {
	return "metric"
}

func (d *Metric) Data2Biz() *entity.Metric {
	if d == nil {
		return nil
	}
	return &entity.Metric{
		Id:                d.Id,
		RuleId:            d.RuleId,
		MetricName:        d.MetricName,
		MetricDescribe:    d.MetricDescribe,
		MetricCollectType: d.MetricCollectType,
		CollectSourceName: d.CollectSourceName,
		CollectType:       d.CollectType,
		CollectDetail:     d.CollectDetail,
		MetricTimeFormat:  d.MetricTimeFormat,
		Arguments:         d.Arguments,
		FunctionType:      d.FunctionType,
		Expression:        d.Expression,
		ShowStatus:        d.ShowStatus,
		CodeType:          d.CodeType,
		MetricType:        d.MetricType,
		Sort:              d.Sort,
		Status:            d.Status,
		UpdateTime:        d.UpdateTime,
		CreateTime:        d.CreateTime,
	}
}

func (d *Metric) Biz2Data(m *entity.Metric) {
	if d == nil {
		return
	}
	*d = Metric{
		Id:                m.Id,
		RuleId:            m.RuleId,
		MetricName:        m.MetricName,
		MetricDescribe:    m.MetricDescribe,
		MetricCollectType: m.MetricCollectType,
		CollectSourceName: m.CollectSourceName,
		CollectType:       m.CollectType,
		CollectDetail:     m.CollectDetail,
		MetricTimeFormat:  m.MetricTimeFormat,
		Arguments:         m.Arguments,
		FunctionType:      m.FunctionType,
		Expression:        m.Expression,
		ShowStatus:        m.ShowStatus,
		CodeType:          m.CodeType,
		MetricType:        m.MetricType,
		Sort:              m.Sort,
		Status:            m.Status,
		UpdateTime:        m.UpdateTime,
		CreateTime:        m.CreateTime,
	}
	return
}

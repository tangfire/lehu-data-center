package model

import (
	"lehu-data-center/app/collect/service/internal/entity"
	"time"
)

type Rule struct {
	Id            int64     `json:"id" gorm:"column:id"`                           // id
	RuleDescribe  string    `json:"rule_describe" gorm:"column:rule_describe"`     // 规则描述
	RuleName      string    `json:"rule_name" gorm:"column:rule_name"`             // 规则名字
	RuleType      int8      `json:"rule_type" gorm:"column:rule_type"`             // 规则类型 1收集 2 查询
	RuleVersionId int64     `json:"rule_version_id" gorm:"column:rule_version_id"` // 如果需要规则变更，直接创建一个新规则，将版本号递增
	Status        int       `json:"status" gorm:"column:status"`                   // 状态 1:启用 0:禁用
	UpdateTime    time.Time `json:"update_time" gorm:"column:update_time"`         // 编辑时间
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time"`         // 创建时间
}

func (r Rule) TableName() string {
	return "rule"
}

func (r *Rule) Data2Biz() *entity.Rule {
	if r == nil {
		return nil
	}
	return &entity.Rule{
		Id:            r.Id,
		RuleDescribe:  r.RuleDescribe,
		RuleName:      r.RuleName,
		RuleType:      r.RuleType,
		RuleVersionId: r.RuleVersionId,
		Status:        r.Status,
		UpdateTime:    r.UpdateTime,
		CreateTime:    r.CreateTime,
	}
}

func (r *Rule) Biz2Data(m *entity.Rule) {
	if m == nil {
		return
	}
	*r = Rule{
		Id:            m.Id,
		RuleDescribe:  m.RuleDescribe,
		RuleName:      m.RuleName,
		RuleType:      m.RuleType,
		RuleVersionId: m.RuleVersionId,
		Status:        m.Status,
		UpdateTime:    m.UpdateTime,
		CreateTime:    m.CreateTime,
	}
}

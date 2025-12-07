package entity

import "time"

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

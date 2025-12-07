package entity

import "time"

type DataSave struct {
	ID                int64     `json:"id" gorm:"column:id"`
	RuleID            int64     `json:"rule_id" gorm:"column:rule_id"`                         // 规则id
	DataSourceName    string    `json:"data_source_name" gorm:"column:data_source_name"`       // 数据源名称
	CreateTablePrefix string    `json:"create_table_prefix" gorm:"column:create_table_prefix"` // 创建表的前缀
	Status            int       `json:"status" gorm:"column:status"`                           // 状态 1:启用 0:禁用
	UpdateTime        time.Time `json:"update_time" gorm:"column:update_time"`                 // 编辑时间
	CreateTime        time.Time `json:"create_time" gorm:"column:create_time"`                 // 创建时间
}

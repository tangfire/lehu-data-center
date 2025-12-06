package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type DimensionGatherModel struct {
	Id                int64     `json:"id" gorm:"column:id"`                                   // id
	RuleId            int64     `json:"rule_id" gorm:"column:rule_id"`                         // 规则id
	CollectType       int8      `json:"collect_type" gorm:"column:collect_type"`               // 收集的方式 1：sql查询 2：调用接口
	CollectDetail     string    `json:"collect_detail" gorm:"column:collect_detail"`           //  具体的收集实现，如果是sql，就是sql脚本。如果是接口，就是url
	CollectSourceName string    `json:"collect_source_name" gorm:"column:collect_source_name"` // 收集的源头名字，如果是数据库，就是数据源名字
	Entity            string    `json:"entity" gorm:"column:entity"`                           // 表名
	Status            int       `json:"status" gorm:"column:status"`                           // 状态 1:启用 0:禁用
	UpdateTime        time.Time `json:"update_time" gorm:"column:update_time"`                 // 编辑时间
	CreateTime        time.Time `json:"create_time" gorm:"column:create_time"`                 // 创建时间
}

type DimensionGatherRepo interface {
	GetDimensionGatherByRuleIdAndEntity(context.Context, int64, string) (*DimensionGatherModel, error)
}

type DimensionGatherUsecase struct {
	repo     DimensionGatherRepo
	ruleRepo RuleRepo
	log      *log.Helper
}

func NewDimensionGatherUsecase(repo DimensionGatherRepo, logger log.Logger) *DimensionGatherUsecase {
	return &DimensionGatherUsecase{repo: repo, log: log.NewHelper(logger)}
}

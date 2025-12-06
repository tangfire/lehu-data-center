package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/biz"
	"lehu-data-center/app/collect/service/internal/data/model"
)

type ruleRepo struct {
	data *Data
	log  *log.Helper
}

func NewRuleRepo(data *Data, logger log.Logger) biz.RuleRepo {
	return &ruleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ruleRepo) GetRuleById(ctx context.Context, id int64) (*biz.RuleModel, error) {
	data := model.Rule{}
	err := r.data.db.Table(model.Rule{}.TableName()).Where("id = ?", id).First(&data).Error
	if err != nil {
		r.log.WithContext(ctx).Errorf("GetRuleById|First fail,data:%+v,err:%v", id, err)
		return nil, err
	}
	ret := data.Data2Biz()
	return ret, nil
}

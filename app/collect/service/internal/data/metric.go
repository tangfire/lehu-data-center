package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/biz"
	"lehu-data-center/app/collect/service/internal/data/model"
	"lehu-data-center/app/collect/service/internal/entity"
)

type metricRepo struct {
	data *Data
	log  *log.Helper
}

func NewMetricRepo(data *Data, logger log.Logger) biz.MetricRepo {
	return &metricRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *metricRepo) GetMetricListByRuleId(ctx context.Context, ruleId int64) ([]*entity.Metric, error) {
	var dataList []*model.Metric
	err := r.data.db.Table(model.Metric{}.TableName()).Where("rule_id = ?", ruleId).Find(&dataList).Error
	if err != nil {
		r.log.WithContext(ctx).Errorf("GetMetricListByRuleId|Find fail,ruleId:%v,err:%v", ruleId, err)
		return nil, err
	}
	var retList []*entity.Metric
	for _, data := range dataList {
		tmp := data.Data2Biz()
		retList = append(retList, tmp)
	}
	return retList, nil
}

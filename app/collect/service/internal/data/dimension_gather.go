package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/biz"
	"lehu-data-center/app/collect/service/internal/data/model"
	"lehu-data-center/app/collect/service/internal/enums"
)

type dimensionGatherRepo struct {
	data *Data
	log  *log.Helper
}

func NewDimensionGatherRepo(data *Data, logger log.Logger) biz.DimensionGatherRepo {
	return &dimensionGatherRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *dimensionGatherRepo) GetDimensionGatherByRuleIdAndEntity(ctx context.Context, ruleId int64, entity enums.EntityType) (*biz.DimensionGatherModel, error) {
	data := model.DimensionGather{}
	err := r.data.db.Table(model.DimensionGather{}.TableName()).
		Where("rule_id = ? and entity = ?", ruleId, entity).
		First(&data).Error
	if err != nil {
		r.log.WithContext(ctx).Errorf("GetDimensionGatherByRuleIdAndEntity|First fail,ruleId:%v,entity:%v, err:%v", ruleId, entity, err)
		return nil, err
	}
	ret := data.Data2Biz()
	return ret, nil
}

package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/biz"
	"lehu-data-center/app/collect/service/internal/data/model"
)

type videoBusinessMessageConsumerRecordRepo struct {
	data *Data
	log  *log.Helper
}

func NewVideoBusinessMessageConsumerRecordRepo(data *Data, logger log.Logger) biz.VideoBusinessMessageConsumerRecordRepo {
	return &videoBusinessMessageConsumerRecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *videoBusinessMessageConsumerRecordRepo) DeleteByCreateTime(createTime string) (int64, error) {
	result := r.data.db.Table(model.VideoBusinessMessageConsumerRecord{}.TableName()).
		Where("DATE_FORMAT(create_time,'%Y-%m-%d') = ?", createTime).
		Delete(&model.VideoBusinessMessageConsumerRecord{})

	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

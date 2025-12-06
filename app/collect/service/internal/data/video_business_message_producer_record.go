package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/biz"
	"lehu-data-center/app/collect/service/internal/data/model"
)

type videoBusinessMessageProducerRecordRepo struct {
	data *Data
	log  *log.Helper
}

func NewVideoBusinessMessageProducerRecordRepo(data *Data, logger log.Logger) biz.VideoBusinessMessageProducerRecordRepo {
	return &videoBusinessMessageProducerRecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *videoBusinessMessageProducerRecordRepo) DeleteByCreateTime(createTime string) (int64, error) {
	result := r.data.db.Table(model.VideoBusinessMessageProducerRecord{}.TableName()).
		Where("DATE_FORMAT(create_time,'%Y-%m-%d') = ?", createTime).
		Delete(&model.VideoBusinessMessageProducerRecord{})

	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

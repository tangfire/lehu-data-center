package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/biz"
	"lehu-data-center/app/collect/service/internal/data/model"
)

type messageConsumerRecordRepo struct {
	data *Data
	log  *log.Helper
}

func NewMessageConsumerRecordRepo(data *Data, logger log.Logger) biz.MessageConsumerRecordRepo {
	return &messageConsumerRecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *messageConsumerRecordRepo) DeleteByConsumerTime(consumerTime string) (int64, error) {
	result := r.data.db.Table(model.MessageConsumerRecord{}.TableName()).
		Where("DATE_FORMAT(consumer_time, '%Y-%m-%d') = ?", consumerTime).
		Delete(&model.MessageConsumerRecord{})

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

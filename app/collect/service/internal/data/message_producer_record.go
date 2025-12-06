package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/biz"
	"lehu-data-center/app/collect/service/internal/data/model"
)

type messageProducerRecordRepo struct {
	data *Data
	log  *log.Helper
}

func NewMessageProducerRecordRepo(data *Data, logger log.Logger) biz.MessageProducerRecordRepo {
	return &messageProducerRecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *messageProducerRecordRepo) DeleteBySendTime(sendTime string) (int64, error) {
	result := r.data.db.Table(model.MessageProducerRecord{}.TableName()).
		Where("DATE_FORMAT(send_time,'%Y-%m-%d') = ?", sendTime).
		Delete(&model.MessageProducerRecord{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type VideoBusinessMessageProducerRecordModel struct {
	Id                      int64     `json:"id" gorm:"column:id"`
	MessageProducerRecordId int64     `json:"message_producer_record_id" gorm:"column:message_producer_record_id"` // 消息发送记录id
	VideoDimensionType      int       `json:"video_dimension_type" gorm:"column:video_dimension_type"`             // 视频维度分类，1-父级视频分类 2-视频分类 3-视频本身
	DateType                int       `json:"date_type" gorm:"column:date_type"`                                   // 日期类型
	Status                  int       `json:"status" gorm:"column:status"`                                         // 状态 1:启用 0:禁用
	UpdateTime              time.Time `json:"update_time" gorm:"column:update_time"`                               // 编辑时间
	CreateTime              time.Time `json:"create_time" gorm:"column:create_time"`                               // 创建时间
}

type VideoBusinessMessageProducerRecordRepo interface {
	DeleteByCreateTime(string) (int64, error)
}

type VideoBusinessMessageProducerRecordUsecase struct {
	repo VideoBusinessMessageProducerRecordRepo
	log  *log.Helper
}

func NewVideoBusinessMessageProducerRecordUsecase(repo VideoBusinessMessageProducerRecordRepo, logger log.Logger) *VideoBusinessMessageProducerRecordUsecase {
	return &VideoBusinessMessageProducerRecordUsecase{repo: repo, log: log.NewHelper(logger)}
}

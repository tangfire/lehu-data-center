package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type MessageConsumerRecordModel struct {
	Id                       int64     `json:"id" gorm:"column:id"`
	MessageParentTraceId     int64     `json:"message_parent_trace_id" gorm:"column:message_parent_trace_id"`       // 消息的父级链路id
	MessageTraceId           int64     `json:"message_trace_id" gorm:"column:message_trace_id"`                     // 消息的链路id
	MessageId                int64     `json:"message_id" gorm:"column:message_id"`                                 // 消息id
	MessageContent           string    `json:"message_content" gorm:"column:message_content"`                       // 消息内容
	MessageConsumerException string    `json:"message_consumer_exception" gorm:"column:message_consumer_exception"` // 消息消费失败的异常信息
	MessageConsumerStatus    int       `json:"message_consumer_status" gorm:"column:message_consumer_status"`       // 消息消费状态 1:未消费 -1:消费失败 2:消费成功
	MessageConsumerCount     int       `json:"message_consumer_count" gorm:"column:message_consumer_count"`         // 消息的消费次数
	ReconciliationStatus     int       `json:"reconciliation_status" gorm:"column:reconciliation_status"`           // 消息对账状态 1:未对账 -1:对账完成有问题 2:对账完成没有问题 3:对账有问题处理完毕
	ConsumerTime             time.Time `json:"consumer_time" gorm:"column:consumer_time"`                           // 消息发送时间
	Status                   int       `json:"status" gorm:"column:status"`                                         // 状态 1:启用 0:禁用
	UpdateTime               time.Time `json:"update_time" gorm:"column:update_time"`                               // 编辑时间
	CreateTime               time.Time `json:"create_time" gorm:"column:create_time"`                               // 创建时间
}

type MessageConsumerRecordRepo interface {
	DeleteByConsumerTime(consumerTime string) (int64, error)
}

type MessageConsumerRecordUsecase struct {
	repo MessageConsumerRecordRepo
	log  *log.Helper
}

func NewMessageConsumerRecordUsecase(repo MessageConsumerRecordRepo, logger log.Logger) *MessageConsumerRecordUsecase {
	return &MessageConsumerRecordUsecase{repo: repo, log: log.NewHelper(logger)}
}

package model

import "time"

type MessageProducerRecord struct {
	Id                   int64     `json:"id" gorm:"column:id"`
	MessageParentTraceId int64     `json:"message_parent_trace_id" gorm:"column:message_parent_trace_id"` // 消息的父级链路id
	MessageTraceId       int64     `json:"message_trace_id" gorm:"column:message_trace_id"`               // 消息的链路id
	MessageId            int64     `json:"message_id" gorm:"column:message_id"`                           // 消息id
	MessageContent       string    `json:"message_content" gorm:"column:message_content"`                 // 消息内容
	MessageSendException string    `json:"message_send_exception" gorm:"column:message_send_exception"`   // 消息发送失败的异常消息
	MessageSendStatus    int       `json:"message_send_status" gorm:"column:message_send_status"`         // 消息发送状态 1-未发送 -1-发送失败 3-发送成功
	ReconciliationStatus int       `json:"reconciliation_status" gorm:"column:reconciliation_status"`     // 消息对账状态 1-未对账 -1-对账完成有问题 2-对账完成没问题 3-对账有问题处理完毕
	SendTime             time.Time `json:"send_time" gorm:"column:send_time"`                             // 消息发送时间
	Status               int       `json:"status" gorm:"column:status"`                                   // 状态 1-启用 0-禁用
	UpdateTime           time.Time `json:"update_time" gorm:"column:update_time"`                         // 更新时间
	CreateTime           time.Time `json:"create_time" gorm:"column:create_time"`                         // 创建时间
}

func (m MessageProducerRecord) TableName() string {
	return "message_producer_record"
}

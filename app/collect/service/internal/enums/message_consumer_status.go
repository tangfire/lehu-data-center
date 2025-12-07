// internal/pkg/enums/message_consumer_status.go
package enums

// MessageConsumerStatus 消息消费状态枚举
type MessageConsumerStatus int

const (
	MessageConsumerStatusUnconsumed      MessageConsumerStatus = 1  // 未消费
	MessageConsumerStatusConsumerFail    MessageConsumerStatus = -1 // 消费失败
	MessageConsumerStatusConsumerSuccess MessageConsumerStatus = 2  // 消费成功
)

// Code 返回枚举的代码值
func (m MessageConsumerStatus) Code() int {
	return int(m)
}

// Msg 返回枚举的描述信息
func (m MessageConsumerStatus) Msg() string {
	switch m {
	case MessageConsumerStatusUnconsumed:
		return "未消费"
	case MessageConsumerStatusConsumerFail:
		return "消费失败"
	case MessageConsumerStatusConsumerSuccess:
		return "消费成功"
	default:
		return ""
	}
}

// GetMsg 根据code获取描述信息
func GetMessageConsumerStatusMsg(code int) string {
	return GetMessageConsumerStatusByCode(code).Msg()
}

// GetMessageConsumerStatusByCode 根据code获取枚举
func GetMessageConsumerStatusByCode(code int) MessageConsumerStatus {
	switch code {
	case 1:
		return MessageConsumerStatusUnconsumed
	case -1:
		return MessageConsumerStatusConsumerFail
	case 2:
		return MessageConsumerStatusConsumerSuccess
	default:
		return 0
	}
}

// internal/pkg/enums/message_send_status.go
package enums

// MessageSendStatus 消息发送状态枚举
type MessageSendStatus int

const (
	MessageSendStatusUnsent      MessageSendStatus = 1  // 未发送
	MessageSendStatusSendFail    MessageSendStatus = -1 // 发送失败
	MessageSendStatusSendSuccess MessageSendStatus = 2  // 发送成功
)

// Code 返回枚举的代码值
func (m MessageSendStatus) Code() int {
	return int(m)
}

// Msg 返回枚举的描述信息
func (m MessageSendStatus) Msg() string {
	switch m {
	case MessageSendStatusUnsent:
		return "未发送"
	case MessageSendStatusSendFail:
		return "发送失败"
	case MessageSendStatusSendSuccess:
		return "发送成功"
	default:
		return ""
	}
}

// GetMessageSendStatusMsg 根据code获取描述信息
func GetMessageSendStatusMsg(code int) string {
	return GetMessageSendStatusByCode(code).Msg()
}

// GetMessageSendStatusByCode 根据code获取枚举
func GetMessageSendStatusByCode(code int) MessageSendStatus {
	switch code {
	case 1:
		return MessageSendStatusUnsent
	case -1:
		return MessageSendStatusSendFail
	case 2:
		return MessageSendStatusSendSuccess
	default:
		return 0
	}
}

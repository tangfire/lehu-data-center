// transfers/rule_handle_output.go
package transfers

// RuleHandleOutput 规则处理输出
type RuleHandleOutput struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewSuccessRuleHandleOutput 创建成功的规则处理输出
func NewSuccessRuleHandleOutput(message string, data interface{}) *RuleHandleOutput {
	return &RuleHandleOutput{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// NewErrorRuleHandleOutput 创建错误的规则处理输出
func NewErrorRuleHandleOutput(message string) *RuleHandleOutput {
	return &RuleHandleOutput{
		Success: false,
		Message: message,
	}
}

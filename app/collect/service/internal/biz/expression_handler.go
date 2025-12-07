package biz

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"lehu-data-center/app/collect/service/internal/enums"
	"strconv"
)

// ExpressionHandler 表达式处理器
type ExpressionHandler struct {
	luaState *lua.LState
}

func NewExpressionHandler() *ExpressionHandler {
	l := lua.NewState()
	return &ExpressionHandler{
		luaState: l,
	}
}

// HandleExpression 处理表达式
func (h *ExpressionHandler) HandleExpression(
	codeClassType interface{},
	expression string,
	params map[string]interface{},
	defaultValue interface{},
) (interface{}, error) {

	defer func() {
		if r := recover(); r != nil {
			// 表达式计算失败，返回默认值
		}
	}()

	// 使用Lua引擎执行表达式
	h.luaState.SetGlobal("params", h.convertToLuaTable(params))

	if err := h.luaState.DoString(fmt.Sprintf("result = %s", expression)); err != nil {
		return defaultValue, fmt.Errorf("execute expression failed: %w", err)
	}

	result := h.luaState.GetGlobal("result")
	return h.convertFromLuaValue(result, codeClassType, defaultValue)
}

// GetDefaultValue 获取数据类型的默认值
func (h *ExpressionHandler) GetDefaultValue(codeType enums.CodeType) interface{} {
	switch codeType {
	case enums.CodeTypeInteger:
		return 0
	case enums.CodeTypeLong:
		return int64(0)
	case enums.CodeTypeDouble:
		return 0.0
	case enums.CodeTypeString:
		return ""
	default:
		return nil
	}
}

func (h *ExpressionHandler) convertToLuaTable(params map[string]interface{}) *lua.LTable {
	table := h.luaState.NewTable()
	for k, v := range params {
		h.luaState.SetField(table, k, h.convertToLuaValue(v))
	}
	return table
}

func (h *ExpressionHandler) convertToLuaValue(v interface{}) lua.LValue {
	switch val := v.(type) {
	case int:
		return lua.LNumber(val)
	case int64:
		return lua.LNumber(val)
	case float64:
		return lua.LNumber(val)
	case string:
		return lua.LString(val)
	case bool:
		return lua.LBool(val)
	default:
		return lua.LNil
	}
}

func (h *ExpressionHandler) convertFromLuaValue(
	lv lua.LValue,
	codeClassType interface{},
	defaultValue interface{},
) (interface{}, error) {

	switch val := lv.(type) {
	case lua.LNumber:
		num := float64(val)
		switch codeClassType.(type) {
		case int:
			return int(num), nil
		case int64:
			return int64(num), nil
		case float64:
			return num, nil
		}
	case lua.LString:
		str := string(val)
		if codeClassType == "string" {
			return str, nil
		}
		// 尝试转换
		switch codeClassType.(type) {
		case int:
			if i, err := strconv.Atoi(str); err == nil {
				return i, nil
			}
		case int64:
			if i, err := strconv.ParseInt(str, 10, 64); err == nil {
				return i, nil
			}
		case float64:
			if f, err := strconv.ParseFloat(str, 64); err == nil {
				return f, nil
			}
		}
	}

	return defaultValue, nil
}

func (h *ExpressionHandler) Close() {
	if h.luaState != nil {
		h.luaState.Close()
	}
}

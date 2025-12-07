// internal/pkg/enums/order_status.go
package enums

// OrderStatus 订单状态
type OrderStatus int

const (
	OrderStatusNoPay  OrderStatus = 1 // 未支付
	OrderStatusCancel OrderStatus = 2 // 已取消
	OrderStatusPay    OrderStatus = 3 // 已支付
	OrderStatusRefund OrderStatus = 4 // 已退单
)

// Code 返回枚举的代码值
func (o OrderStatus) Code() int {
	return int(o)
}

// Msg 返回枚举的描述信息
func (o OrderStatus) Msg() string {
	switch o {
	case OrderStatusNoPay:
		return "未支付"
	case OrderStatusCancel:
		return "已取消"
	case OrderStatusPay:
		return "已支付"
	case OrderStatusRefund:
		return "已退单"
	default:
		return ""
	}
}

// GetOrderStatusMsg 根据code获取描述信息
func GetOrderStatusMsg(code int) string {
	return GetOrderStatusByCode(code).Msg()
}

// GetOrderStatusByCode 根据code获取枚举
func GetOrderStatusByCode(code int) OrderStatus {
	switch code {
	case 1:
		return OrderStatusNoPay
	case 2:
		return OrderStatusCancel
	case 3:
		return OrderStatusPay
	case 4:
		return OrderStatusRefund
	default:
		return 0
	}
}

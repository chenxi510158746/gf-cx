// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Orders is the golang structure of table c_orders for DAO operations like Where/Data.
type Orders struct {
	g.Meta     `orm:"table:c_orders, do:true"`
	Id         interface{} // 订单ID
	OrderSn    interface{} // 订单编号
	UserId     interface{} // 用户ID
	GoodsId    interface{} // 商品ID
	GoodsPrice interface{} // 商品价格（单位：分）
	PayAmount  interface{} // 支付金额（单位：分）
	PayStatus  interface{} // 支付状态（1-待支付；2-已支付；3-已退款）
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 删除时间
}

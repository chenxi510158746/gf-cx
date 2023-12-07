// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Orders is the golang structure for table orders.
type Orders struct {
	Id         uint        `json:"id"         description:"订单ID"`
	OrderSn    string      `json:"orderSn"    description:"订单编号"`
	UserId     int         `json:"userId"     description:"用户ID"`
	GoodsId    int         `json:"goodsId"    description:"商品ID"`
	GoodsPrice int         `json:"goodsPrice" description:"商品价格（单位：分）"`
	PayAmount  int         `json:"payAmount"  description:"支付金额（单位：分）"`
	PayStatus  int         `json:"payStatus"  description:"支付状态（1-待支付；2-已支付；3-已退款）"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"更新时间"`
	DeletedAt  *gtime.Time `json:"deletedAt"  description:"删除时间"`
}

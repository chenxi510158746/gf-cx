package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type OrderCreateItem struct {
	OrderSn    string
	UserId     uint
	GoodsId    uint
	GoodsPrice float32
	PayAmount  float32
	PayStatus  uint
}

type OrderGetOneOut struct {
	Order *OrderOutItem `json:"order" dc:"订单详情"`
}

type OrderGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

type OrderGetListOut struct {
	List  []*OrderOutItem `json:"list" dc:"订单列表"`
	Page  int             `json:"page" dc:"分页码"`
	Size  int             `json:"size" dc:"分页数"`
	Total int             `json:"total" dc:"总数量"`
}

type OrderOutItem struct {
	gmeta.Meta     `orm:"table:c_orders"`
	Id             uint          `json:"id" dc:"订单ID"`
	OrderSn        string        `json:"order_sn" dc:"订单编号"`
	UserId         uint          `json:"user_id" dc:"用户ID"`
	GoodsId        uint          `json:"goods_id" dc:"商品ID"`
	GoodsPrice     float32       `json:"goods_price" dc:"商品价格"`
	PayAmount      float32       `json:"pay_amount" dc:"订单价格"`
	GoodsPriceText string        `json:"goods_price_text" dc:"商品价格"`
	PayAmountText  string        `json:"pay_amount_text" dc:"订单价格"`
	PayStatus      uint          `json:"pay_status" dc:"支付状态"`
	PayStatusText  string        `json:"pay_status_text" dc:"支付状态描述"`
	CreatedAt      *gtime.Time   `json:"created_at" dc:"创建时间"` // 创建时间
	UpdatedAt      *gtime.Time   `json:"updated_at" dc:"更新时间"` // 更新时间
	Goods          *GoodsOutItem `orm:"with:id=goods_id" json:"goods"`
	User           *UserOutItem  `orm:"with:id=user_id" json:"user"`
}

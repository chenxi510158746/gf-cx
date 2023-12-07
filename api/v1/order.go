package v1

import (
	"gf-cx/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type OrderCreateReq struct {
	g.Meta  `path:"/order/create" method:"post" tags:"订单" summary:"订单创建"`
	GoodsId uint `json:"goods_id" v:"required#商品ID不能为空" dc:"商品ID"`
}

type OrderCreateRes struct {
	OrderId uint `json:"order_id" dc:"订单ID"`
}

type OrderGetOneReq struct {
	g.Meta  `path:"/order/info" method:"get" tags:"订单" summary:"获取订单详情"`
	OrderId uint `json:"order_id" v:"required#订单ID不能为空" dc:"订单ID"`
}

type OrderGetOneRes struct {
	Order *model.OrderOutItem `json:"order" dc:"订单详情"`
}

type OrderGetListReq struct {
	g.Meta `path:"/order/list" method:"get" tags:"订单" summary:"获取订单列表"`
	Page   int `json:"page" d:"1" v:"min:1" dc:"分页码"`
	Size   int `json:"size" d:"10" v:"max:50" dc:"分页数量"`
}

type OrderGetListRes struct {
	List  []*model.OrderOutItem `json:"list" dc:"订单列表"`
	Page  int                   `json:"page" dc:"分页码"`
	Size  int                   `json:"size" dc:"分页数"`
	Total int                   `json:"total" dc:"总数量"`
}

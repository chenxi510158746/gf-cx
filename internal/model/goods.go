package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type GoodsCreateInput struct {
	Title  string
	Desc   string
	Price  float32
	Status uint
}

type GoodsUpdateInput struct {
	GoodsId uint
	Title   string
	Desc    string
	Price   float32
	Status  uint
}

type GoodsGetOneOut struct {
	GoodsInfo *GoodsOutItem `json:"goods_info"`
}

type GoodsGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

type GoodsGetListOut struct {
	List  []*GoodsOutItem `json:"list" description:"列表"`
	Page  int             `json:"page" description:"分页码"`
	Size  int             `json:"size" description:"分页数量"`
	Total int             `json:"total" description:"数据总数"`
}

type GoodsOutItem struct {
	gmeta.Meta `orm:"table:c_goods"`
	Id         uint        `json:"id" dc:"商品ID"`
	Title      string      `json:"title" dc:"商品名称"`
	Desc       string      `json:"desc" dc:"商品描述"`
	Price      float32     `json:"price" dc:"商品价格"`
	PriceText  string      `json:"price_text" dc:"商品价格"`
	Status     uint        `json:"status" dc:"状态"`
	StatusText string      `json:"status_text" dc:"状态描述"`
	CreatedAt  *gtime.Time `json:"created_at" dc:"创建时间"` // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at" dc:"修改时间"` // 修改时间
}

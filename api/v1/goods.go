package v1

import (
	"gf-cx/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type GoodsCreateReq struct {
	g.Meta `path:"/goods/create" method:"post" tags:"商品" summary:"商品创建"`
	Title  string  `v:"required|length:1,150#标题不能为空|标题长度为1～150" json:"title" dc:"名称"`
	Desc   string  `v:"required|length:1,300#描述不能为空|描述长度为1～300" json:"desc" dc:"描述"`
	Price  float32 `v:"required#价格不能为空" json:"price" dc:"价额"`
	Status uint    `v:"in:0,1#状态为0或1" json:"status" dc:"状态"`
}

type GoodsCreateRes struct {
	GoodsId uint `json:"goods_id"`
}

type GoodsUpdateReq struct {
	g.Meta  `path:"/goods/update" method:"post" tags:"商品" summary:"商品更新"`
	GoodsId uint    `v:"required#商品ID不能为空" json:"goods_id" dc:"商品ID"`
	Title   string  `v:"required|length:1,150#标题不能为空|标题长度为1～150" json:"title" dc:"名称"`
	Desc    string  `v:"required|length:1,300#描述不能为空|描述长度为1～300" json:"desc" dc:"描述"`
	Price   float32 `v:"required#价格不能为空" json:"price" dc:"价额"`
	Status  uint    `v:"in:0,1#状态为0或1" json:"status" dc:"状态"`
}

type GoodsUpdateRes struct{}

type GoodsDeleteReq struct {
	g.Meta  `path:"/goods/delete" method:"post" tags:"商品" summary:"商品删除"`
	GoodsId uint `v:"required#商品ID不能为空" json:"goods_id" dc:"商品ID"`
}

type GoodsDeleteRes struct{}

type GoodsGetOneReq struct {
	g.Meta  `path:"/goods/info" method:"get" tags:"商品" summary:"获取商品详情"`
	GoodsId uint `v:"required#商品ID不能为空" json:"goods_id" dc:"商品ID"`
}

type GoodsGetOneRes struct {
	GoodsInfo *model.GoodsOutItem `json:"goods_info"`
}

type GoodsGetListReq struct {
	g.Meta `path:"/goods/list" method:"get" tags:"商品" summary:"获取商品列表"`
	Page   int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	Size   int `json:"size" in:"query" d:"10" v:"max:50#分页数量最大50条" dc:"分页数量，最大50"`
}

type GoodsGetListRes struct {
	List  []*model.GoodsOutItem `json:"list" description:"列表"`
	Page  int                   `json:"page" description:"分页码"`
	Size  int                   `json:"size" description:"分页数量"`
	Total int                   `json:"total" description:"数据总数"`
}

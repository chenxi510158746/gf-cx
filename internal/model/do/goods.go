// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Goods is the golang structure of table c_goods for DAO operations like Where/Data.
type Goods struct {
	g.Meta    `orm:"table:c_goods, do:true"`
	Id        interface{} // 商品ID
	Title     interface{} // 商品名称
	Desc      interface{} // 商品描述
	Price     interface{} // 商品价格（单位：分）
	Status    interface{} // 商品状态（0-下架；1-上架）
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}

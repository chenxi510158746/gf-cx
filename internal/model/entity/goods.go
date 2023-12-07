// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Goods is the golang structure for table goods.
type Goods struct {
	Id        uint        `json:"id"        description:"商品ID"`
	Title     string      `json:"title"     description:"商品名称"`
	Desc      string      `json:"desc"      description:"商品描述"`
	Price     int         `json:"price"     description:"商品价格（单位：分）"`
	Status    int         `json:"status"    description:"商品状态（0-下架；1-上架）"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}

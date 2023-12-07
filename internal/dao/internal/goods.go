// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoodsDao is the data access object for table c_goods.
type GoodsDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns GoodsColumns // columns contains all the column names of Table for convenient usage.
}

// GoodsColumns defines and stores column names for table c_goods.
type GoodsColumns struct {
	Id        string // 商品ID
	Title     string // 商品名称
	Desc      string // 商品描述
	Price     string // 商品价格（单位：分）
	Status    string // 商品状态（0-下架；1-上架）
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// goodsColumns holds the columns for table c_goods.
var goodsColumns = GoodsColumns{
	Id:        "id",
	Title:     "title",
	Desc:      "desc",
	Price:     "price",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewGoodsDao creates and returns a new DAO object for table data access.
func NewGoodsDao() *GoodsDao {
	return &GoodsDao{
		group:   "default",
		table:   "c_goods",
		columns: goodsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GoodsDao) Columns() GoodsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GoodsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

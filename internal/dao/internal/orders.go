// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrdersDao is the data access object for table c_orders.
type OrdersDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns OrdersColumns // columns contains all the column names of Table for convenient usage.
}

// OrdersColumns defines and stores column names for table c_orders.
type OrdersColumns struct {
	Id         string // 订单ID
	OrderSn    string // 订单编号
	UserId     string // 用户ID
	GoodsId    string // 商品ID
	GoodsPrice string // 商品价格（单位：分）
	PayAmount  string // 支付金额（单位：分）
	PayStatus  string // 支付状态（1-待支付；2-已支付；3-已退款）
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	DeletedAt  string // 删除时间
}

// ordersColumns holds the columns for table c_orders.
var ordersColumns = OrdersColumns{
	Id:         "id",
	OrderSn:    "order_sn",
	UserId:     "user_id",
	GoodsId:    "goods_id",
	GoodsPrice: "goods_price",
	PayAmount:  "pay_amount",
	PayStatus:  "pay_status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewOrdersDao creates and returns a new DAO object for table data access.
func NewOrdersDao() *OrdersDao {
	return &OrdersDao{
		group:   "default",
		table:   "c_orders",
		columns: ordersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrdersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrdersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrdersDao) Columns() OrdersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrdersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrdersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrdersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

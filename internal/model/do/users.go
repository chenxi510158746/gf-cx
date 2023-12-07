// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table c_users for DAO operations like Where/Data.
type Users struct {
	g.Meta    `orm:"table:c_users, do:true"`
	Id        interface{} // 用户ID
	UserName  interface{} // 用户姓名
	Mobile    interface{} // 手机号码
	Password  interface{} // 密码
	Age       interface{} // 年龄
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}

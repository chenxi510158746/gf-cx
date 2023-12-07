// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id        uint        `json:"id"        description:"用户ID"`
	UserName  string      `json:"userName"  description:"用户姓名"`
	Mobile    string      `json:"mobile"    description:"手机号码"`
	Password  string      `json:"password"  description:"密码"`
	Age       int         `json:"age"       description:"年龄"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}

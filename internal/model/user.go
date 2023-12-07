package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type UserSignUpInput struct {
	UserName   string
	Mobile     string
	Password   string
	RePassword string
	Age        int
}

type UserSignInInput struct {
	UserName string
	Password string
}

type UserInfoOut struct {
	User *UserOutItem `json:"user" dc:"用户信息"`
}

type UserOutItem struct {
	gmeta.Meta `orm:"table:c_users"`
	Id         uint        `json:"id" dc:"用户ID"`
	UserName   string      `json:"user_name" dc:"姓名"`
	Mobile     string      `json:"mobile" dc:"手机号"`
	Age        int         `json:"age" dc:"年龄"`
	CreatedAt  *gtime.Time `json:"created_at" dc:"创建时间"` // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at" dc:"更新时间"` // 更新时间
}

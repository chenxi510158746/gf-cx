package v1

import (
	"gf-cx/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type UserSignUpReq struct {
	g.Meta     `path:"/user/sign-up" method:"post" tags:"用户" summary:"用户注册"`
	UserName   string `v:"required|length:1,15#用户名不能为空|用户名长度为1～15" json:"user_name" dc:"用户名"`
	Mobile     string `v:"required|phone#手机号不能为空|手机号格式错误" json:"mobile" dc:"手机号"`
	Password   string `v:"required|length:6,15#密码不能为空|密码长度为6～15" json:"password" dc:"密码"`
	RePassword string `v:"required|length:6,15|same:Password#重复密码不能为空|重复密码长度为6～15|与密码不一致" json:"re_password" dc:"重复密码"`
	Age        int    `json:"age" dc:"年龄"`
}

type UserSignUpRes struct {
	UserId uint `json:"user_id" dc:"用户ID"`
}

type UserSignInReq struct {
	g.Meta   `path:"/user/sign-in" method:"post" tags:"用户" summary:"用户登陆"`
	UserName string `v:"required#用户名不能为空" json:"user_name" dc:"用户名"`
	Password string `v:"required#密码不能为空" json:"password" dc:"密码"`
}

type UserSignInRes struct{}

type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"用户" summary:"获取用户信息"`
}

type UserInfoRes struct {
	User *model.UserOutItem `json:"user" dc:"用户信息"`
}

type UserSignOutReq struct {
	g.Meta `path:"/user/sign-out" method:"post" tags:"用户" summary:"用户退出"`
}

type UserSignOutRes struct{}

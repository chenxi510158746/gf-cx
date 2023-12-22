package v1

import "github.com/gogf/gf/v2/frame/g"

type GameStartReq struct {
	g.Meta  `path:"/game/start" method:"post" tags:"游戏" summary:"游戏启动"`
	Secret  string `v:"required#参数错误" json:"secret"`
	OrderId int    `v:"required#参数错误" json:"order_id"`
}

type GameStartRes struct {
	Msg string `json:"msg"`
}

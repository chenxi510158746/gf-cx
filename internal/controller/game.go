package controller

import (
	"context"
	"crypto/tls"
	"fmt"
	v1 "gf-cx/api/v1"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var Game = cGame{}

type cGame struct {
}

func (c *cGame) Start(ctx context.Context, req *v1.GameStartReq) (res *v1.GameStartRes, err error) {
	client := gclient.NewWebSocket()
	client.HandshakeTimeout = time.Second    // 设置超时时间
	client.Proxy = http.ProxyFromEnvironment // 设置代理
	client.TLSClientConfig = &tls.Config{}   // 设置 tls 配置
	conn, _, err := client.Dial("ws://union.gtslfw.cn:12234", nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	tJson := gjson.New(fmt.Sprintf(`{"type":"start_game", "secret":"%v", "order_id":"%v"}`, req.Secret, req.OrderId))
	err = conn.WriteMessage(websocket.TextMessage, []byte(tJson.MustToJsonString()))
	if err != nil {
		panic(err)
	}
	mt, data, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	fmt.Println(mt, string(data))
	return &v1.GameStartRes{string(data)}, nil
}

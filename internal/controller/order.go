package controller

import (
	"context"
	v1 "gf-cx/api/v1"
	"gf-cx/internal/model"
	"gf-cx/internal/service"
)

var Order = cOrder{}

type cOrder struct{}

func (o *cOrder) Create(ctx context.Context, req *v1.OrderCreateReq) (res *v1.OrderCreateRes, err error) {
	orderId, err := service.Order().Create(ctx, req.GoodsId)
	if err != nil {
		return nil, err
	}
	return &v1.OrderCreateRes{OrderId: orderId}, err
}

func (o *cOrder) GetOne(ctx context.Context, req *v1.OrderGetOneReq) (res *v1.OrderGetOneRes, err error) {
	out, err := service.Order().GetOne(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	return &v1.OrderGetOneRes{
		Order: out.Order,
	}, err
}

func (o *cOrder) GetList(ctx context.Context, req *v1.OrderGetListReq) (res *v1.OrderGetListRes, err error) {
	out, err := service.Order().GetList(ctx, model.OrderGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.OrderGetListRes{
		List:  out.List,
		Page:  out.Page,
		Size:  out.Size,
		Total: out.Total,
	}, err
}

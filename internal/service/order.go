package service

import (
	"context"
	"gf-cx/internal/model"
)

type IOrder interface {
	Create(ctx context.Context, goodsId uint) (orderId uint, err error)
	GetOne(ctx context.Context, orderId uint) (out *model.OrderGetOneOut, err error)
	GetList(ctx context.Context, req model.OrderGetListInput) (out *model.OrderGetListOut, err error)
	OutFormat(ctx context.Context, order *model.OrderOutItem) error
}

var localOrder IOrder

func Order() IOrder {
	if localGoods == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localOrder
}

func RegisterOrder(i IOrder) {
	localOrder = i
}

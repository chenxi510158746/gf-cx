package service

import (
	"context"
	"gf-cx/internal/model"
)

type IGoods interface {
	Create(ctx context.Context, req model.GoodsCreateInput) (goodsId uint, err error)
	Update(ctx context.Context, req model.GoodsUpdateInput) error
	Delete(ctx context.Context, goodsId uint) error
	GetOne(ctx context.Context, goodsId uint) (out *model.GoodsGetOneOut, err error)
	GetList(ctx context.Context, req model.GoodsGetListInput) (out *model.GoodsGetListOut, err error)
	OutFormat(ctx context.Context, goods *model.GoodsOutItem) error
}

var localGoods IGoods

func Goods() IGoods {
	if localGoods == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localGoods
}

func RegisterGoods(i IGoods) {
	localGoods = i
}

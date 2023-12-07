package controller

import (
	"context"
	v1 "gf-cx/api/v1"
	"gf-cx/internal/model"
	"gf-cx/internal/service"
)

var Goods = cGoods{}

type cGoods struct {
}

func (g *cGoods) Create(ctx context.Context, req *v1.GoodsCreateReq) (res *v1.GoodsCreateRes, err error) {
	goodsId, err := service.Goods().Create(ctx, model.GoodsCreateInput{
		Title:  req.Title,
		Desc:   req.Desc,
		Price:  req.Price * 100,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GoodsCreateRes{GoodsId: goodsId}, err
}

func (g *cGoods) Update(ctx context.Context, req *v1.GoodsUpdateReq) (res *v1.GoodsUpdateRes, err error) {
	err = service.Goods().Update(ctx, model.GoodsUpdateInput{
		GoodsId: req.GoodsId,
		Title:   req.Title,
		Desc:    req.Desc,
		Price:   req.Price * 100,
		Status:  req.Status,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (g *cGoods) Delete(ctx context.Context, req *v1.GoodsDeleteReq) (res *v1.GoodsDeleteRes, err error) {
	err = service.Goods().Delete(ctx, req.GoodsId)
	if err != nil {
		return nil, err
	}
	return
}

func (g *cGoods) GetOne(ctx context.Context, req *v1.GoodsGetOneReq) (res *v1.GoodsGetOneRes, err error) {
	out, err := service.Goods().GetOne(ctx, req.GoodsId)
	if err != nil {
		return nil, err
	}
	return &v1.GoodsGetOneRes{GoodsInfo: out.GoodsInfo}, err
}

func (g *cGoods) GetList(ctx context.Context, req *v1.GoodsGetListReq) (res *v1.GoodsGetListRes, err error) {
	out, err := service.Goods().GetList(ctx, model.GoodsGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GoodsGetListRes{
		List:  out.List,
		Page:  out.Page,
		Size:  out.Size,
		Total: out.Total,
	}, err
}

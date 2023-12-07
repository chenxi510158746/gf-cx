package logic

import (
	"context"
	"fmt"
	"gf-cx/internal/consts"
	"gf-cx/internal/dao"
	"gf-cx/internal/model"
	"gf-cx/internal/model/do"
	"gf-cx/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
)

type sGoods struct{}

func NewGoods() service.IGoods {
	return &sGoods{}
}

func init() {
	service.RegisterGoods(NewGoods())
}

func (s *sGoods) Create(ctx context.Context, in model.GoodsCreateInput) (goodsId uint, err error) {
	id, err := dao.Goods.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return 0, err
	}
	return uint(id), err
}

func (s *sGoods) Update(ctx context.Context, in model.GoodsUpdateInput) error {
	return dao.Goods.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.Goods.Ctx(ctx).Data(in).
			Where(do.Goods{Id: in.GoodsId}).
			Update()
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *sGoods) Delete(ctx context.Context, goodsId uint) error {
	return dao.Goods.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.Goods.Ctx(ctx).
			Where(do.Goods{Id: goodsId}).
			Delete()
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *sGoods) GetOne(ctx context.Context, goodsId uint) (out *model.GoodsGetOneOut, err error) {
	out = &model.GoodsGetOneOut{}
	err = dao.Goods.Ctx(ctx).WherePri(goodsId).Scan(&out.GoodsInfo)
	if err != nil {
		return nil, err
	}
	err = s.OutFormat(ctx, out.GoodsInfo)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *sGoods) GetList(ctx context.Context, in model.GoodsGetListInput) (out *model.GoodsGetListOut, err error) {
	out = &model.GoodsGetListOut{
		Page: in.Page,
		Size: in.Size,
	}
	err = dao.Goods.Ctx(ctx).
		Page(in.Page, in.Size).
		OrderDesc(dao.Goods.Columns().Id).
		ScanAndCount(&out.List, &out.Total, true)
	if err != nil {
		return nil, err
	}
	for _, goods := range out.List {
		err = s.OutFormat(ctx, goods)
		if err != nil {
			return nil, err
		}
	}
	//g.Dump(out)
	//g.Log().Info(ctx, out)
	return out, err
}

func (s *sGoods) OutFormat(ctx context.Context, goods *model.GoodsOutItem) error {
	goodsStatusMaps := map[uint]string{consts.GoodsStatusOff: "下架", consts.GoodsStatusOn: "上架"}
	//商品自定义字段
	goods.PriceText = "¥" + fmt.Sprintf("%.2f", goods.Price/100)
	goods.StatusText = goodsStatusMaps[goods.Status]
	return nil
}

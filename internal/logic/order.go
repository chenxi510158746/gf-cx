package logic

import (
	"context"
	"fmt"
	"gf-cx/internal/consts"
	"gf-cx/internal/dao"
	"gf-cx/internal/model"
	"gf-cx/internal/model/do"
	"gf-cx/internal/service"
	"github.com/gogf/gf/v2/util/guid"
	"strings"
)

type sOrder struct{}

func NewOrder() service.IOrder {
	return &sOrder{}
}

func init() {
	service.RegisterOrder(NewOrder())
}

func (s *sOrder) Create(ctx context.Context, goodsId uint) (orderId uint, err error) {
	price, err := dao.Goods.Ctx(ctx).
		Fields(dao.Goods.Columns().Price).
		WherePri(goodsId).
		Value()
	if err != nil {
		return 0, err
	}
	orderSn := "O" + strings.ToUpper(guid.S())
	in := model.OrderCreateItem{
		OrderSn:    orderSn,
		UserId:     service.Session().GetUser(ctx).Id,
		GoodsId:    goodsId,
		GoodsPrice: price.Float32(),
		PayAmount:  price.Float32(),
		PayStatus:  consts.PayStatusPaid,
	}
	id, err := dao.Orders.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return 0, err
	}
	return uint(id), err
}

func (s *sOrder) GetOne(ctx context.Context, orderId uint) (out *model.OrderGetOneOut, err error) {
	out = &model.OrderGetOneOut{}
	userId := service.Session().GetUser(ctx).Id
	err = dao.Orders.Ctx(ctx).
		With(model.GoodsOutItem{}, model.UserOutItem{}).
		Where(do.Orders{
			Id:     orderId,
			UserId: userId,
		}).Scan(&out.Order)
	if err != nil {
		return nil, err
	}
	if out.Order == nil {
		return out, nil
	}
	err = s.OutFormat(ctx, out.Order)
	if err != nil {
		return nil, err
	}
	return out, err
}

func (s *sOrder) GetList(ctx context.Context, in model.OrderGetListInput) (out *model.OrderGetListOut, err error) {
	out = &model.OrderGetListOut{
		Page: in.Page,
		Size: in.Size,
	}
	userId := service.Session().GetUser(ctx).Id
	err = dao.Orders.Ctx(ctx).
		With(model.GoodsOutItem{}, model.UserOutItem{}).
		Where(do.Orders{UserId: userId}).
		Page(in.Page, in.Size).
		OrderDesc(dao.Orders.Columns().Id).
		ScanAndCount(&out.List, &out.Total, true)
	if err != nil {
		return nil, err
	}
	if out.List == nil {
		return out, nil
	}
	for _, order := range out.List {
		err := s.OutFormat(ctx, order)
		if err != nil {
			return nil, err
		}
	}
	return out, err
}

func (s *sOrder) OutFormat(ctx context.Context, order *model.OrderOutItem) error {
	orderStatusMaps := map[uint]string{consts.PayStatusWait: "未支付", consts.PayStatusPaid: "已支付", consts.PayStatusRefund: "退退款"}
	//订单自定义字段
	order.GoodsPriceText = "¥" + fmt.Sprintf("%.2f", order.GoodsPrice/100)
	order.PayAmountText = "¥" + fmt.Sprintf("%.2f", order.PayAmount/100)
	order.PayStatusText = orderStatusMaps[order.PayStatus]
	//商品自定义字段
	err := service.Goods().OutFormat(ctx, order.Goods)
	if err != nil {
		return err
	}
	return nil
}

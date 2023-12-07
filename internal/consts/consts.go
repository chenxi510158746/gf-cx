package consts

const (
	//OpenAPI配置信息
	OpenAPITitle       = "gf-cx项目"
	OpenAPIDescription = "gf-cx项目API，Enjoy 💖 "
	OpenAPIUserVersion = "v1"

	//Session配置信息
	SessionKeyUser = "kii17Yl6TDiCnjxDB1"
	ContextKey     = "k6NSTWtssBxSGG3T"

	//商品配置信息
	GoodsStatusOff = 0 //商品下架
	GoodsStatusOn  = 1 //商品上架

	//订单配置信息
	PayStatusWait   = 1 //待支付
	PayStatusPaid   = 2 //已支付
	PayStatusRefund = 3 //已退款
)

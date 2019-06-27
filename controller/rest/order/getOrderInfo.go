package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【订单表】信息接口

// 请求
type getOrderInfoRequest struct {

	// 订单ID
	orderid    string
	orderidInt int
}

// 方法
func GetOrderInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "order.GetOrderInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getOrderInfoRequest
	//endregion

	//region 验证orderid参数
	request.orderid = c.Param("id")
	if golibs.Length(request.orderid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.GetOrderInfo.orderid is null",
			Message: "缺少【订单ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.orderid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.GetOrderInfo.orderid is number",
			Message: "【订单ID】参数格式不正确",
		})
		return
	}
	request.orderidInt, err = strconv.Atoi(request.orderid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.GetOrderInfo.orderid parse err",
			Message: "orderid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.orderidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.GetOrderInfo.orderid value err",
			Message: "【订单ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【订单表】信息
	orderInfo, err := service.GetOrderInfo(request.orderidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.GetOrderInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if orderInfo.Orderid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.GetOrderInfo.not found",
			Message: "没有找到【订单表】信息",
		})
		return
	}
	if orderInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.GetOrderInfo.has delete",
			Message: "【订单表】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【订单表】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"orderid":      orderInfo.Orderid,                                     //订单ID
			"pkgcode":      orderInfo.Pkgcode,                                     //套餐号
			"ordercode":    orderInfo.Ordercode,                                   //订单号
			"buyer":        orderInfo.Buyer,                                       //用户id
			"paytype":      orderInfo.Paytype,                                     //支付方式
			"totalprice":   orderInfo.Totalprice,                                  //订单总价（元）
			"payprice":     orderInfo.Payprice,                                    //购买总价（元）
			"paytime":      orderInfo.Paytime.Format(golibs.Time_TIMEStandard),    //购买日期
			"isinvalid":    orderInfo.Isinvalid,                                   //是否失效
			"isused":       orderInfo.Isused,                                      //是否使用
			"isrefund":     orderInfo.Isrefund,                                    //是否退款
			"refundprice":  orderInfo.Refundprice,                                 //退款金额
			"refundtime":   orderInfo.Refundtime.Format(golibs.Time_TIMEStandard), //退款日期
			"remark":       orderInfo.Remark,                                      //订单备注
			"adder":        orderInfo.Adder,                                       //创建人
			"addtime":      orderInfo.Addtime.Format(golibs.Time_TIMEStandard),    //创建时间
			"moder":        orderInfo.Moder,                                       //修改人
			"modtime":      orderInfo.Modtime.Format(golibs.Time_TIMEStandard),    //修改时间
			"deleteStatus": orderInfo.DeleteStatus,                                //0:未知，1：未删除，2：已删除
		},
	})
	//endregion
}

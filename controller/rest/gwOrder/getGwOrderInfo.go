package gwOrder

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【订单管理】信息接口

// 请求
type getGwOrderInfoRequest struct {

	// 订单ID
	orderid    string
	orderidInt int
}

// 方法
func GetGwOrderInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwOrder.GetGwOrderInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getGwOrderInfoRequest
	//endregion

	//region 验证orderid参数
	request.orderid = c.Param("id")
	if golibs.Length(request.orderid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.GetGwOrderInfo.orderid is null",
			Message: "缺少【订单ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.orderid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.GetGwOrderInfo.orderid is number",
			Message: "【订单ID】参数格式不正确",
		})
		return
	}
	request.orderidInt, err = strconv.Atoi(request.orderid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.GetGwOrderInfo.orderid parse err",
			Message: "orderid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.orderidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.GetGwOrderInfo.orderid value err",
			Message: "【订单ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【订单管理】信息
	gwOrderInfo, err := service.GetGwOrderInfo(request.orderidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.GetGwOrderInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if gwOrderInfo.Orderid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.GetGwOrderInfo.not found",
			Message: "没有找到【订单管理】信息",
		})
		return
	}
	if gwOrderInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.GetGwOrderInfo.has delete",
			Message: "【订单管理】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【订单管理】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"orderid":           gwOrderInfo.Orderid,                                          //订单ID
			"pkg_code":          gwOrderInfo.PkgCode,                                          //套餐号
			"order_code":        gwOrderInfo.OrderCode,                                        //订单号
			"order_buyer":       gwOrderInfo.OrderBuyer,                                       //用户id
			"order_paytype":     gwOrderInfo.OrderPaytype,                                     //支付方式
			"order_totalprice":  gwOrderInfo.OrderTotalprice,                                  //订单总价（元）
			"order_payprice":    gwOrderInfo.OrderPayprice,                                    //购买总价（元）
			"order_paytime":     gwOrderInfo.OrderPaytime.Format(golibs.Time_TIMEStandard),    //购买日期
			"order_isinvalid":   gwOrderInfo.OrderIsinvalid,                                   //是否失效
			"order_isused":      gwOrderInfo.OrderIsused,                                      //是否使用
			"order_isrefund":    gwOrderInfo.OrderIsrefund,                                    //是否退款
			"order_refundprice": gwOrderInfo.OrderRefundprice,                                 //退款金额
			"order_refundtime":  gwOrderInfo.OrderRefundtime.Format(golibs.Time_TIMEStandard), //退款日期
			"order_remark":      gwOrderInfo.OrderRemark,                                      //订单备注
			"adder":             gwOrderInfo.Adder,                                            //创建人
			"addtime":           gwOrderInfo.Addtime.Format(golibs.Time_TIMEStandard),         //创建时间
			"moder":             gwOrderInfo.Moder,                                            //修改人
			"modtime":           gwOrderInfo.Modtime.Format(golibs.Time_TIMEStandard),         //修改时间
			"deleteStatus":      gwOrderInfo.DeleteStatus,                                     //0:未知，1：未删除，2：已删除
		},
	})
	//endregion
}

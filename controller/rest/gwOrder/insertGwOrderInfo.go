package gwOrder

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"time"
)

// 插入【订单管理】信息接口

// 请求
type insertGwOrderInfoRequest struct {

	// 套餐号
	PkgCode string `form:"pkg_code"`

	// 订单号
	OrderCode string `form:"order_code"`

	// 用户id
	OrderBuyer int `form:"order_buyer"`

	// 支付方式
	OrderPaytype int `form:"order_paytype"`

	// 订单总价（元）
	OrderTotalprice float64 `form:"order_totalprice"`

	// 购买总价（元）
	OrderPayprice float64 `form:"order_payprice"`

	// 购买日期
	OrderPaytime      string `form:"order_paytime"`
	order_paytimeTime time.Time

	// 是否失效
	OrderIsinvalid int `form:"order_isinvalid"`

	// 是否使用
	OrderIsused int `form:"order_isused"`

	// 是否退款
	OrderIsrefund int `form:"order_isrefund"`

	// 退款金额
	OrderRefundprice float64 `form:"order_refundprice"`

	// 退款日期
	OrderRefundtime      string `form:"order_refundtime"`
	order_refundtimeTime time.Time

	// 订单备注
	OrderRemark string `form:"order_remark"`

	// 创建人
	Adder int `form:"adder"`

	// 创建时间
	Addtime     string `form:"addtime"`
	addtimeTime time.Time

	// 修改人
	Moder int `form:"moder"`

	// 修改时间
	Modtime     string `form:"modtime"`
	modtimeTime time.Time
}

// 方法
func InsertGwOrderInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwOrder.InsertGwOrderInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request insertGwOrderInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证pkg_code参数
	if golibs.Length(request.PkgCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.pkg_code is null",
			Message: "缺少【套餐号】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.PkgCode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.pkg_code format err",
			Message: "【套餐号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.PkgCode) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.pkg_code length err",
			Message: "【套餐号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证order_code参数
	if golibs.Length(request.OrderCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_code is null",
			Message: "缺少【订单号】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.OrderCode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_code format err",
			Message: "【订单号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.OrderCode) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_code length err",
			Message: "【订单号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证order_buyer参数
	if request.OrderBuyer <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_buyer value err",
			Message: "order_buyer参数值错误",
		})
		return
	}
	//endregion

	//region 验证order_paytype参数
	if request.OrderPaytype <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_paytype value err",
			Message: "order_paytype参数值错误",
		})
		return
	}
	//endregion

	//region 验证order_paytime参数
	if golibs.Length(request.OrderPaytime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_paytime length err",
			Message: "缺少【购买日期】参数",
		})
		return
	}
	request.order_paytimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.OrderPaytime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_paytime parse err",
			Message: "【购买日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.order_paytimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_paytime value err",
			Message: "【购买日期】参数值错误:" + request.OrderPaytime,
		})
		return
	}
	//endregion

	//region 验证order_isinvalid参数
	if request.OrderIsinvalid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_isinvalid value err",
			Message: "order_isinvalid参数值错误",
		})
		return
	}
	//endregion

	//region 验证order_isused参数
	if request.OrderIsused <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_isused value err",
			Message: "order_isused参数值错误",
		})
		return
	}
	//endregion

	//region 验证order_isrefund参数
	if request.OrderIsrefund <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_isrefund value err",
			Message: "order_isrefund参数值错误",
		})
		return
	}
	//endregion

	//region 验证order_refundtime参数
	if golibs.Length(request.OrderRefundtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_refundtime length err",
			Message: "缺少【退款日期】参数",
		})
		return
	}
	request.order_refundtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.OrderRefundtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_refundtime parse err",
			Message: "【退款日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.order_refundtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_refundtime value err",
			Message: "【退款日期】参数值错误:" + request.OrderRefundtime,
		})
		return
	}
	//endregion

	//region 验证order_remark参数
	if golibs.Length(request.OrderRemark) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_remark is null",
			Message: "缺少【订单备注】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.OrderRemark) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_remark format err",
			Message: "【订单备注】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.OrderRemark) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.order_remark length err",
			Message: "【订单备注】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.adder value err",
			Message: "adder参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.moder value err",
			Message: "moder参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 插入【订单管理】信息
	orderid, err := service.InsertGwOrderInfo(request.PkgCode, request.OrderCode, request.OrderBuyer, request.OrderPaytype, request.OrderTotalprice, request.OrderPayprice, request.order_paytimeTime, int8(request.OrderIsinvalid), int8(request.OrderIsused), int8(request.OrderIsrefund), request.OrderRefundprice, request.order_refundtimeTime, request.OrderRemark, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.insert err",
			Message: "插入出错:" + err.Error(),
		})
		return
	}
	if orderid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwOrder.InsertGwOrderInfo.insert failure",
			Message: "插入失败",
		})
		return
	}
	//endregion

	//region 返回插入【订单管理】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result":  true,
			"orderid": orderid,
		},
	})
	//endregion
}
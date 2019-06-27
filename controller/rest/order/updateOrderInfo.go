package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
	"time"
)

// 修改【订单表】信息接口

// 请求
type updateOrderInfoRequest struct {

	// 订单ID
	orderid    string
	orderidInt int

	// 套餐号
	Pkgcode string `form:"pkgcode"`

	// 订单号
	Ordercode string `form:"ordercode"`

	// 用户id
	Buyer int `form:"buyer"`

	// 支付方式
	Paytype int `form:"paytype"`

	// 订单总价（元）
	Totalprice float64 `form:"totalprice"`

	// 购买总价（元）
	Payprice float64 `form:"payprice"`

	// 购买日期
	Paytime     string `form:"paytime"`
	paytimeTime time.Time

	// 是否失效
	Isinvalid bool `form:"isinvalid"`

	// 是否使用
	Isused bool `form:"isused"`

	// 是否退款
	Isrefund bool `form:"isrefund"`

	// 退款金额
	Refundprice float64 `form:"refundprice"`

	// 退款日期
	Refundtime     string `form:"refundtime"`
	refundtimeTime time.Time

	// 订单备注
	Remark string `form:"remark"`

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
func UpdateOrderInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "order.UpdateOrderInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateOrderInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证orderid参数
	request.orderid = c.Param("id")
	if golibs.Length(request.orderid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.orderid is null",
			Message: "缺少【订单ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.orderid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.orderid is number",
			Message: "【订单ID】参数格式不正确",
		})
		return
	}
	request.orderidInt, err = strconv.Atoi(request.orderid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.orderid parse err",
			Message: "【订单ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.orderidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.orderid value err",
			Message: "【订单ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证pkgcode参数
	if golibs.Length(request.Pkgcode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.pkgcode is null",
			Message: "缺少【套餐号】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.Pkgcode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.pkgcode format err",
			Message: "【套餐号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Pkgcode) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.pkgcode length err",
			Message: "【套餐号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证ordercode参数
	if golibs.Length(request.Ordercode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.ordercode is null",
			Message: "缺少【订单号】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.Ordercode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.ordercode format err",
			Message: "【订单号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Ordercode) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.ordercode length err",
			Message: "【订单号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证buyer参数
	if request.Buyer <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.buyer value err",
			Message: "【用户id】参数值错误",
		})
		return
	}
	//endregion

	//region 验证paytype参数
	if request.Paytype <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.paytype value err",
			Message: "【支付方式】参数值错误",
		})
		return
	}
	//endregion

	//region 验证paytime参数
	if golibs.Length(request.Paytime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.paytime length err",
			Message: "缺少【购买日期】参数",
		})
		return
	}
	request.paytimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Paytime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.paytime parse err",
			Message: "【购买日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.paytimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.paytime value err",
			Message: "【购买日期】参数值错误:" + request.Paytime,
		})
		return
	}
	//endregion

	//region 验证refundtime参数
	if golibs.Length(request.Refundtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.refundtime length err",
			Message: "缺少【退款日期】参数",
		})
		return
	}
	request.refundtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Refundtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.refundtime parse err",
			Message: "【退款日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.refundtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.refundtime value err",
			Message: "【退款日期】参数值错误:" + request.Refundtime,
		})
		return
	}
	//endregion

	//region 验证remark参数
	if golibs.Length(request.Remark) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.remark is null",
			Message: "缺少【订单备注】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Remark) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.remark format err",
			Message: "【订单备注】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Remark) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.remark length err",
			Message: "【订单备注】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.adder value err",
			Message: "【创建人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.moder value err",
			Message: "【修改人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 修改【订单表】信息
	isSuccess, err := service.UpdateOrderInfo(request.orderidInt, request.Pkgcode, request.Ordercode, request.Buyer, request.Paytype, request.Totalprice, request.Payprice, request.paytimeTime, request.Isinvalid, request.Isused, request.Isrefund, request.Refundprice, request.refundtimeTime, request.Remark, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.UpdateOrderInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【订单表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

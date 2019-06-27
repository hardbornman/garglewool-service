package voucher

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"time"
)

// 插入【抵用券管理】信息接口

// 请求
type insertVoucherInfoRequest struct {

	// 抵用券号
	Code string `form:"code"`

	// 用户id
	Guestid int `form:"guestid"`

	// 额度
	Quota float64 `form:"quota"`

	// 抵消
	Reduce float64 `form:"reduce"`

	// 有效期（天）
	Validdays int `form:"validdays"`

	// 是否失效
	Isinvalid bool `form:"isinvalid"`

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
func InsertVoucherInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "voucher.InsertVoucherInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request insertVoucherInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证code参数
	if golibs.Length(request.Code) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.code is null",
			Message: "缺少【抵用券号】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.Code) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.code format err",
			Message: "【抵用券号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Code) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.code length err",
			Message: "【抵用券号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证guestid参数
	if request.Guestid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.guestid value err",
			Message: "guestid参数值错误",
		})
		return
	}
	//endregion

	//region 验证validdays参数
	if request.Validdays <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.validdays value err",
			Message: "validdays参数值错误",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.adder value err",
			Message: "adder参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.moder value err",
			Message: "moder参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 插入【抵用券管理】信息
	voucherid, err := service.InsertVoucherInfo(request.Code, request.Guestid, request.Quota, request.Reduce, request.Validdays, request.Isinvalid, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.insert err",
			Message: "插入出错:" + err.Error(),
		})
		return
	}
	if voucherid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.InsertVoucherInfo.insert failure",
			Message: "插入失败",
		})
		return
	}
	//endregion

	//region 返回插入【抵用券管理】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result":    true,
			"voucherid": voucherid,
		},
	})
	//endregion
}

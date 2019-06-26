package gwVoucher

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
type insertGwVoucherInfoRequest struct {

	// 抵用券号
	VoucherCode string `form:"voucher_code"`

	// 用户id
	VoucherUserid int `form:"voucher_userid"`

	// 额度
	VoucherQuota float64 `form:"voucher_quota"`

	// 抵消
	VoucherReduce float64 `form:"voucher_reduce"`

	// 创建时间
	VoucherCreatetime      string `form:"voucher_createtime"`
	voucher_createtimeTime time.Time

	// 有效期（天）
	VoucherValiddays int `form:"voucher_validdays"`

	// 是否失效
	VoucherIsinvalid int `form:"voucher_isinvalid"`

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
func InsertGwVoucherInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwVoucher.InsertGwVoucherInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request insertGwVoucherInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证voucher_code参数
	if golibs.Length(request.VoucherCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.voucher_code is null",
			Message: "缺少【抵用券号】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.VoucherCode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.voucher_code format err",
			Message: "【抵用券号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.VoucherCode) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.voucher_code length err",
			Message: "【抵用券号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证voucher_userid参数
	if request.VoucherUserid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.voucher_userid value err",
			Message: "voucher_userid参数值错误",
		})
		return
	}
	//endregion

	//region 验证voucher_createtime参数
	if golibs.Length(request.VoucherCreatetime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.voucher_createtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.voucher_createtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.VoucherCreatetime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.voucher_createtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.voucher_createtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.voucher_createtime value err",
			Message: "【创建时间】参数值错误:" + request.VoucherCreatetime,
		})
		return
	}
	//endregion

	//region 验证voucher_validdays参数
	if request.VoucherValiddays <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.voucher_validdays value err",
			Message: "voucher_validdays参数值错误",
		})
		return
	}
	//endregion

	//region 验证voucher_isinvalid参数
	if request.VoucherIsinvalid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.voucher_isinvalid value err",
			Message: "voucher_isinvalid参数值错误",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.adder value err",
			Message: "adder参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.moder value err",
			Message: "moder参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 插入【抵用券管理】信息
	voucherid, err := service.InsertGwVoucherInfo(request.VoucherCode, request.VoucherUserid, request.VoucherQuota, request.VoucherReduce, request.voucher_createtimeTime, request.VoucherValiddays, int8(request.VoucherIsinvalid), request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.insert err",
			Message: "插入出错:" + err.Error(),
		})
		return
	}
	if voucherid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.InsertGwVoucherInfo.insert failure",
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

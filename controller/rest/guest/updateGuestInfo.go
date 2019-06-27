package guest

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

// 修改【买家客户表】信息接口

// 请求
type updateGuestInfoRequest struct {

	// 用户ID
	guestid    string
	guestidInt int

	// 用户名
	Name string `form:"name"`

	// 用户密码
	Password string `form:"password"`

	// 手机号
	Phone string `form:"phone"`

	// 用户金币数量
	Golds int `form:"golds"`

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
func UpdateGuestInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "guest.UpdateGuestInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateGuestInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证guestid参数
	request.guestid = c.Param("id")
	if golibs.Length(request.guestid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.guestid is null",
			Message: "缺少【用户ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.guestid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.guestid is number",
			Message: "【用户ID】参数格式不正确",
		})
		return
	}
	request.guestidInt, err = strconv.Atoi(request.guestid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.guestid parse err",
			Message: "【用户ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.guestidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.guestid value err",
			Message: "【用户ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证name参数
	if golibs.Length(request.Name) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.name is null",
			Message: "缺少【用户名】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Name) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.name format err",
			Message: "【用户名】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Name) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.name length err",
			Message: "【用户名】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证password参数
	if golibs.Length(request.Password) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.password is null",
			Message: "缺少【用户密码】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Password) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.password format err",
			Message: "【用户密码】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Password) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.password length err",
			Message: "【用户密码】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证phone参数
	if golibs.Length(request.Phone) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.phone is null",
			Message: "缺少【手机号】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Phone) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.phone format err",
			Message: "【手机号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Phone) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.phone length err",
			Message: "【手机号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证golds参数
	if request.Golds <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.golds value err",
			Message: "【用户金币数量】参数值错误",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.adder value err",
			Message: "【创建人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.moder value err",
			Message: "【修改人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 修改【买家客户表】信息
	isSuccess, err := service.UpdateGuestInfo(request.guestidInt, request.Name, request.Password, request.Phone, request.Golds, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.UpdateGuestInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【买家客户表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

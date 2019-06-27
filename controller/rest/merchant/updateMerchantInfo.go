package merchant

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

// 修改【商家用户表】信息接口

// 请求
type updateMerchantInfoRequest struct {

	// 商家ID
	merchantid    string
	merchantidInt int

	// 商家名
	Merchantname string `form:"merchantname"`

	// 手机号
	Phone string `form:"phone"`

	// 用户角色ID
	Userroleid int `form:"userroleid"`

	// 登录密码
	Loginpwd string `form:"loginpwd"`

	// 登录账户
	Loginaccount string `form:"loginaccount"`

	// 昵称
	Nick string `form:"nick"`

	// 微信账号
	Wechataccount string `form:"wechataccount"`

	// 微信签名
	Wechatsign string `form:"wechatsign"`

	// 备注
	Remark string `form:"remark"`

	// 最近登录时间
	Lastlogintime     string `form:"lastlogintime"`
	lastlogintimeTime time.Time

	// 最近登录地址
	Lastloginaddr string `form:"lastloginaddr"`

	// 注册时间
	Rigstertime     string `form:"rigstertime"`
	rigstertimeTime time.Time

	// 是否启用
	Enable bool `form:"enable"`

	// 地址
	Addr string `form:"addr"`
}

// 方法
func UpdateMerchantInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "merchant.UpdateMerchantInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateMerchantInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证merchantid参数
	request.merchantid = c.Param("id")
	if golibs.Length(request.merchantid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.merchantid is null",
			Message: "缺少【商家ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.merchantid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.merchantid is number",
			Message: "【商家ID】参数格式不正确",
		})
		return
	}
	request.merchantidInt, err = strconv.Atoi(request.merchantid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.merchantid parse err",
			Message: "【商家ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.merchantidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.merchantid value err",
			Message: "【商家ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证merchantname参数
	if golibs.Length(request.Merchantname) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.merchantname is null",
			Message: "缺少【商家名】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Merchantname) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.merchantname format err",
			Message: "【商家名】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Merchantname) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.merchantname length err",
			Message: "【商家名】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证phone参数
	if golibs.Length(request.Phone) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.phone is null",
			Message: "缺少【手机号】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Phone) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.phone format err",
			Message: "【手机号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Phone) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.phone length err",
			Message: "【手机号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证userroleid参数
	if request.Userroleid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.userroleid value err",
			Message: "【用户角色ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证loginpwd参数
	if golibs.Length(request.Loginpwd) <= 0 {
		request.Loginpwd = "kong_zi_fu_chuan"
	}
	if !golibs.IsUtf8(request.Loginpwd) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.loginpwd format err",
			Message: "【登录密码】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Loginpwd) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.loginpwd length err",
			Message: "【登录密码】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证loginaccount参数
	if golibs.Length(request.Loginaccount) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.loginaccount is null",
			Message: "缺少【登录账户】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Loginaccount) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.loginaccount format err",
			Message: "【登录账户】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Loginaccount) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.loginaccount length err",
			Message: "【登录账户】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证nick参数
	if golibs.Length(request.Nick) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.nick is null",
			Message: "缺少【昵称】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Nick) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.nick format err",
			Message: "【昵称】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Nick) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.nick length err",
			Message: "【昵称】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证wechataccount参数
	if golibs.Length(request.Wechataccount) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.wechataccount is null",
			Message: "缺少【微信账号】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Wechataccount) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.wechataccount format err",
			Message: "【微信账号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Wechataccount) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.wechataccount length err",
			Message: "【微信账号】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证wechatsign参数
	if golibs.Length(request.Wechatsign) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.wechatsign is null",
			Message: "缺少【微信签名】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Wechatsign) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.wechatsign format err",
			Message: "【微信签名】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Wechatsign) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.wechatsign length err",
			Message: "【微信签名】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证remark参数
	if golibs.Length(request.Remark) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.remark is null",
			Message: "缺少【备注】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Remark) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.remark format err",
			Message: "【备注】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Remark) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.remark length err",
			Message: "【备注】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证lastlogintime参数
	if golibs.Length(request.Lastlogintime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.lastlogintime length err",
			Message: "缺少【最近登录时间】参数",
		})
		return
	}
	request.lastlogintimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Lastlogintime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.lastlogintime parse err",
			Message: "【最近登录时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.lastlogintimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.lastlogintime value err",
			Message: "【最近登录时间】参数值错误:" + request.Lastlogintime,
		})
		return
	}
	//endregion

	//region 验证lastloginaddr参数
	if golibs.Length(request.Lastloginaddr) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.lastloginaddr is null",
			Message: "缺少【最近登录地址】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Lastloginaddr) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.lastloginaddr format err",
			Message: "【最近登录地址】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Lastloginaddr) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.lastloginaddr length err",
			Message: "【最近登录地址】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证rigstertime参数
	if golibs.Length(request.Rigstertime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.rigstertime length err",
			Message: "缺少【注册时间】参数",
		})
		return
	}
	request.rigstertimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Rigstertime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.rigstertime parse err",
			Message: "【注册时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.rigstertimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.rigstertime value err",
			Message: "【注册时间】参数值错误:" + request.Rigstertime,
		})
		return
	}
	//endregion

	//region 验证addr参数
	if golibs.Length(request.Addr) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.addr is null",
			Message: "缺少【地址】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Addr) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.addr format err",
			Message: "【地址】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Addr) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.addr length err",
			Message: "【地址】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 修改【商家用户表】信息
	isSuccess, err := service.UpdateMerchantInfo(request.merchantidInt, request.Merchantname, request.Phone, request.Userroleid, request.Loginpwd, request.Loginaccount, request.Nick, request.Wechataccount, request.Wechatsign, request.Remark, request.lastlogintimeTime, request.Lastloginaddr, request.rigstertimeTime, request.Enable, request.Addr)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.UpdateMerchantInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【商家用户表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

package merchant

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"time"
)

// 插入【商家用户表】信息接口

// 请求
type insertMerchantInfoRequest struct {

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
func InsertMerchantInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "merchant.InsertMerchantInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request insertMerchantInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证merchantname参数
	if golibs.Length(request.Merchantname) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.merchantname is null",
			Message: "缺少【商家名】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Merchantname) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.merchantname format err",
			Message: "【商家名】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Merchantname) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.merchantname length err",
			Message: "【商家名】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证phone参数
	if golibs.Length(request.Phone) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.phone is null",
			Message: "缺少【手机号】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Phone) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.phone format err",
			Message: "【手机号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Phone) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.phone length err",
			Message: "【手机号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证userroleid参数
	if request.Userroleid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.userroleid value err",
			Message: "userroleid参数值错误",
		})
		return
	}
	//endregion

	//region 验证loginpwd参数
	if golibs.Length(request.Loginpwd) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.loginpwd is null",
			Message: "缺少【登录密码】参数",
		})
		return
	}
	if !golibs.IsUtf8(request.Loginpwd) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.loginpwd format err",
			Message: "【登录密码】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Loginpwd) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.loginpwd length err",
			Message: "【登录密码】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证loginaccount参数
	if golibs.Length(request.Loginaccount) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.loginaccount is null",
			Message: "缺少【登录账户】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Loginaccount) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.loginaccount format err",
			Message: "【登录账户】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Loginaccount) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.loginaccount length err",
			Message: "【登录账户】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证nick参数
	if golibs.Length(request.Nick) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.nick is null",
			Message: "缺少【昵称】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Nick) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.nick format err",
			Message: "【昵称】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Nick) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.nick length err",
			Message: "【昵称】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证wechataccount参数
	if golibs.Length(request.Wechataccount) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.wechataccount is null",
			Message: "缺少【微信账号】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Wechataccount) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.wechataccount format err",
			Message: "【微信账号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Wechataccount) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.wechataccount length err",
			Message: "【微信账号】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证wechatsign参数
	if golibs.Length(request.Wechatsign) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.wechatsign is null",
			Message: "缺少【微信签名】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Wechatsign) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.wechatsign format err",
			Message: "【微信签名】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Wechatsign) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.wechatsign length err",
			Message: "【微信签名】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证remark参数
	if golibs.Length(request.Remark) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.remark is null",
			Message: "缺少【备注】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Remark) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.remark format err",
			Message: "【备注】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Remark) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.remark length err",
			Message: "【备注】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证lastlogintime参数
	if golibs.Length(request.Lastlogintime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.lastlogintime length err",
			Message: "缺少【最近登录时间】参数",
		})
		return
	}
	request.lastlogintimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Lastlogintime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.lastlogintime parse err",
			Message: "【最近登录时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.lastlogintimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.lastlogintime value err",
			Message: "【最近登录时间】参数值错误:" + request.Lastlogintime,
		})
		return
	}
	//endregion

	//region 验证lastloginaddr参数
	if golibs.Length(request.Lastloginaddr) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.lastloginaddr is null",
			Message: "缺少【最近登录地址】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Lastloginaddr) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.lastloginaddr format err",
			Message: "【最近登录地址】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Lastloginaddr) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.lastloginaddr length err",
			Message: "【最近登录地址】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证rigstertime参数
	if golibs.Length(request.Rigstertime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.rigstertime length err",
			Message: "缺少【注册时间】参数",
		})
		return
	}
	request.rigstertimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Rigstertime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.rigstertime parse err",
			Message: "【注册时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.rigstertimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.rigstertime value err",
			Message: "【注册时间】参数值错误:" + request.Rigstertime,
		})
		return
	}
	//endregion

	//region 验证addr参数
	if golibs.Length(request.Addr) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.addr is null",
			Message: "缺少【地址】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Addr) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.addr format err",
			Message: "【地址】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Addr) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.addr length err",
			Message: "【地址】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 插入【商家用户表】信息
	merchantid, err := service.InsertMerchantInfo(request.Merchantname, request.Phone, request.Userroleid, request.Loginpwd, request.Loginaccount, request.Nick, request.Wechataccount, request.Wechatsign, request.Remark, request.lastlogintimeTime, request.Lastloginaddr, request.rigstertimeTime, request.Enable, request.Addr)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.insert err",
			Message: "插入出错:" + err.Error(),
		})
		return
	}
	if merchantid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.InsertMerchantInfo.insert failure",
			Message: "插入失败",
		})
		return
	}
	//endregion

	//region 返回插入【商家用户表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result":     true,
			"merchantid": merchantid,
		},
	})
	//endregion
}

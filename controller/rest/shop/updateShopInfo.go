package shop

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

// 修改【店铺表】信息接口

// 请求
type updateShopInfoRequest struct {

	// 店铺ID
	shopid    string
	shopidInt int

	// 店铺号
	Shopcode string `form:"shopcode"`

	// 店铺名称
	Shopname string `form:"shopname"`

	// 店铺地址-省
	Province string `form:"province"`

	// 店铺地址-市
	City string `form:"city"`

	// 店铺地址-区
	District string `form:"district"`

	// 店铺详细地址
	Address string `form:"address"`

	// 电话
	Phone string `form:"phone"`

	// 加盟平台日期
	Leaguetime     string `form:"leaguetime"`
	leaguetimeTime time.Time

	// 退出平台日期
	Exittime     string `form:"exittime"`
	exittimeTime time.Time

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

	// 商家ID
	Merchantid int `form:"merchantid"`

	// 经度
	Longtitude string `form:"longtitude"`

	// 纬度
	Latitude string `form:"latitude"`
}

// 方法
func UpdateShopInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "shop.UpdateShopInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateShopInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证shopid参数
	request.shopid = c.Param("id")
	if golibs.Length(request.shopid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.shopid is null",
			Message: "缺少【店铺ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.shopid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.shopid is number",
			Message: "【店铺ID】参数格式不正确",
		})
		return
	}
	request.shopidInt, err = strconv.Atoi(request.shopid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.shopid parse err",
			Message: "【店铺ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.shopidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.shopid value err",
			Message: "【店铺ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证shopcode参数
	if golibs.Length(request.Shopcode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.shopcode is null",
			Message: "缺少【店铺号】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.Shopcode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.shopcode format err",
			Message: "【店铺号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Shopcode) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.shopcode length err",
			Message: "【店铺号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证shopname参数
	if golibs.Length(request.Shopname) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.shopname is null",
			Message: "缺少【店铺名称】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Shopname) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.shopname format err",
			Message: "【店铺名称】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Shopname) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.shopname length err",
			Message: "【店铺名称】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证province参数
	if golibs.Length(request.Province) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.province is null",
			Message: "缺少【店铺地址-省】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Province) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.province format err",
			Message: "【店铺地址-省】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Province) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.province length err",
			Message: "【店铺地址-省】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证city参数
	if golibs.Length(request.City) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.city is null",
			Message: "缺少【店铺地址-市】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.City) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.city format err",
			Message: "【店铺地址-市】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.City) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.city length err",
			Message: "【店铺地址-市】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证district参数
	if golibs.Length(request.District) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.district is null",
			Message: "缺少【店铺地址-区】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.District) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.district format err",
			Message: "【店铺地址-区】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.District) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.district length err",
			Message: "【店铺地址-区】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证address参数
	if golibs.Length(request.Address) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.address is null",
			Message: "缺少【店铺详细地址】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Address) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.address format err",
			Message: "【店铺详细地址】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Address) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.address length err",
			Message: "【店铺详细地址】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证phone参数
	if golibs.Length(request.Phone) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.phone is null",
			Message: "缺少【电话】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Phone) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.phone format err",
			Message: "【电话】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Phone) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.phone length err",
			Message: "【电话】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证leaguetime参数
	if golibs.Length(request.Leaguetime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.leaguetime length err",
			Message: "缺少【加盟平台日期】参数",
		})
		return
	}
	request.leaguetimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Leaguetime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.leaguetime parse err",
			Message: "【加盟平台日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.leaguetimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.leaguetime value err",
			Message: "【加盟平台日期】参数值错误:" + request.Leaguetime,
		})
		return
	}
	//endregion

	//region 验证exittime参数
	if golibs.Length(request.Exittime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.exittime length err",
			Message: "缺少【退出平台日期】参数",
		})
		return
	}
	request.exittimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Exittime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.exittime parse err",
			Message: "【退出平台日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.exittimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.exittime value err",
			Message: "【退出平台日期】参数值错误:" + request.Exittime,
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.adder value err",
			Message: "【创建人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.moder value err",
			Message: "【修改人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 验证merchantid参数
	if request.Merchantid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.merchantid value err",
			Message: "【商家ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证longtitude参数
	if golibs.Length(request.Longtitude) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.longtitude is null",
			Message: "缺少【经度】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Longtitude) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.longtitude format err",
			Message: "【经度】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Longtitude) > 20 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.longtitude length err",
			Message: "【经度】参数长度不能超过20个字符",
		})
		return
	}
	//endregion

	//region 验证latitude参数
	if golibs.Length(request.Latitude) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.latitude is null",
			Message: "缺少【纬度】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Latitude) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.latitude format err",
			Message: "【纬度】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Latitude) > 20 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.latitude length err",
			Message: "【纬度】参数长度不能超过20个字符",
		})
		return
	}
	//endregion

	//region 修改【店铺表】信息
	isSuccess, err := service.UpdateShopInfo(request.shopidInt, request.Shopcode, request.Shopname, request.Province, request.City, request.District, request.Address, request.Phone, request.leaguetimeTime, request.exittimeTime, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime, request.Merchantid, request.Longtitude, request.Latitude)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.UpdateShopInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【店铺表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

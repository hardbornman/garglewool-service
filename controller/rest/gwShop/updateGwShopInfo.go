package gwShop

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

// 修改【店铺管理】信息接口

// 请求
type updateGwShopInfoRequest struct {

	// 店铺ID
	shopid    string
	shopidInt int

	// 店铺号
	ShopCode string `form:"shop_code"`

	// 店铺名称
	ShopName string `form:"shop_name"`

	// 店铺地址-省
	ShopProvince string `form:"shop_province"`

	// 店铺地址-市
	ShopCity string `form:"shop_city"`

	// 店铺地址-区
	ShopDistrict string `form:"shop_district"`

	// 店铺详细地址
	ShopAddress string `form:"shop_address"`

	// 加盟平台日期
	ShopPhone string `form:"shop_phone"`

	// 加盟平台日期
	ShopCreatetime      string `form:"shop_createtime"`
	shop_createtimeTime time.Time

	// 退出平台日期
	ShopExittime      string `form:"shop_exittime"`
	shop_exittimeTime time.Time

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
func UpdateGwShopInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwShop.UpdateGwShopInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateGwShopInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证shopid参数
	request.shopid = c.Param("id")
	if golibs.Length(request.shopid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shopid is null",
			Message: "缺少【店铺ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.shopid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shopid is number",
			Message: "【店铺ID】参数格式不正确",
		})
		return
	}
	request.shopidInt, err = strconv.Atoi(request.shopid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shopid parse err",
			Message: "【店铺ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.shopidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shopid value err",
			Message: "【店铺ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证shop_code参数
	if golibs.Length(request.ShopCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_code is null",
			Message: "缺少【店铺号】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.ShopCode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_code format err",
			Message: "【店铺号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.ShopCode) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_code length err",
			Message: "【店铺号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证shop_name参数
	if golibs.Length(request.ShopName) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_name is null",
			Message: "缺少【店铺名称】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.ShopName) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_name format err",
			Message: "【店铺名称】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.ShopName) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_name length err",
			Message: "【店铺名称】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证shop_province参数
	if golibs.Length(request.ShopProvince) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_province is null",
			Message: "缺少【店铺地址-省】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.ShopProvince) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_province format err",
			Message: "【店铺地址-省】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.ShopProvince) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_province length err",
			Message: "【店铺地址-省】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证shop_city参数
	if golibs.Length(request.ShopCity) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_city is null",
			Message: "缺少【店铺地址-市】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.ShopCity) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_city format err",
			Message: "【店铺地址-市】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.ShopCity) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_city length err",
			Message: "【店铺地址-市】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证shop_district参数
	if golibs.Length(request.ShopDistrict) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_district is null",
			Message: "缺少【店铺地址-区】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.ShopDistrict) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_district format err",
			Message: "【店铺地址-区】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.ShopDistrict) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_district length err",
			Message: "【店铺地址-区】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证shop_address参数
	if golibs.Length(request.ShopAddress) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_address is null",
			Message: "缺少【店铺详细地址】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.ShopAddress) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_address format err",
			Message: "【店铺详细地址】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.ShopAddress) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_address length err",
			Message: "【店铺详细地址】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证shop_phone参数
	if golibs.Length(request.ShopPhone) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_phone is null",
			Message: "缺少【加盟平台日期】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.ShopPhone) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_phone format err",
			Message: "【加盟平台日期】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.ShopPhone) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_phone length err",
			Message: "【加盟平台日期】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证shop_createtime参数
	if golibs.Length(request.ShopCreatetime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_createtime length err",
			Message: "缺少【加盟平台日期】参数",
		})
		return
	}
	request.shop_createtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.ShopCreatetime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_createtime parse err",
			Message: "【加盟平台日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.shop_createtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_createtime value err",
			Message: "【加盟平台日期】参数值错误:" + request.ShopCreatetime,
		})
		return
	}
	//endregion

	//region 验证shop_exittime参数
	if golibs.Length(request.ShopExittime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_exittime length err",
			Message: "缺少【退出平台日期】参数",
		})
		return
	}
	request.shop_exittimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.ShopExittime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_exittime parse err",
			Message: "【退出平台日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.shop_exittimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.shop_exittime value err",
			Message: "【退出平台日期】参数值错误:" + request.ShopExittime,
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.adder value err",
			Message: "【创建人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.moder value err",
			Message: "【修改人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 修改【店铺管理】信息
	isSuccess, err := service.UpdateGwShopInfo(request.shopidInt, request.ShopCode, request.ShopName, request.ShopProvince, request.ShopCity, request.ShopDistrict, request.ShopAddress, request.ShopPhone, request.shop_createtimeTime, request.shop_exittimeTime, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.UpdateGwShopInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【店铺管理】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

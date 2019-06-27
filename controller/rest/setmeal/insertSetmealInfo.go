package setmeal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"time"
)

// 插入【套餐表】信息接口

// 请求
type insertSetmealInfoRequest struct {

	// 店铺ID
	Shopid int `form:"shopid"`

	// 套餐号
	Pkgcode string `form:"pkgcode"`

	// 套餐类别
	Setmealtype string `form:"setmealtype"`

	// 套餐标题
	Title string `form:"title"`

	// 人数
	People int `form:"people"`

	// 是否需要预约
	Isorder bool `form:"isorder"`

	// 是否支持退款
	Isrefund bool `form:"isrefund"`

	// 是否仅限堂食
	Isinhouse bool `form:"isinhouse"`

	// 是否新品
	Isnew bool `form:"isnew"`

	// 是否强烈推荐
	Isrecommend bool `form:"isrecommend"`

	// 有效期（天）
	Validdays int `form:"validdays"`

	// 今日关注人数
	Watchers int `form:"watchers"`

	// 生效日期
	Validtime     string `form:"validtime"`
	validtimeTime time.Time

	// 下架日期
	Exittime     string `form:"exittime"`
	exittimeTime time.Time

	// 更多优惠链接地址
	Links string `form:"links"`

	// 套餐说明
	Info string `form:"info"`

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
func InsertSetmealInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "setmeal.InsertSetmealInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request insertSetmealInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证shopid参数
	if request.Shopid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.shopid value err",
			Message: "shopid参数值错误",
		})
		return
	}
	//endregion

	//region 验证pkgcode参数
	if golibs.Length(request.Pkgcode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.pkgcode is null",
			Message: "缺少【套餐号】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.Pkgcode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.pkgcode format err",
			Message: "【套餐号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Pkgcode) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.pkgcode length err",
			Message: "【套餐号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证setmealtype参数
	if golibs.Length(request.Setmealtype) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.setmealtype is null",
			Message: "缺少【套餐类别】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Setmealtype) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.setmealtype format err",
			Message: "【套餐类别】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Setmealtype) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.setmealtype length err",
			Message: "【套餐类别】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证title参数
	if golibs.Length(request.Title) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.title is null",
			Message: "缺少【套餐标题】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Title) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.title format err",
			Message: "【套餐标题】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Title) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.title length err",
			Message: "【套餐标题】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证people参数
	if request.People <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.people value err",
			Message: "people参数值错误",
		})
		return
	}
	//endregion

	//region 验证validdays参数
	if request.Validdays <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.validdays value err",
			Message: "validdays参数值错误",
		})
		return
	}
	//endregion

	//region 验证watchers参数
	if request.Watchers <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.watchers value err",
			Message: "watchers参数值错误",
		})
		return
	}
	//endregion

	//region 验证validtime参数
	if golibs.Length(request.Validtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.validtime length err",
			Message: "缺少【生效日期】参数",
		})
		return
	}
	request.validtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Validtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.validtime parse err",
			Message: "【生效日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.validtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.validtime value err",
			Message: "【生效日期】参数值错误:" + request.Validtime,
		})
		return
	}
	//endregion

	//region 验证exittime参数
	if golibs.Length(request.Exittime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.exittime length err",
			Message: "缺少【下架日期】参数",
		})
		return
	}
	request.exittimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Exittime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.exittime parse err",
			Message: "【下架日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.exittimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.exittime value err",
			Message: "【下架日期】参数值错误:" + request.Exittime,
		})
		return
	}
	//endregion

	//region 验证links参数
	if golibs.Length(request.Links) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.links is null",
			Message: "缺少【更多优惠链接地址】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Links) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.links format err",
			Message: "【更多优惠链接地址】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Links) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.links length err",
			Message: "【更多优惠链接地址】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证info参数
	if golibs.Length(request.Info) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.info is null",
			Message: "缺少【套餐说明】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Info) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.info format err",
			Message: "【套餐说明】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Info) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.info length err",
			Message: "【套餐说明】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.adder value err",
			Message: "adder参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.moder value err",
			Message: "moder参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 插入【套餐表】信息
	setmealid, err := service.InsertSetmealInfo(request.Shopid, request.Pkgcode, request.Setmealtype, request.Title, request.People, request.Isorder, request.Isrefund, request.Isinhouse, request.Isnew, request.Isrecommend, request.Validdays, request.Watchers, request.validtimeTime, request.exittimeTime, request.Links, request.Info, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.insert err",
			Message: "插入出错:" + err.Error(),
		})
		return
	}
	if setmealid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.InsertSetmealInfo.insert failure",
			Message: "插入失败",
		})
		return
	}
	//endregion

	//region 返回插入【套餐表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result":    true,
			"setmealid": setmealid,
		},
	})
	//endregion
}

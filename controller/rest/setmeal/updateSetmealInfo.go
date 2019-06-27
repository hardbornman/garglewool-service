package setmeal

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

// 修改【套餐表】信息接口

// 请求
type updateSetmealInfoRequest struct {

	// 套餐ID
	setmealid    string
	setmealidInt int

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
func UpdateSetmealInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "setmeal.UpdateSetmealInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateSetmealInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证setmealid参数
	request.setmealid = c.Param("id")
	if golibs.Length(request.setmealid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.setmealid is null",
			Message: "缺少【套餐ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.setmealid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.setmealid is number",
			Message: "【套餐ID】参数格式不正确",
		})
		return
	}
	request.setmealidInt, err = strconv.Atoi(request.setmealid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.setmealid parse err",
			Message: "【套餐ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.setmealidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.setmealid value err",
			Message: "【套餐ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证shopid参数
	if request.Shopid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.shopid value err",
			Message: "【店铺ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证pkgcode参数
	if golibs.Length(request.Pkgcode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.pkgcode is null",
			Message: "缺少【套餐号】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.Pkgcode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.pkgcode format err",
			Message: "【套餐号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Pkgcode) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.pkgcode length err",
			Message: "【套餐号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证setmealtype参数
	if golibs.Length(request.Setmealtype) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.setmealtype is null",
			Message: "缺少【套餐类别】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Setmealtype) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.setmealtype format err",
			Message: "【套餐类别】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Setmealtype) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.setmealtype length err",
			Message: "【套餐类别】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证title参数
	if golibs.Length(request.Title) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.title is null",
			Message: "缺少【套餐标题】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Title) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.title format err",
			Message: "【套餐标题】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Title) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.title length err",
			Message: "【套餐标题】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证people参数
	if request.People <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.people value err",
			Message: "【人数】参数值错误",
		})
		return
	}
	//endregion

	//region 验证validdays参数
	if request.Validdays <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.validdays value err",
			Message: "【有效期（天）】参数值错误",
		})
		return
	}
	//endregion

	//region 验证watchers参数
	if request.Watchers <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.watchers value err",
			Message: "【今日关注人数】参数值错误",
		})
		return
	}
	//endregion

	//region 验证validtime参数
	if golibs.Length(request.Validtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.validtime length err",
			Message: "缺少【生效日期】参数",
		})
		return
	}
	request.validtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Validtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.validtime parse err",
			Message: "【生效日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.validtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.validtime value err",
			Message: "【生效日期】参数值错误:" + request.Validtime,
		})
		return
	}
	//endregion

	//region 验证exittime参数
	if golibs.Length(request.Exittime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.exittime length err",
			Message: "缺少【下架日期】参数",
		})
		return
	}
	request.exittimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Exittime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.exittime parse err",
			Message: "【下架日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.exittimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.exittime value err",
			Message: "【下架日期】参数值错误:" + request.Exittime,
		})
		return
	}
	//endregion

	//region 验证links参数
	if golibs.Length(request.Links) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.links is null",
			Message: "缺少【更多优惠链接地址】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Links) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.links format err",
			Message: "【更多优惠链接地址】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Links) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.links length err",
			Message: "【更多优惠链接地址】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证info参数
	if golibs.Length(request.Info) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.info is null",
			Message: "缺少【套餐说明】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Info) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.info format err",
			Message: "【套餐说明】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Info) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.info length err",
			Message: "【套餐说明】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.adder value err",
			Message: "【创建人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.moder value err",
			Message: "【修改人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 修改【套餐表】信息
	isSuccess, err := service.UpdateSetmealInfo(request.setmealidInt, request.Shopid, request.Pkgcode, request.Setmealtype, request.Title, request.People, request.Isorder, request.Isrefund, request.Isinhouse, request.Isnew, request.Isrecommend, request.Validdays, request.Watchers, request.validtimeTime, request.exittimeTime, request.Links, request.Info, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.UpdateSetmealInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【套餐表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

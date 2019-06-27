package dictionarycategory

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

// 修改【字典分类】信息接口

// 请求
type updateDictionarycategoryInfoRequest struct {

	// 字典分类ID
	dictionarycategoryid    string
	dictionarycategoryidInt int

	// 字典类别key
	Categorykey string `form:"categorykey"`

	// 字典类别value
	Categoryvalue string `form:"categoryvalue"`

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
func UpdateDictionarycategoryInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "dictionarycategory.UpdateDictionarycategoryInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateDictionarycategoryInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证dictionarycategoryid参数
	request.dictionarycategoryid = c.Param("id")
	if golibs.Length(request.dictionarycategoryid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.dictionarycategoryid is null",
			Message: "缺少【字典分类ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.dictionarycategoryid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.dictionarycategoryid is number",
			Message: "【字典分类ID】参数格式不正确",
		})
		return
	}
	request.dictionarycategoryidInt, err = strconv.Atoi(request.dictionarycategoryid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.dictionarycategoryid parse err",
			Message: "【字典分类ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.dictionarycategoryidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.dictionarycategoryid value err",
			Message: "【字典分类ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证categorykey参数
	if golibs.Length(request.Categorykey) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.categorykey is null",
			Message: "缺少【字典类别key】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Categorykey) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.categorykey format err",
			Message: "【字典类别key】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Categorykey) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.categorykey length err",
			Message: "【字典类别key】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证categoryvalue参数
	if golibs.Length(request.Categoryvalue) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.categoryvalue is null",
			Message: "缺少【字典类别value】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Categoryvalue) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.categoryvalue format err",
			Message: "【字典类别value】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Categoryvalue) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.categoryvalue length err",
			Message: "【字典类别value】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.adder value err",
			Message: "【创建人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.moder value err",
			Message: "【修改人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 修改【字典分类】信息
	isSuccess, err := service.UpdateDictionarycategoryInfo(request.dictionarycategoryidInt, request.Categorykey, request.Categoryvalue, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.UpdateDictionarycategoryInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【字典分类】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

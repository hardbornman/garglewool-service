package dictionarycategory

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"time"
)

// 插入【字典分类】信息接口

// 请求
type insertDictionarycategoryInfoRequest struct {

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
func InsertDictionarycategoryInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "dictionarycategory.InsertDictionarycategoryInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request insertDictionarycategoryInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证categorykey参数
	if golibs.Length(request.Categorykey) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.categorykey is null",
			Message: "缺少【字典类别key】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Categorykey) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.categorykey format err",
			Message: "【字典类别key】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Categorykey) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.categorykey length err",
			Message: "【字典类别key】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证categoryvalue参数
	if golibs.Length(request.Categoryvalue) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.categoryvalue is null",
			Message: "缺少【字典类别value】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Categoryvalue) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.categoryvalue format err",
			Message: "【字典类别value】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Categoryvalue) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.categoryvalue length err",
			Message: "【字典类别value】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.adder value err",
			Message: "adder参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.moder value err",
			Message: "moder参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 插入【字典分类】信息
	dictionarycategoryid, err := service.InsertDictionarycategoryInfo(request.Categorykey, request.Categoryvalue, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.insert err",
			Message: "插入出错:" + err.Error(),
		})
		return
	}
	if dictionarycategoryid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.InsertDictionarycategoryInfo.insert failure",
			Message: "插入失败",
		})
		return
	}
	//endregion

	//region 返回插入【字典分类】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result":               true,
			"dictionarycategoryid": dictionarycategoryid,
		},
	})
	//endregion
}

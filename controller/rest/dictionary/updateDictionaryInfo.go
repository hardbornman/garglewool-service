package dictionary

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

// 修改【字典表】信息接口

// 请求
type updateDictionaryInfoRequest struct {

	// 字典ID
	dictionaryid    string
	dictionaryidInt int

	// 字典类别key
	Categorykey string `form:"categorykey"`

	// key
	Dictkey string `form:"dictkey"`

	// value
	Dictvalue string `form:"dictvalue"`

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
func UpdateDictionaryInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "dictionary.UpdateDictionaryInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateDictionaryInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证dictionaryid参数
	request.dictionaryid = c.Param("id")
	if golibs.Length(request.dictionaryid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.dictionaryid is null",
			Message: "缺少【字典ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.dictionaryid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.dictionaryid is number",
			Message: "【字典ID】参数格式不正确",
		})
		return
	}
	request.dictionaryidInt, err = strconv.Atoi(request.dictionaryid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.dictionaryid parse err",
			Message: "【字典ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.dictionaryidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.dictionaryid value err",
			Message: "【字典ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证categorykey参数
	if golibs.Length(request.Categorykey) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.categorykey is null",
			Message: "缺少【字典类别key】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Categorykey) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.categorykey format err",
			Message: "【字典类别key】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Categorykey) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.categorykey length err",
			Message: "【字典类别key】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证dictkey参数
	if golibs.Length(request.Dictkey) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.dictkey is null",
			Message: "缺少【key】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Dictkey) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.dictkey format err",
			Message: "【key】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Dictkey) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.dictkey length err",
			Message: "【key】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证dictvalue参数
	if golibs.Length(request.Dictvalue) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.dictvalue is null",
			Message: "缺少【value】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Dictvalue) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.dictvalue format err",
			Message: "【value】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Dictvalue) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.dictvalue length err",
			Message: "【value】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.adder value err",
			Message: "【创建人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.moder value err",
			Message: "【修改人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 修改【字典表】信息
	isSuccess, err := service.UpdateDictionaryInfo(request.dictionaryidInt, request.Categorykey, request.Dictkey, request.Dictvalue, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionary.UpdateDictionaryInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【字典表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

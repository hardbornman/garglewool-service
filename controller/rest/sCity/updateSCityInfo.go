package sCity

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"time"
)

// 修改【市】信息接口

// 请求
type updateSCityInfoRequest struct {

	// 市ID
	cityId    string
	cityIdInt int64

	// 市名称
	CityName string `form:"cityName"`

	// 编码
	ZipCode string `form:"zipCode"`

	// 省ID
	ProvinceId int64 `form:"provinceId"`

	// 创建日期
	DateCreated     string `form:"dateCreated"`
	dateCreatedTime time.Time

	// 更新日期
	DateUpdated     string `form:"dateUpdated"`
	dateUpdatedTime time.Time
}

// 方法
func UpdateSCityInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sCity.UpdateSCityInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateSCityInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证cityId参数
	request.cityId = c.Param("id")
	if golibs.Length(request.cityId) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.cityId is null",
			Message: "缺少【市ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.cityId) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.cityId is number",
			Message: "【市ID】参数格式不正确",
		})
		return
	}
	request.cityIdInt, err = strconv.ParseInt(request.cityId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.cityId parse err",
			Message: "【市ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.cityIdInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.cityId value err",
			Message: "【市ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证cityName参数
	if golibs.Length(request.CityName) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.cityName is null",
			Message: "缺少【市名称】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.CityName) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.cityName format err",
			Message: "【市名称】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.CityName) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.cityName length err",
			Message: "【市名称】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证zipCode参数
	if golibs.Length(request.ZipCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.zipCode is null",
			Message: "缺少【编码】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.ZipCode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.zipCode format err",
			Message: "【编码】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.ZipCode) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.zipCode length err",
			Message: "【编码】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证provinceId参数
	if request.ProvinceId <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.provinceId value err",
			Message: "【省ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证dateCreated参数
	if golibs.Length(request.DateCreated) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.dateCreated length err",
			Message: "缺少【创建日期】参数",
		})
		return
	}
	request.dateCreatedTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.DateCreated, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.dateCreated parse err",
			Message: "【创建日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.dateCreatedTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.dateCreated value err",
			Message: "【创建日期】参数值错误:" + request.DateCreated,
		})
		return
	}
	//endregion

	//region 验证dateUpdated参数
	if golibs.Length(request.DateUpdated) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.dateUpdated length err",
			Message: "缺少【更新日期】参数",
		})
		return
	}
	request.dateUpdatedTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.DateUpdated, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.dateUpdated parse err",
			Message: "【更新日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.dateUpdatedTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.dateUpdated value err",
			Message: "【更新日期】参数值错误:" + request.DateUpdated,
		})
		return
	}
	//endregion

	//region 修改【市】信息
	isSuccess, err := service.UpdateSCityInfo(request.cityIdInt, request.CityName, request.ZipCode, request.ProvinceId, request.dateCreatedTime, request.dateUpdatedTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.UpdateSCityInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【市】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

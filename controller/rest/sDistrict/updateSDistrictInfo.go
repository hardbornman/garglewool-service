package sDistrict

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"time"
)

// 修改【区县】信息接口

// 请求
type updateSDistrictInfoRequest struct {

	// 区ID
	districtId    string
	districtIdInt int64

	// 区名称
	DistrictName string `form:"districtName"`

	// 市ID
	CityId int64 `form:"cityId"`

	// 创建日期
	DateCreated     string `form:"dateCreated"`
	dateCreatedTime time.Time

	// 更新日期
	DateUpdated     string `form:"dateUpdated"`
	dateUpdatedTime time.Time
}

// 方法
func UpdateSDistrictInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sDistrict.UpdateSDistrictInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateSDistrictInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证districtId参数
	request.districtId = c.Param("id")
	if golibs.Length(request.districtId) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.districtId is null",
			Message: "缺少【区ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.districtId) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.districtId is number",
			Message: "【区ID】参数格式不正确",
		})
		return
	}
	request.districtIdInt, err = strconv.ParseInt(request.districtId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.districtId parse err",
			Message: "【区ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.districtIdInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.districtId value err",
			Message: "【区ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证districtName参数
	if golibs.Length(request.DistrictName) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.districtName is null",
			Message: "缺少【区名称】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.DistrictName) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.districtName format err",
			Message: "【区名称】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.DistrictName) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.districtName length err",
			Message: "【区名称】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证cityId参数
	if request.CityId <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.cityId value err",
			Message: "【市ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证dateCreated参数
	if golibs.Length(request.DateCreated) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.dateCreated length err",
			Message: "缺少【创建日期】参数",
		})
		return
	}
	request.dateCreatedTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.DateCreated, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.dateCreated parse err",
			Message: "【创建日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.dateCreatedTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.dateCreated value err",
			Message: "【创建日期】参数值错误:" + request.DateCreated,
		})
		return
	}
	//endregion

	//region 验证dateUpdated参数
	if golibs.Length(request.DateUpdated) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.dateUpdated length err",
			Message: "缺少【更新日期】参数",
		})
		return
	}
	request.dateUpdatedTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.DateUpdated, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.dateUpdated parse err",
			Message: "【更新日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.dateUpdatedTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.dateUpdated value err",
			Message: "【更新日期】参数值错误:" + request.DateUpdated,
		})
		return
	}
	//endregion

	//region 修改【区县】信息
	isSuccess, err := service.UpdateSDistrictInfo(request.districtIdInt, request.DistrictName, request.CityId, request.dateCreatedTime, request.dateUpdatedTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.UpdateSDistrictInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【区县】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

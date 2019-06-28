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

// 插入【区县】信息接口

// 请求
type insertSDistrictInfoRequest struct {

	// 区ID
	DistrictId int64 `form:"districtId"`

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
func InsertSDistrictInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sDistrict.InsertSDistrictInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request insertSDistrictInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.InsertSDistrictInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证districtId参数
	if request.DistrictId <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.InsertSDistrictInfo.districtId value err",
			Message: "districtId参数值错误",
		})
		return
	}
	//endregion

	//region 验证districtName参数
	if golibs.Length(request.DistrictName) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.InsertSDistrictInfo.districtName is null",
			Message: "缺少【区名称】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.DistrictName) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.InsertSDistrictInfo.districtName format err",
			Message: "【区名称】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.DistrictName) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.InsertSDistrictInfo.districtName length err",
			Message: "【区名称】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证cityId参数
	if request.CityId <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.InsertSDistrictInfo.cityId value err",
			Message: "cityId参数值错误",
		})
		return
	}
	//endregion

	//region 验证dateCreated参数
	if golibs.Length(request.DateCreated) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.InsertSDistrictInfo.dateCreated length err",
			Message: "缺少【创建日期】参数",
		})
		return
	}
	request.dateCreatedTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.DateCreated, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.InsertSDistrictInfo.dateCreated parse err",
			Message: "【创建日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.dateCreatedTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.InsertSDistrictInfo.dateCreated value err",
			Message: "【创建日期】参数值错误:" + request.DateCreated,
		})
		return
	}
	//endregion

	//region 验证dateUpdated参数
	if golibs.Length(request.DateUpdated) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.InsertSDistrictInfo.dateUpdated length err",
			Message: "缺少【更新日期】参数",
		})
		return
	}
	request.dateUpdatedTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.DateUpdated, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.InsertSDistrictInfo.dateUpdated parse err",
			Message: "【更新日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.dateUpdatedTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.InsertSDistrictInfo.dateUpdated value err",
			Message: "【更新日期】参数值错误:" + request.DateUpdated,
		})
		return
	}
	//endregion

	//region 插入【区县】信息
	isSuccess, err := service.InsertSDistrictInfo(request.DistrictId, request.DistrictName, request.CityId, request.dateCreatedTime, request.dateUpdatedTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.InsertSDistrictInfo.insert err",
			Message: "插入出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.InsertSDistrictInfo.insert failure",
			Message: "插入失败",
		})
		return
	}
	//endregion

	//region 返回插入【区县】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

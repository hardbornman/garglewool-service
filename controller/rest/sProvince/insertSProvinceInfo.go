package sProvince

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"time"
)

// 插入【省，直辖市】信息接口

// 请求
type insertSProvinceInfoRequest struct {

	// 省ID
	ProvinceId int64 `form:"provinceId"`

	// 省名称
	ProvinceName string `form:"provinceName"`

	// 创建日期
	DateCreated     string `form:"dateCreated"`
	dateCreatedTime time.Time

	// 更新日期
	DateUpdated     string `form:"dateUpdated"`
	dateUpdatedTime time.Time
}

// 方法
func InsertSProvinceInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sProvince.InsertSProvinceInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request insertSProvinceInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.InsertSProvinceInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证provinceId参数
	if request.ProvinceId <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.InsertSProvinceInfo.provinceId value err",
			Message: "provinceId参数值错误",
		})
		return
	}
	//endregion

	//region 验证provinceName参数
	if golibs.Length(request.ProvinceName) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.InsertSProvinceInfo.provinceName is null",
			Message: "缺少【省名称】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.ProvinceName) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.InsertSProvinceInfo.provinceName format err",
			Message: "【省名称】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.ProvinceName) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.InsertSProvinceInfo.provinceName length err",
			Message: "【省名称】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证dateCreated参数
	if golibs.Length(request.DateCreated) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.InsertSProvinceInfo.dateCreated length err",
			Message: "缺少【创建日期】参数",
		})
		return
	}
	request.dateCreatedTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.DateCreated, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.InsertSProvinceInfo.dateCreated parse err",
			Message: "【创建日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.dateCreatedTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.InsertSProvinceInfo.dateCreated value err",
			Message: "【创建日期】参数值错误:" + request.DateCreated,
		})
		return
	}
	//endregion

	//region 验证dateUpdated参数
	if golibs.Length(request.DateUpdated) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.InsertSProvinceInfo.dateUpdated length err",
			Message: "缺少【更新日期】参数",
		})
		return
	}
	request.dateUpdatedTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.DateUpdated, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.InsertSProvinceInfo.dateUpdated parse err",
			Message: "【更新日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.dateUpdatedTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.InsertSProvinceInfo.dateUpdated value err",
			Message: "【更新日期】参数值错误:" + request.DateUpdated,
		})
		return
	}
	//endregion

	//region 插入【省，直辖市】信息
	isSuccess, err := service.InsertSProvinceInfo(request.ProvinceId, request.ProvinceName, request.dateCreatedTime, request.dateUpdatedTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.InsertSProvinceInfo.insert err",
			Message: "插入出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.InsertSProvinceInfo.insert failure",
			Message: "插入失败",
		})
		return
	}
	//endregion

	//region 返回插入【省，直辖市】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

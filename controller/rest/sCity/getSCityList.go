package sCity

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

// 获取【市】列表接口

// 请求
type getSCityListRequest struct {

	// 当前页码
	pageIndex    string
	pageIndexInt int

	// 每页记录数
	pageSize    string
	pageSizeInt int

	// 市名称
	cityName string

	// 省ID
	provinceId    string
	provinceIdInt int64
}

// 方法
func GetSCityList(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sCity.GetSCityList.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getSCityListRequest
	//endregion

	//region 验证【当前页码】参数
	request.pageIndexInt = 1
	request.pageIndex = c.DefaultQuery("pageIndex", "")
	if golibs.Length(request.pageIndex) > 0 {
		if !golibs.IsNumber(request.pageIndex) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sCity.GetSCityList.pageIndex is number",
				Message: "pageIndex参数格式不正确",
			})
			return
		}
		request.pageIndexInt, err = strconv.Atoi(request.pageIndex)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sCity.GetSCityList.pageIndex parse err",
				Message: "pageIndex参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageIndexInt < 1 || request.pageIndexInt > 100000 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sCity.GetSCityList.pageIndex value err",
				Message: "pageIndex参数值错误",
			})
			return
		}
	}
	//endregion

	//region 验证【每页记录数】参数
	request.pageSizeInt = 15
	request.pageSize = c.DefaultQuery("pageSize", "")
	if golibs.Length(request.pageSize) > 0 {
		if !golibs.IsNumber(request.pageSize) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sCity.GetSCityList.pageSize is number",
				Message: "pageSize参数格式不正确",
			})
			return
		}
		request.pageSizeInt, err = strconv.Atoi(request.pageSize)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sCity.GetSCityList.pageSize parse err",
				Message: "pageSize参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageSizeInt < 1 || request.pageSizeInt > 500 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sCity.GetSCityList.pageSize value err",
				Message: "pageSize参数值错误",
			})
			return
		}
	}
	//endregion

	//region 验证【市名称】参数,可选
	request.cityName = c.DefaultQuery("cityName", "")
	if golibs.Length(request.cityName) > 0 {
		if !golibs.IsGeneralString(request.cityName) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sCity.GetSCityList.cityName format err",
				Message: "【市名称】参数格式不正确",
			})
			return
		}
		if golibs.Length(request.cityName) > 50 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sCity.GetSCityList.cityName length err",
				Message: "【市名称】参数长度不能超过50个字符",
			})
			return
		}
	}
	//endregion

	//region 验证【省ID】参数,可选
	request.provinceId = c.DefaultQuery("provinceId", "")
	if golibs.Length(request.provinceId) > 0 {
		if !golibs.IsNumber(request.provinceId) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sCity.GetSCityList.provinceId is not a number",
				Message: "【省ID】参数格式不正确",
			})
			return
		}
		request.provinceIdInt, err = strconv.ParseInt(request.provinceId, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sCity.GetSCityList.provinceId parse err",
				Message: "【省ID】参数解析出错:" + err.Error(),
			})
			return
		}
		if request.provinceIdInt <= 0 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sCity.GetSCityList.provinceId value err",
				Message: "【省ID】参数值错误",
			})
			return
		}
	}
	//endregion

	//region 查询【市】列表
	list, total, err := service.GetSCityList(request.cityName, request.provinceIdInt, request.pageIndexInt, request.pageSizeInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.GetSCityList.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 组装【市】列表
	sCitysArray := make([]map[string]interface{}, len(list))
	if len(list) > 0 {
		for i, v := range list {
			sCitysArray[i] = map[string]interface{}{
				"cityId":      v.CityId,                                       //市ID
				"cityName":    v.CityName,                                     //市名称
				"zipCode":     v.ZipCode,                                      //编码
				"provinceId":  v.ProvinceId,                                   //省ID
				"dateCreated": v.DateCreated.Format(golibs.Time_TIMEStandard), //创建日期
				"dateUpdated": v.DateUpdated.Format(golibs.Time_TIMEStandard), //更新日期
			}
		}
	}
	//endregion

	//region 返回【市】列表
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"total": total,
			"list":  sCitysArray,
		},
	})
	//endregion
}

package sDistrict

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

// 获取【区县】列表接口

// 请求
type getSDistrictListRequest struct {

	// 当前页码
	pageIndex    string
	pageIndexInt int

	// 每页记录数
	pageSize    string
	pageSizeInt int

	// 区名称
	districtName string

	// 市ID
	cityId    string
	cityIdInt int64
}

// 方法
func GetSDistrictList(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sDistrict.GetSDistrictList.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getSDistrictListRequest
	//endregion

	//region 验证【当前页码】参数
	request.pageIndexInt = 1
	request.pageIndex = c.DefaultQuery("pageIndex", "")
	if golibs.Length(request.pageIndex) > 0 {
		if !golibs.IsNumber(request.pageIndex) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sDistrict.GetSDistrictList.pageIndex is number",
				Message: "pageIndex参数格式不正确",
			})
			return
		}
		request.pageIndexInt, err = strconv.Atoi(request.pageIndex)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sDistrict.GetSDistrictList.pageIndex parse err",
				Message: "pageIndex参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageIndexInt < 1 || request.pageIndexInt > 100000 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sDistrict.GetSDistrictList.pageIndex value err",
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
				Code:    "sDistrict.GetSDistrictList.pageSize is number",
				Message: "pageSize参数格式不正确",
			})
			return
		}
		request.pageSizeInt, err = strconv.Atoi(request.pageSize)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sDistrict.GetSDistrictList.pageSize parse err",
				Message: "pageSize参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageSizeInt < 1 || request.pageSizeInt > 500 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sDistrict.GetSDistrictList.pageSize value err",
				Message: "pageSize参数值错误",
			})
			return
		}
	}
	//endregion

	//region 验证【区名称】参数,可选
	request.districtName = c.DefaultQuery("districtName", "")
	if golibs.Length(request.districtName) > 0 {
		if !golibs.IsGeneralString(request.districtName) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sDistrict.GetSDistrictList.districtName format err",
				Message: "【区名称】参数格式不正确",
			})
			return
		}
		if golibs.Length(request.districtName) > 50 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sDistrict.GetSDistrictList.districtName length err",
				Message: "【区名称】参数长度不能超过50个字符",
			})
			return
		}
	}
	//endregion

	//region 验证【市ID】参数,可选
	request.cityId = c.DefaultQuery("cityId", "")
	if golibs.Length(request.cityId) > 0 {
		if !golibs.IsNumber(request.cityId) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sDistrict.GetSDistrictList.cityId is not a number",
				Message: "【市ID】参数格式不正确",
			})
			return
		}
		request.cityIdInt, err = strconv.ParseInt(request.cityId, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sDistrict.GetSDistrictList.cityId parse err",
				Message: "【市ID】参数解析出错:" + err.Error(),
			})
			return
		}
		if request.cityIdInt <= 0 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sDistrict.GetSDistrictList.cityId value err",
				Message: "【市ID】参数值错误",
			})
			return
		}
	}
	//endregion

	//region 查询【区县】列表
	list, total, err := service.GetSDistrictList(request.districtName, request.cityIdInt, request.pageIndexInt, request.pageSizeInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.GetSDistrictList.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 组装【区县】列表
	sDistrictsArray := make([]map[string]interface{}, len(list))
	if len(list) > 0 {
		for i, v := range list {
			sDistrictsArray[i] = map[string]interface{}{
				"districtId":   v.DistrictId,                                   //区ID
				"districtName": v.DistrictName,                                 //区名称
				"cityId":       v.CityId,                                       //市ID
				"dateCreated":  v.DateCreated.Format(golibs.Time_TIMEStandard), //创建日期
				"dateUpdated":  v.DateUpdated.Format(golibs.Time_TIMEStandard), //更新日期
			}
		}
	}
	//endregion

	//region 返回【区县】列表
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"total": total,
			"list":  sDistrictsArray,
		},
	})
	//endregion
}

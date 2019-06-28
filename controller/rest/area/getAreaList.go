package area

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【区域表】列表接口

// 请求
type getAreaListRequest struct {

	// 当前页码
	pageIndex    string
	pageIndexInt int

	// 每页记录数
	pageSize    string
	pageSizeInt int

	// 区域名称
	regionName string

	// 经度
	longitude string

	// 纬度
	latitude string
}

// 方法
func GetAreaList(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "area.GetAreaList.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getAreaListRequest
	//endregion

	//region 验证【当前页码】参数
	request.pageIndexInt = 1
	request.pageIndex = c.DefaultQuery("pageIndex", "")
	if golibs.Length(request.pageIndex) > 0 {
		if !golibs.IsNumber(request.pageIndex) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "area.GetAreaList.pageIndex is number",
				Message: "pageIndex参数格式不正确",
			})
			return
		}
		request.pageIndexInt, err = strconv.Atoi(request.pageIndex)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "area.GetAreaList.pageIndex parse err",
				Message: "pageIndex参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageIndexInt < 1 || request.pageIndexInt > 100000 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "area.GetAreaList.pageIndex value err",
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
				Code:    "area.GetAreaList.pageSize is number",
				Message: "pageSize参数格式不正确",
			})
			return
		}
		request.pageSizeInt, err = strconv.Atoi(request.pageSize)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "area.GetAreaList.pageSize parse err",
				Message: "pageSize参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageSizeInt < 1 || request.pageSizeInt > 500 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "area.GetAreaList.pageSize value err",
				Message: "pageSize参数值错误",
			})
			return
		}
	}
	//endregion

	//region 验证【区域名称】参数,可选
	request.regionName = c.DefaultQuery("regionName", "")
	if golibs.Length(request.regionName) > 0 {
		if !golibs.IsGeneralString(request.regionName) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "area.GetAreaList.regionName format err",
				Message: "【区域名称】参数格式不正确",
			})
			return
		}
		if golibs.Length(request.regionName) > 50 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "area.GetAreaList.regionName length err",
				Message: "【区域名称】参数长度不能超过50个字符",
			})
			return
		}
	}
	//endregion

	//region 验证【经度】参数,可选
	request.longitude = c.DefaultQuery("longitude", "")
	if golibs.Length(request.longitude) > 0 {
		if !golibs.IsGeneralString(request.longitude) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "area.GetAreaList.longitude format err",
				Message: "【经度】参数格式不正确",
			})
			return
		}
		if golibs.Length(request.longitude) > 20 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "area.GetAreaList.longitude length err",
				Message: "【经度】参数长度不能超过20个字符",
			})
			return
		}
	}
	//endregion

	//region 验证【纬度】参数,可选
	request.latitude = c.DefaultQuery("latitude", "")
	if golibs.Length(request.latitude) > 0 {
		if !golibs.IsGeneralString(request.latitude) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "area.GetAreaList.latitude format err",
				Message: "【纬度】参数格式不正确",
			})
			return
		}
		if golibs.Length(request.latitude) > 20 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "area.GetAreaList.latitude length err",
				Message: "【纬度】参数长度不能超过20个字符",
			})
			return
		}
	}
	//endregion

	//region 查询【区域表】列表
	list, total, err := service.GetAreaList(request.regionName, request.longitude, request.latitude, request.pageIndexInt, request.pageSizeInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.GetAreaList.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 组装【区域表】列表
	areasArray := make([]map[string]interface{}, len(list))
	if len(list) > 0 {
		for i, v := range list {
			areasArray[i] = map[string]interface{}{
				"areaid":     v.Areaid,     //序号
				"regionName": v.RegionName, //区域名称
				"regionCode": v.RegionCode, //区域编码
				"parentCode": v.ParentCode, //父级编码
				"longitude":  v.Longitude,  //经度
				"latitude":   v.Latitude,   //纬度
			}
		}
	}
	//endregion

	//region 返回【区域表】列表
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"total": total,
			"list":  areasArray,
		},
	})
	//endregion
}

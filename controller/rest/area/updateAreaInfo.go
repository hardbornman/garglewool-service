package area

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
)

// 修改【区域表】信息接口

// 请求
type updateAreaInfoRequest struct {

	// 序号
	areaid    string
	areaidInt int

	// 区域名称
	RegionName string `form:"regionName"`

	// 区域编码
	RegionCode string `form:"regionCode"`

	// 父级编码
	ParentCode string `form:"parentCode"`

	// 经度
	Longitude string `form:"longitude"`

	// 纬度
	Latitude string `form:"latitude"`
}

// 方法
func UpdateAreaInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "area.UpdateAreaInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateAreaInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证areaid参数
	request.areaid = c.Param("id")
	if golibs.Length(request.areaid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.areaid is null",
			Message: "缺少【序号】参数",
		})
		return
	}
	if !golibs.IsNumber(request.areaid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.areaid is number",
			Message: "【序号】参数格式不正确",
		})
		return
	}
	request.areaidInt, err = strconv.Atoi(request.areaid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.areaid parse err",
			Message: "【序号】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.areaidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.areaid value err",
			Message: "【序号】参数值错误",
		})
		return
	}
	//endregion

	//region 验证regionName参数
	if golibs.Length(request.RegionName) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.regionName is null",
			Message: "缺少【区域名称】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.RegionName) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.regionName format err",
			Message: "【区域名称】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.RegionName) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.regionName length err",
			Message: "【区域名称】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证regionCode参数
	if golibs.Length(request.RegionCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.regionCode is null",
			Message: "缺少【区域编码】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.RegionCode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.regionCode format err",
			Message: "【区域编码】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.RegionCode) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.regionCode length err",
			Message: "【区域编码】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证parentCode参数
	if golibs.Length(request.ParentCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.parentCode is null",
			Message: "缺少【父级编码】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.ParentCode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.parentCode format err",
			Message: "【父级编码】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.ParentCode) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.parentCode length err",
			Message: "【父级编码】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证longitude参数
	if golibs.Length(request.Longitude) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.longitude is null",
			Message: "缺少【经度】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Longitude) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.longitude format err",
			Message: "【经度】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Longitude) > 20 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.longitude length err",
			Message: "【经度】参数长度不能超过20个字符",
		})
		return
	}
	//endregion

	//region 验证latitude参数
	if golibs.Length(request.Latitude) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.latitude is null",
			Message: "缺少【纬度】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Latitude) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.latitude format err",
			Message: "【纬度】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Latitude) > 20 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.latitude length err",
			Message: "【纬度】参数长度不能超过20个字符",
		})
		return
	}
	//endregion

	//region 修改【区域表】信息
	isSuccess, err := service.UpdateAreaInfo(request.areaidInt, request.RegionName, request.RegionCode, request.ParentCode, request.Longitude, request.Latitude)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.UpdateAreaInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【区域表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

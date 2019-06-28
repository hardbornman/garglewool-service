package sProvince

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

// 获取【省，直辖市】列表接口

// 请求
type getSProvinceListRequest struct {

	// 当前页码
	pageIndex    string
	pageIndexInt int

	// 每页记录数
	pageSize    string
	pageSizeInt int

	// 省名称
	provinceName string
}

// 方法
func GetSProvinceList(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sProvince.GetSProvinceList.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getSProvinceListRequest
	//endregion

	//region 验证【当前页码】参数
	request.pageIndexInt = 1
	request.pageIndex = c.DefaultQuery("pageIndex", "")
	if golibs.Length(request.pageIndex) > 0 {
		if !golibs.IsNumber(request.pageIndex) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sProvince.GetSProvinceList.pageIndex is number",
				Message: "pageIndex参数格式不正确",
			})
			return
		}
		request.pageIndexInt, err = strconv.Atoi(request.pageIndex)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sProvince.GetSProvinceList.pageIndex parse err",
				Message: "pageIndex参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageIndexInt < 1 || request.pageIndexInt > 100000 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sProvince.GetSProvinceList.pageIndex value err",
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
				Code:    "sProvince.GetSProvinceList.pageSize is number",
				Message: "pageSize参数格式不正确",
			})
			return
		}
		request.pageSizeInt, err = strconv.Atoi(request.pageSize)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sProvince.GetSProvinceList.pageSize parse err",
				Message: "pageSize参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageSizeInt < 1 || request.pageSizeInt > 500 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sProvince.GetSProvinceList.pageSize value err",
				Message: "pageSize参数值错误",
			})
			return
		}
	}
	//endregion

	//region 验证【省名称】参数,可选
	request.provinceName = c.DefaultQuery("provinceName", "")
	if golibs.Length(request.provinceName) > 0 {
		if !golibs.IsGeneralString(request.provinceName) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sProvince.GetSProvinceList.provinceName format err",
				Message: "【省名称】参数格式不正确",
			})
			return
		}
		if golibs.Length(request.provinceName) > 50 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sProvince.GetSProvinceList.provinceName length err",
				Message: "【省名称】参数长度不能超过50个字符",
			})
			return
		}
	}
	//endregion

	//region 查询【省，直辖市】列表
	list, total, err := service.GetSProvinceList(request.provinceName, request.pageIndexInt, request.pageSizeInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.GetSProvinceList.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 组装【省，直辖市】列表
	sProvincesArray := make([]map[string]interface{}, len(list))
	if len(list) > 0 {
		for i, v := range list {
			sProvincesArray[i] = map[string]interface{}{
				"provinceId":   v.ProvinceId,                                   //省ID
				"provinceName": v.ProvinceName,                                 //省名称
				"dateCreated":  v.DateCreated.Format(golibs.Time_TIMEStandard), //创建日期
				"dateUpdated":  v.DateUpdated.Format(golibs.Time_TIMEStandard), //更新日期
			}
		}
	}
	//endregion

	//region 返回【省，直辖市】列表
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"total": total,
			"list":  sProvincesArray,
		},
	})
	//endregion
}

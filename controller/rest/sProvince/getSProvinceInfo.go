package sProvince

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【省，直辖市】信息接口

// 请求
type getSProvinceInfoRequest struct {

	// 省ID
	provinceId    string
	provinceIdInt int64
}

// 方法
func GetSProvinceInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sProvince.GetSProvinceInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getSProvinceInfoRequest
	//endregion

	//region 验证provinceId参数
	request.provinceId = c.Param("id")
	if golibs.Length(request.provinceId) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.GetSProvinceInfo.provinceId is null",
			Message: "缺少【省ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.provinceId) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.GetSProvinceInfo.provinceId is number",
			Message: "【省ID】参数格式不正确",
		})
		return
	}
	request.provinceIdInt, err = strconv.ParseInt(request.provinceId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.GetSProvinceInfo.provinceId parse err",
			Message: "provinceId参数解析出错:" + err.Error(),
		})
		return
	}
	if request.provinceIdInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.GetSProvinceInfo.provinceId value err",
			Message: "【省ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【省，直辖市】信息
	sProvinceInfo, err := service.GetSProvinceInfo(request.provinceIdInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.GetSProvinceInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if sProvinceInfo.ProvinceId <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sProvince.GetSProvinceInfo.not found",
			Message: "没有找到【省，直辖市】信息",
		})
		return
	}
	//endregion

	//region 返回【省，直辖市】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"provinceId":   sProvinceInfo.ProvinceId,                                   //省ID
			"provinceName": sProvinceInfo.ProvinceName,                                 //省名称
			"dateCreated":  sProvinceInfo.DateCreated.Format(golibs.Time_TIMEStandard), //创建日期
			"dateUpdated":  sProvinceInfo.DateUpdated.Format(golibs.Time_TIMEStandard), //更新日期
		},
	})
	//endregion
}

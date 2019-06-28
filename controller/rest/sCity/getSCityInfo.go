package sCity

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【市】信息接口

// 请求
type getSCityInfoRequest struct {

	// 市ID
	cityId    string
	cityIdInt int64
}

// 方法
func GetSCityInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sCity.GetSCityInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getSCityInfoRequest
	//endregion

	//region 验证cityId参数
	request.cityId = c.Param("id")
	if golibs.Length(request.cityId) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.GetSCityInfo.cityId is null",
			Message: "缺少【市ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.cityId) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.GetSCityInfo.cityId is number",
			Message: "【市ID】参数格式不正确",
		})
		return
	}
	request.cityIdInt, err = strconv.ParseInt(request.cityId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.GetSCityInfo.cityId parse err",
			Message: "cityId参数解析出错:" + err.Error(),
		})
		return
	}
	if request.cityIdInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.GetSCityInfo.cityId value err",
			Message: "【市ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【市】信息
	sCityInfo, err := service.GetSCityInfo(request.cityIdInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.GetSCityInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if sCityInfo.CityId <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sCity.GetSCityInfo.not found",
			Message: "没有找到【市】信息",
		})
		return
	}
	//endregion

	//region 返回【市】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"cityId":      sCityInfo.CityId,                                       //市ID
			"cityName":    sCityInfo.CityName,                                     //市名称
			"zipCode":     sCityInfo.ZipCode,                                      //编码
			"provinceId":  sCityInfo.ProvinceId,                                   //省ID
			"dateCreated": sCityInfo.DateCreated.Format(golibs.Time_TIMEStandard), //创建日期
			"dateUpdated": sCityInfo.DateUpdated.Format(golibs.Time_TIMEStandard), //更新日期
		},
	})
	//endregion
}

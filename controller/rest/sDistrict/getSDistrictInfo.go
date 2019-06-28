package sDistrict

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【区县】信息接口

// 请求
type getSDistrictInfoRequest struct {

	// 区ID
	districtId    string
	districtIdInt int64
}

// 方法
func GetSDistrictInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sDistrict.GetSDistrictInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getSDistrictInfoRequest
	//endregion

	//region 验证districtId参数
	request.districtId = c.Param("id")
	if golibs.Length(request.districtId) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.GetSDistrictInfo.districtId is null",
			Message: "缺少【区ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.districtId) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.GetSDistrictInfo.districtId is number",
			Message: "【区ID】参数格式不正确",
		})
		return
	}
	request.districtIdInt, err = strconv.ParseInt(request.districtId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.GetSDistrictInfo.districtId parse err",
			Message: "districtId参数解析出错:" + err.Error(),
		})
		return
	}
	if request.districtIdInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.GetSDistrictInfo.districtId value err",
			Message: "【区ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【区县】信息
	sDistrictInfo, err := service.GetSDistrictInfo(request.districtIdInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.GetSDistrictInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if sDistrictInfo.DistrictId <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sDistrict.GetSDistrictInfo.not found",
			Message: "没有找到【区县】信息",
		})
		return
	}
	//endregion

	//region 返回【区县】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"districtId":   sDistrictInfo.DistrictId,                                   //区ID
			"districtName": sDistrictInfo.DistrictName,                                 //区名称
			"cityId":       sDistrictInfo.CityId,                                       //市ID
			"dateCreated":  sDistrictInfo.DateCreated.Format(golibs.Time_TIMEStandard), //创建日期
			"dateUpdated":  sDistrictInfo.DateUpdated.Format(golibs.Time_TIMEStandard), //更新日期
		},
	})
	//endregion
}

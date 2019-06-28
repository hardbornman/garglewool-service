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

// 获取【区域表】信息接口

// 请求
type getAreaInfoRequest struct {

	// 序号
	areaid    string
	areaidInt int
}

// 方法
func GetAreaInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "area.GetAreaInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getAreaInfoRequest
	//endregion

	//region 验证areaid参数
	request.areaid = c.Param("id")
	if golibs.Length(request.areaid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.GetAreaInfo.areaid is null",
			Message: "缺少【序号】参数",
		})
		return
	}
	if !golibs.IsNumber(request.areaid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.GetAreaInfo.areaid is number",
			Message: "【序号】参数格式不正确",
		})
		return
	}
	request.areaidInt, err = strconv.Atoi(request.areaid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.GetAreaInfo.areaid parse err",
			Message: "areaid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.areaidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.GetAreaInfo.areaid value err",
			Message: "【序号】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【区域表】信息
	areaInfo, err := service.GetAreaInfo(request.areaidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.GetAreaInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if areaInfo.Areaid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "area.GetAreaInfo.not found",
			Message: "没有找到【区域表】信息",
		})
		return
	}
	//endregion

	//region 返回【区域表】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"areaid":     areaInfo.Areaid,     //序号
			"regionName": areaInfo.RegionName, //区域名称
			"regionCode": areaInfo.RegionCode, //区域编码
			"parentCode": areaInfo.ParentCode, //父级编码
			"longitude":  areaInfo.Longitude,  //经度
			"latitude":   areaInfo.Latitude,   //纬度
		},
	})
	//endregion
}

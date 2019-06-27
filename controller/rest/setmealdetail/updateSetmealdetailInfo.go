package setmealdetail

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

// 修改【套餐明细表】信息接口

// 请求
type updateSetmealdetailInfoRequest struct {

	// 套餐明细ID
	setmealdetailid    string
	setmealdetailidInt int

	// 套餐ID
	Setmealid int `form:"setmealid"`

	// 商品名称
	Name string `form:"name"`

	// 数量
	Nums int `form:"nums"`

	// 单价（元）
	Price float64 `form:"price"`

	// 创建人
	Adder int `form:"adder"`

	// 创建时间
	Addtime     string `form:"addtime"`
	addtimeTime time.Time

	// 修改人
	Moder int `form:"moder"`

	// 修改时间
	Modtime     string `form:"modtime"`
	modtimeTime time.Time
}

// 方法
func UpdateSetmealdetailInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "setmealdetail.UpdateSetmealdetailInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateSetmealdetailInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证setmealdetailid参数
	request.setmealdetailid = c.Param("id")
	if golibs.Length(request.setmealdetailid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.setmealdetailid is null",
			Message: "缺少【套餐明细ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.setmealdetailid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.setmealdetailid is number",
			Message: "【套餐明细ID】参数格式不正确",
		})
		return
	}
	request.setmealdetailidInt, err = strconv.Atoi(request.setmealdetailid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.setmealdetailid parse err",
			Message: "【套餐明细ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.setmealdetailidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.setmealdetailid value err",
			Message: "【套餐明细ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证setmealid参数
	if request.Setmealid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.setmealid value err",
			Message: "【套餐ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证name参数
	if golibs.Length(request.Name) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.name is null",
			Message: "缺少【商品名称】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Name) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.name format err",
			Message: "【商品名称】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Name) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.name length err",
			Message: "【商品名称】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证nums参数
	if request.Nums <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.nums value err",
			Message: "【数量】参数值错误",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.adder value err",
			Message: "【创建人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.moder value err",
			Message: "【修改人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 修改【套餐明细表】信息
	isSuccess, err := service.UpdateSetmealdetailInfo(request.setmealdetailidInt, request.Setmealid, request.Name, request.Nums, request.Price, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.UpdateSetmealdetailInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【套餐明细表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

package setmealdetail

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"time"
)

// 插入【套餐明细表】信息接口

// 请求
type insertSetmealdetailInfoRequest struct {

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
func InsertSetmealdetailInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "setmealdetail.InsertSetmealdetailInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request insertSetmealdetailInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证setmealid参数
	if request.Setmealid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.setmealid value err",
			Message: "setmealid参数值错误",
		})
		return
	}
	//endregion

	//region 验证name参数
	if golibs.Length(request.Name) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.name is null",
			Message: "缺少【商品名称】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Name) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.name format err",
			Message: "【商品名称】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Name) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.name length err",
			Message: "【商品名称】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证nums参数
	if request.Nums <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.nums value err",
			Message: "nums参数值错误",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.adder value err",
			Message: "adder参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.moder value err",
			Message: "moder参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 插入【套餐明细表】信息
	setmealdetailid, err := service.InsertSetmealdetailInfo(request.Setmealid, request.Name, request.Nums, request.Price, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.insert err",
			Message: "插入出错:" + err.Error(),
		})
		return
	}
	if setmealdetailid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.InsertSetmealdetailInfo.insert failure",
			Message: "插入失败",
		})
		return
	}
	//endregion

	//region 返回插入【套餐明细表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result":          true,
			"setmealdetailid": setmealdetailid,
		},
	})
	//endregion
}

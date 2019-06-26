package gwUser

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"time"
)

// 插入【用户管理】信息接口

// 请求
type insertGwUserInfoRequest struct {

	// 用户名
	UserName string `form:"user_name"`

	// 手机号
	UserPhone string `form:"user_phone"`

	// 用户金币数量
	UserGolds int `form:"user_golds"`

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
func InsertGwUserInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwUser.InsertGwUserInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request insertGwUserInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证user_name参数
	if golibs.Length(request.UserName) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.user_name is null",
			Message: "缺少【用户名】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.UserName) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.user_name format err",
			Message: "【用户名】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.UserName) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.user_name length err",
			Message: "【用户名】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证user_phone参数
	if golibs.Length(request.UserPhone) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.user_phone is null",
			Message: "缺少【手机号】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.UserPhone) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.user_phone format err",
			Message: "【手机号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.UserPhone) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.user_phone length err",
			Message: "【手机号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证user_golds参数
	if request.UserGolds <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.user_golds value err",
			Message: "user_golds参数值错误",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.adder value err",
			Message: "adder参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.moder value err",
			Message: "moder参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 插入【用户管理】信息
	userid, err := service.InsertGwUserInfo(request.UserName, request.UserPhone, request.UserGolds, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.insert err",
			Message: "插入出错:" + err.Error(),
		})
		return
	}
	if userid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.InsertGwUserInfo.insert failure",
			Message: "插入失败",
		})
		return
	}
	//endregion

	//region 返回插入【用户管理】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
			"userid": userid,
		},
	})
	//endregion
}

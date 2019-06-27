package userrole

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 修改【用户角色表】信息接口

// 请求
type updateUserroleInfoRequest struct {

	// 用户角色ID
	userroleid    string
	userroleidInt int

	// 角色名称
	Rolename string `form:"rolename"`

	// 描述
	Desc string `form:"desc"`

	// 备注
	Remark string `form:"remark"`
}

// 方法
func UpdateUserroleInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "userrole.UpdateUserroleInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateUserroleInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证userroleid参数
	request.userroleid = c.Param("id")
	if golibs.Length(request.userroleid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.userroleid is null",
			Message: "缺少【用户角色ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.userroleid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.userroleid is number",
			Message: "【用户角色ID】参数格式不正确",
		})
		return
	}
	request.userroleidInt, err = strconv.Atoi(request.userroleid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.userroleid parse err",
			Message: "【用户角色ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.userroleidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.userroleid value err",
			Message: "【用户角色ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证rolename参数
	if golibs.Length(request.Rolename) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.rolename is null",
			Message: "缺少【角色名称】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Rolename) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.rolename format err",
			Message: "【角色名称】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Rolename) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.rolename length err",
			Message: "【角色名称】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证desc参数
	if golibs.Length(request.Desc) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.desc is null",
			Message: "缺少【描述】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Desc) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.desc format err",
			Message: "【描述】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Desc) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.desc length err",
			Message: "【描述】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证remark参数
	if golibs.Length(request.Remark) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.remark is null",
			Message: "缺少【备注】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Remark) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.remark format err",
			Message: "【备注】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Remark) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.remark length err",
			Message: "【备注】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 修改【用户角色表】信息
	isSuccess, err := service.UpdateUserroleInfo(request.userroleidInt, request.Rolename, request.Desc, request.Remark)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.UpdateUserroleInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【用户角色表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

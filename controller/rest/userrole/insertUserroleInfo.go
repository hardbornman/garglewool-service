package userrole

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
)

// 插入【用户角色表】信息接口

// 请求
type insertUserroleInfoRequest struct {

	// 角色名称
	Rolename string `form:"rolename"`

	// 描述
	Desc string `form:"desc"`

	// 备注
	Remark string `form:"remark"`
}

// 方法
func InsertUserroleInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "userrole.InsertUserroleInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request insertUserroleInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.InsertUserroleInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证rolename参数
	if golibs.Length(request.Rolename) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.InsertUserroleInfo.rolename is null",
			Message: "缺少【角色名称】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Rolename) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.InsertUserroleInfo.rolename format err",
			Message: "【角色名称】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Rolename) > 50 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.InsertUserroleInfo.rolename length err",
			Message: "【角色名称】参数长度不能超过50个字符",
		})
		return
	}
	//endregion

	//region 验证desc参数
	if golibs.Length(request.Desc) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.InsertUserroleInfo.desc is null",
			Message: "缺少【描述】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Desc) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.InsertUserroleInfo.desc format err",
			Message: "【描述】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Desc) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.InsertUserroleInfo.desc length err",
			Message: "【描述】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证remark参数
	if golibs.Length(request.Remark) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.InsertUserroleInfo.remark is null",
			Message: "缺少【备注】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Remark) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.InsertUserroleInfo.remark format err",
			Message: "【备注】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Remark) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.InsertUserroleInfo.remark length err",
			Message: "【备注】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 插入【用户角色表】信息
	userroleid, err := service.InsertUserroleInfo(request.Rolename, request.Desc, request.Remark)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.InsertUserroleInfo.insert err",
			Message: "插入出错:" + err.Error(),
		})
		return
	}
	if userroleid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.InsertUserroleInfo.insert failure",
			Message: "插入失败",
		})
		return
	}
	//endregion

	//region 返回插入【用户角色表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result":     true,
			"userroleid": userroleid,
		},
	})
	//endregion
}

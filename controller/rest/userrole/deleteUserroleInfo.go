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

// 删除【用户角色表】信息接口

// 请求
type deleteUserroleInfoRequest struct {

	// 用户角色ID
	userroleid    string
	userroleidInt int
}

// 方法
func DeleteUserroleInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "userrole.DeleteUserroleInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteUserroleInfoRequest
	//endregion

	//region 验证userroleid参数
	request.userroleid = c.Param("id")
	if golibs.Length(request.userroleid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.DeleteUserroleInfo.userroleid is null",
			Message: "缺少【用户角色ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.userroleid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.DeleteUserroleInfo.userroleid is number",
			Message: "【用户角色ID】参数格式不正确",
		})
		return
	}
	request.userroleidInt, err = strconv.Atoi(request.userroleid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.DeleteUserroleInfo.userroleid parse err",
			Message: "【用户角色ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.userroleidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.DeleteUserroleInfo.userroleid value err",
			Message: "【用户角色ID】参数值错误",
		})
		return
	}
	//endregion

	//region 删除【用户角色表】信息
	isSuccess, err := service.DeleteUserroleInfo(request.userroleidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.DeleteUserroleInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.DeleteUserroleInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【用户角色表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

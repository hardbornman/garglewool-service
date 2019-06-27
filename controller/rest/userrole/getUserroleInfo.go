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

// 获取【用户角色表】信息接口

// 请求
type getUserroleInfoRequest struct {

	// 用户角色ID
	userroleid    string
	userroleidInt int
}

// 方法
func GetUserroleInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "userrole.GetUserroleInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getUserroleInfoRequest
	//endregion

	//region 验证userroleid参数
	request.userroleid = c.Param("id")
	if golibs.Length(request.userroleid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.GetUserroleInfo.userroleid is null",
			Message: "缺少【用户角色ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.userroleid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.GetUserroleInfo.userroleid is number",
			Message: "【用户角色ID】参数格式不正确",
		})
		return
	}
	request.userroleidInt, err = strconv.Atoi(request.userroleid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.GetUserroleInfo.userroleid parse err",
			Message: "userroleid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.userroleidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.GetUserroleInfo.userroleid value err",
			Message: "【用户角色ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【用户角色表】信息
	userroleInfo, err := service.GetUserroleInfo(request.userroleidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.GetUserroleInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if userroleInfo.Userroleid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.GetUserroleInfo.not found",
			Message: "没有找到【用户角色表】信息",
		})
		return
	}
	if userroleInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "userrole.GetUserroleInfo.has delete",
			Message: "【用户角色表】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【用户角色表】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"userroleid":   userroleInfo.Userroleid,                                  //用户角色ID
			"rolename":     userroleInfo.Rolename,                                    //角色名称
			"desc":         userroleInfo.Desc,                                        //描述
			"remark":       userroleInfo.Remark,                                      //备注
			"deleteStatus": userroleInfo.DeleteStatus,                                //0:未知，1：未删除，2：已删除
			"createtime":   userroleInfo.Createtime.Format(golibs.Time_TIMEStandard), //创建时间
			"updatetime":   userroleInfo.Updatetime.Format(golibs.Time_TIMEStandard), //更新时间
		},
	})
	//endregion
}

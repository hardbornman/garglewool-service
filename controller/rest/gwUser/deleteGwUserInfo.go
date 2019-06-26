package gwUser

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 删除【用户管理】信息接口

// 请求
type deleteGwUserInfoRequest struct {

	// 用户ID
	userid    string
	useridInt int
}

// 方法
func DeleteGwUserInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwUser.DeleteGwUserInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteGwUserInfoRequest
	//endregion

	//region 验证userid参数
	request.userid = c.Param("id")
	if golibs.Length(request.userid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.DeleteGwUserInfo.userid is null",
			Message: "缺少【用户ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.userid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.DeleteGwUserInfo.userid is number",
			Message: "【用户ID】参数格式不正确",
		})
		return
	}
	request.useridInt, err = strconv.Atoi(request.userid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.DeleteGwUserInfo.userid parse err",
			Message: "【用户ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.useridInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.DeleteGwUserInfo.userid value err",
			Message: "【用户ID】参数值错误",
		})
		return
	}
	//endregion

	//region 删除【用户管理】信息
	isSuccess, err := service.DeleteGwUserInfo(request.useridInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.DeleteGwUserInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.DeleteGwUserInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【用户管理】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

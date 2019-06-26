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

// 获取【用户管理】信息接口

// 请求
type getGwUserInfoRequest struct {

	// 用户ID
	userid    string
	useridInt int
}

// 方法
func GetGwUserInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwUser.GetGwUserInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getGwUserInfoRequest
	//endregion

	//region 验证userid参数
	request.userid = c.Param("id")
	if golibs.Length(request.userid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.GetGwUserInfo.userid is null",
			Message: "缺少【用户ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.userid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.GetGwUserInfo.userid is number",
			Message: "【用户ID】参数格式不正确",
		})
		return
	}
	request.useridInt, err = strconv.Atoi(request.userid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.GetGwUserInfo.userid parse err",
			Message: "userid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.useridInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.GetGwUserInfo.userid value err",
			Message: "【用户ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【用户管理】信息
	gwUserInfo, err := service.GetGwUserInfo(request.useridInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.GetGwUserInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if gwUserInfo.Userid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.GetGwUserInfo.not found",
			Message: "没有找到【用户管理】信息",
		})
		return
	}
	if gwUserInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwUser.GetGwUserInfo.has delete",
			Message: "【用户管理】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【用户管理】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"userid":       gwUserInfo.Userid,                                   //用户ID
			"user_name":    gwUserInfo.UserName,                                 //用户名
			"user_phone":   gwUserInfo.UserPhone,                                //手机号
			"user_golds":   gwUserInfo.UserGolds,                                //用户金币数量
			"adder":        gwUserInfo.Adder,                                    //创建人
			"addtime":      gwUserInfo.Addtime.Format(golibs.Time_TIMEStandard), //创建时间
			"moder":        gwUserInfo.Moder,                                    //修改人
			"modtime":      gwUserInfo.Modtime.Format(golibs.Time_TIMEStandard), //修改时间
			"deleteStatus": gwUserInfo.DeleteStatus,                             //0:未知，1：未删除，2：已删除
		},
	})
	//endregion
}

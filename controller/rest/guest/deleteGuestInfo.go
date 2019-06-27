package guest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 删除【买家客户表】信息接口

// 请求
type deleteGuestInfoRequest struct {

	// 用户ID
	guestid    string
	guestidInt int
}

// 方法
func DeleteGuestInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "guest.DeleteGuestInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteGuestInfoRequest
	//endregion

	//region 验证guestid参数
	request.guestid = c.Param("id")
	if golibs.Length(request.guestid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.DeleteGuestInfo.guestid is null",
			Message: "缺少【用户ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.guestid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.DeleteGuestInfo.guestid is number",
			Message: "【用户ID】参数格式不正确",
		})
		return
	}
	request.guestidInt, err = strconv.Atoi(request.guestid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.DeleteGuestInfo.guestid parse err",
			Message: "【用户ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.guestidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.DeleteGuestInfo.guestid value err",
			Message: "【用户ID】参数值错误",
		})
		return
	}
	//endregion

	//region 删除【买家客户表】信息
	isSuccess, err := service.DeleteGuestInfo(request.guestidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.DeleteGuestInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.DeleteGuestInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【买家客户表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

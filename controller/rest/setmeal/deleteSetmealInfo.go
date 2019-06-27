package setmeal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 删除【套餐表】信息接口

// 请求
type deleteSetmealInfoRequest struct {

	// 套餐ID
	setmealid    string
	setmealidInt int
}

// 方法
func DeleteSetmealInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "setmeal.DeleteSetmealInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteSetmealInfoRequest
	//endregion

	//region 验证setmealid参数
	request.setmealid = c.Param("id")
	if golibs.Length(request.setmealid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.DeleteSetmealInfo.setmealid is null",
			Message: "缺少【套餐ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.setmealid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.DeleteSetmealInfo.setmealid is number",
			Message: "【套餐ID】参数格式不正确",
		})
		return
	}
	request.setmealidInt, err = strconv.Atoi(request.setmealid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.DeleteSetmealInfo.setmealid parse err",
			Message: "【套餐ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.setmealidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.DeleteSetmealInfo.setmealid value err",
			Message: "【套餐ID】参数值错误",
		})
		return
	}
	//endregion

	//region 删除【套餐表】信息
	isSuccess, err := service.DeleteSetmealInfo(request.setmealidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.DeleteSetmealInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.DeleteSetmealInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【套餐表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

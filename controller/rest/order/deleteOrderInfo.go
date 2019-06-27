package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 删除【订单表】信息接口

// 请求
type deleteOrderInfoRequest struct {

	// 订单ID
	orderid    string
	orderidInt int
}

// 方法
func DeleteOrderInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "order.DeleteOrderInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteOrderInfoRequest
	//endregion

	//region 验证orderid参数
	request.orderid = c.Param("id")
	if golibs.Length(request.orderid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.DeleteOrderInfo.orderid is null",
			Message: "缺少【订单ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.orderid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.DeleteOrderInfo.orderid is number",
			Message: "【订单ID】参数格式不正确",
		})
		return
	}
	request.orderidInt, err = strconv.Atoi(request.orderid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.DeleteOrderInfo.orderid parse err",
			Message: "【订单ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.orderidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.DeleteOrderInfo.orderid value err",
			Message: "【订单ID】参数值错误",
		})
		return
	}
	//endregion

	//region 删除【订单表】信息
	isSuccess, err := service.DeleteOrderInfo(request.orderidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.DeleteOrderInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "order.DeleteOrderInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【订单表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

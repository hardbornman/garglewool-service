package gwShop

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 删除【店铺管理】信息接口

// 请求
type deleteGwShopInfoRequest struct {

	// 店铺ID
	shopid    string
	shopidInt int
}

// 方法
func DeleteGwShopInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwShop.DeleteGwShopInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteGwShopInfoRequest
	//endregion

	//region 验证shopid参数
	request.shopid = c.Param("id")
	if golibs.Length(request.shopid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.DeleteGwShopInfo.shopid is null",
			Message: "缺少【店铺ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.shopid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.DeleteGwShopInfo.shopid is number",
			Message: "【店铺ID】参数格式不正确",
		})
		return
	}
	request.shopidInt, err = strconv.Atoi(request.shopid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.DeleteGwShopInfo.shopid parse err",
			Message: "【店铺ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.shopidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.DeleteGwShopInfo.shopid value err",
			Message: "【店铺ID】参数值错误",
		})
		return
	}
	//endregion

	//region 删除【店铺管理】信息
	isSuccess, err := service.DeleteGwShopInfo(request.shopidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.DeleteGwShopInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.DeleteGwShopInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【店铺管理】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

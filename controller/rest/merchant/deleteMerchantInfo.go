package merchant

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 删除【商家用户表】信息接口

// 请求
type deleteMerchantInfoRequest struct {

	// 商家ID
	merchantid    string
	merchantidInt int
}

// 方法
func DeleteMerchantInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "merchant.DeleteMerchantInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteMerchantInfoRequest
	//endregion

	//region 验证merchantid参数
	request.merchantid = c.Param("id")
	if golibs.Length(request.merchantid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.DeleteMerchantInfo.merchantid is null",
			Message: "缺少【商家ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.merchantid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.DeleteMerchantInfo.merchantid is number",
			Message: "【商家ID】参数格式不正确",
		})
		return
	}
	request.merchantidInt, err = strconv.Atoi(request.merchantid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.DeleteMerchantInfo.merchantid parse err",
			Message: "【商家ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.merchantidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.DeleteMerchantInfo.merchantid value err",
			Message: "【商家ID】参数值错误",
		})
		return
	}
	//endregion

	//region 删除【商家用户表】信息
	isSuccess, err := service.DeleteMerchantInfo(request.merchantidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.DeleteMerchantInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.DeleteMerchantInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【商家用户表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

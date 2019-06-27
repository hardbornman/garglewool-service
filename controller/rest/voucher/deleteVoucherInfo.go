package voucher

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 删除【抵用券管理】信息接口

// 请求
type deleteVoucherInfoRequest struct {

	// 抵用券ID
	voucherid    string
	voucheridInt int
}

// 方法
func DeleteVoucherInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "voucher.DeleteVoucherInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteVoucherInfoRequest
	//endregion

	//region 验证voucherid参数
	request.voucherid = c.Param("id")
	if golibs.Length(request.voucherid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.DeleteVoucherInfo.voucherid is null",
			Message: "缺少【抵用券ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.voucherid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.DeleteVoucherInfo.voucherid is number",
			Message: "【抵用券ID】参数格式不正确",
		})
		return
	}
	request.voucheridInt, err = strconv.Atoi(request.voucherid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.DeleteVoucherInfo.voucherid parse err",
			Message: "【抵用券ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.voucheridInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.DeleteVoucherInfo.voucherid value err",
			Message: "【抵用券ID】参数值错误",
		})
		return
	}
	//endregion

	//region 删除【抵用券管理】信息
	isSuccess, err := service.DeleteVoucherInfo(request.voucheridInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.DeleteVoucherInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.DeleteVoucherInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【抵用券管理】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

package gwPackagedetail

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 删除【套餐明细管理】信息接口

// 请求
type deleteGwPackagedetailInfoRequest struct {

	// 套餐明细ID
	packagedetailid    string
	packagedetailidInt int
}

// 方法
func DeleteGwPackagedetailInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwPackagedetail.DeleteGwPackagedetailInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteGwPackagedetailInfoRequest
	//endregion

	//region 验证packagedetailid参数
	request.packagedetailid = c.Param("id")
	if golibs.Length(request.packagedetailid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.DeleteGwPackagedetailInfo.packagedetailid is null",
			Message: "缺少【套餐明细ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.packagedetailid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.DeleteGwPackagedetailInfo.packagedetailid is number",
			Message: "【套餐明细ID】参数格式不正确",
		})
		return
	}
	request.packagedetailidInt, err = strconv.Atoi(request.packagedetailid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.DeleteGwPackagedetailInfo.packagedetailid parse err",
			Message: "【套餐明细ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.packagedetailidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.DeleteGwPackagedetailInfo.packagedetailid value err",
			Message: "【套餐明细ID】参数值错误",
		})
		return
	}
	//endregion

	//region 删除【套餐明细管理】信息
	isSuccess, err := service.DeleteGwPackagedetailInfo(request.packagedetailidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.DeleteGwPackagedetailInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.DeleteGwPackagedetailInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【套餐明细管理】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

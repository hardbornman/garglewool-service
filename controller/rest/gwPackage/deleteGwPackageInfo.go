package gwPackage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 删除【套餐管理】信息接口

// 请求
type deleteGwPackageInfoRequest struct {

	// 套餐ID
	packageid    string
	packageidInt int
}

// 方法
func DeleteGwPackageInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwPackage.DeleteGwPackageInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteGwPackageInfoRequest
	//endregion

	//region 验证packageid参数
	request.packageid = c.Param("id")
	if golibs.Length(request.packageid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.DeleteGwPackageInfo.packageid is null",
			Message: "缺少【套餐ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.packageid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.DeleteGwPackageInfo.packageid is number",
			Message: "【套餐ID】参数格式不正确",
		})
		return
	}
	request.packageidInt, err = strconv.Atoi(request.packageid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.DeleteGwPackageInfo.packageid parse err",
			Message: "【套餐ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.packageidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.DeleteGwPackageInfo.packageid value err",
			Message: "【套餐ID】参数值错误",
		})
		return
	}
	//endregion

	//region 删除【套餐管理】信息
	isSuccess, err := service.DeleteGwPackageInfo(request.packageidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.DeleteGwPackageInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.DeleteGwPackageInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【套餐管理】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

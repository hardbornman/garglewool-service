package sysDictionarycategory

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 删除【字典分类】信息接口

// 请求
type deleteSysDictionarycategoryInfoRequest struct {

	// 字典分类ID
	dictionarycategoryid    string
	dictionarycategoryidInt int
}

// 方法
func DeleteSysDictionarycategoryInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sysDictionarycategory.DeleteSysDictionarycategoryInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteSysDictionarycategoryInfoRequest
	//endregion

	//region 验证dictionarycategoryid参数
	request.dictionarycategoryid = c.Param("id")
	if golibs.Length(request.dictionarycategoryid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionarycategory.DeleteSysDictionarycategoryInfo.dictionarycategoryid is null",
			Message: "缺少【字典分类ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.dictionarycategoryid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionarycategory.DeleteSysDictionarycategoryInfo.dictionarycategoryid is number",
			Message: "【字典分类ID】参数格式不正确",
		})
		return
	}
	request.dictionarycategoryidInt, err = strconv.Atoi(request.dictionarycategoryid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionarycategory.DeleteSysDictionarycategoryInfo.dictionarycategoryid parse err",
			Message: "【字典分类ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.dictionarycategoryidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionarycategory.DeleteSysDictionarycategoryInfo.dictionarycategoryid value err",
			Message: "【字典分类ID】参数值错误",
		})
		return
	}
	//endregion

	//region 删除【字典分类】信息
	isSuccess, err := service.DeleteSysDictionarycategoryInfo(request.dictionarycategoryidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionarycategory.DeleteSysDictionarycategoryInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionarycategory.DeleteSysDictionarycategoryInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【字典分类】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

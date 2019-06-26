package sysDictionary

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 删除【字典表】信息接口

// 请求
type deleteSysDictionaryInfoRequest struct {

	// 字典ID
	dictionaryid    string
	dictionaryidInt int
}

// 方法
func DeleteSysDictionaryInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sysDictionary.DeleteSysDictionaryInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteSysDictionaryInfoRequest
	//endregion

	//region 验证dictionaryid参数
	request.dictionaryid = c.Param("id")
	if golibs.Length(request.dictionaryid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionary.DeleteSysDictionaryInfo.dictionaryid is null",
			Message: "缺少【字典ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.dictionaryid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionary.DeleteSysDictionaryInfo.dictionaryid is number",
			Message: "【字典ID】参数格式不正确",
		})
		return
	}
	request.dictionaryidInt, err = strconv.Atoi(request.dictionaryid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionary.DeleteSysDictionaryInfo.dictionaryid parse err",
			Message: "【字典ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.dictionaryidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionary.DeleteSysDictionaryInfo.dictionaryid value err",
			Message: "【字典ID】参数值错误",
		})
		return
	}
	//endregion

	//region 删除【字典表】信息
	isSuccess, err := service.DeleteSysDictionaryInfo(request.dictionaryidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionary.DeleteSysDictionaryInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionary.DeleteSysDictionaryInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【字典表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

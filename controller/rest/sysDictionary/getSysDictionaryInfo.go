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

// 获取【字典表】信息接口

// 请求
type getSysDictionaryInfoRequest struct {

	// 字典ID
	dictionaryid    string
	dictionaryidInt int
}

// 方法
func GetSysDictionaryInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "sysDictionary.GetSysDictionaryInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getSysDictionaryInfoRequest
	//endregion

	//region 验证dictionaryid参数
	request.dictionaryid = c.Param("id")
	if golibs.Length(request.dictionaryid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionary.GetSysDictionaryInfo.dictionaryid is null",
			Message: "缺少【字典ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.dictionaryid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionary.GetSysDictionaryInfo.dictionaryid is number",
			Message: "【字典ID】参数格式不正确",
		})
		return
	}
	request.dictionaryidInt, err = strconv.Atoi(request.dictionaryid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionary.GetSysDictionaryInfo.dictionaryid parse err",
			Message: "dictionaryid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.dictionaryidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionary.GetSysDictionaryInfo.dictionaryid value err",
			Message: "【字典ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【字典表】信息
	sysDictionaryInfo, err := service.GetSysDictionaryInfo(request.dictionaryidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionary.GetSysDictionaryInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if sysDictionaryInfo.Dictionaryid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionary.GetSysDictionaryInfo.not found",
			Message: "没有找到【字典表】信息",
		})
		return
	}
	if sysDictionaryInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "sysDictionary.GetSysDictionaryInfo.has delete",
			Message: "【字典表】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【字典表】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"dictionaryid": sysDictionaryInfo.Dictionaryid,                             //字典ID
			"categorykey":  sysDictionaryInfo.Categorykey,                              //字典类别key
			"dictkey":      sysDictionaryInfo.Dictkey,                                  //key
			"dictvalue":    sysDictionaryInfo.Dictvalue,                                //value
			"adder":        sysDictionaryInfo.Adder,                                    //创建人
			"addtime":      sysDictionaryInfo.Addtime.Format(golibs.Time_TIMEStandard), //创建时间
			"moder":        sysDictionaryInfo.Moder,                                    //修改人
			"modtime":      sysDictionaryInfo.Modtime.Format(golibs.Time_TIMEStandard), //修改时间
			"deleteStatus": sysDictionaryInfo.DeleteStatus,                             //0:未知，1：未删除，2：已删除
		},
	})
	//endregion
}

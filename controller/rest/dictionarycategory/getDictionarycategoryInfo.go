package dictionarycategory

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【字典分类】信息接口

// 请求
type getDictionarycategoryInfoRequest struct {

	// 字典分类ID
	dictionarycategoryid    string
	dictionarycategoryidInt int
}

// 方法
func GetDictionarycategoryInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "dictionarycategory.GetDictionarycategoryInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getDictionarycategoryInfoRequest
	//endregion

	//region 验证dictionarycategoryid参数
	request.dictionarycategoryid = c.Param("id")
	if golibs.Length(request.dictionarycategoryid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.GetDictionarycategoryInfo.dictionarycategoryid is null",
			Message: "缺少【字典分类ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.dictionarycategoryid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.GetDictionarycategoryInfo.dictionarycategoryid is number",
			Message: "【字典分类ID】参数格式不正确",
		})
		return
	}
	request.dictionarycategoryidInt, err = strconv.Atoi(request.dictionarycategoryid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.GetDictionarycategoryInfo.dictionarycategoryid parse err",
			Message: "dictionarycategoryid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.dictionarycategoryidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.GetDictionarycategoryInfo.dictionarycategoryid value err",
			Message: "【字典分类ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【字典分类】信息
	dictionarycategoryInfo, err := service.GetDictionarycategoryInfo(request.dictionarycategoryidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.GetDictionarycategoryInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if dictionarycategoryInfo.Dictionarycategoryid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.GetDictionarycategoryInfo.not found",
			Message: "没有找到【字典分类】信息",
		})
		return
	}
	if dictionarycategoryInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.GetDictionarycategoryInfo.has delete",
			Message: "【字典分类】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【字典分类】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"dictionarycategoryid": dictionarycategoryInfo.Dictionarycategoryid,                     //字典分类ID
			"categorykey":          dictionarycategoryInfo.Categorykey,                              //字典类别key
			"categoryvalue":        dictionarycategoryInfo.Categoryvalue,                            //字典类别value
			"adder":                dictionarycategoryInfo.Adder,                                    //创建人
			"addtime":              dictionarycategoryInfo.Addtime.Format(golibs.Time_TIMEStandard), //创建时间
			"moder":                dictionarycategoryInfo.Moder,                                    //修改人
			"modtime":              dictionarycategoryInfo.Modtime.Format(golibs.Time_TIMEStandard), //修改时间
			"deleteStatus":         dictionarycategoryInfo.DeleteStatus,                             //0:未知，1：未删除，2：已删除
		},
	})
	//endregion
}

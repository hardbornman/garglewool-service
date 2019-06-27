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

// 获取【字典分类】列表接口

// 请求
type getDictionarycategoryListRequest struct {

	// 当前页码
	pageIndex    string
	pageIndexInt int

	// 每页记录数
	pageSize    string
	pageSizeInt int
}

// 方法
func GetDictionarycategoryList(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "dictionarycategory.GetDictionarycategoryList.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getDictionarycategoryListRequest
	//endregion

	//region 验证【当前页码】参数
	request.pageIndexInt = 1
	request.pageIndex = c.DefaultQuery("pageIndex", "")
	if golibs.Length(request.pageIndex) > 0 {
		if !golibs.IsNumber(request.pageIndex) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "dictionarycategory.GetDictionarycategoryList.pageIndex is number",
				Message: "pageIndex参数格式不正确",
			})
			return
		}
		request.pageIndexInt, err = strconv.Atoi(request.pageIndex)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "dictionarycategory.GetDictionarycategoryList.pageIndex parse err",
				Message: "pageIndex参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageIndexInt < 1 || request.pageIndexInt > 100000 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "dictionarycategory.GetDictionarycategoryList.pageIndex value err",
				Message: "pageIndex参数值错误",
			})
			return
		}
	}
	//endregion

	//region 验证【每页记录数】参数
	request.pageSizeInt = 15
	request.pageSize = c.DefaultQuery("pageSize", "")
	if golibs.Length(request.pageSize) > 0 {
		if !golibs.IsNumber(request.pageSize) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "dictionarycategory.GetDictionarycategoryList.pageSize is number",
				Message: "pageSize参数格式不正确",
			})
			return
		}
		request.pageSizeInt, err = strconv.Atoi(request.pageSize)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "dictionarycategory.GetDictionarycategoryList.pageSize parse err",
				Message: "pageSize参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageSizeInt < 1 || request.pageSizeInt > 500 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "dictionarycategory.GetDictionarycategoryList.pageSize value err",
				Message: "pageSize参数值错误",
			})
			return
		}
	}
	//endregion

	//region 查询【字典分类】列表
	list, total, err := service.GetDictionarycategoryList(request.pageIndexInt, request.pageSizeInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "dictionarycategory.GetDictionarycategoryList.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 组装【字典分类】列表
	dictionarycategorysArray := make([]map[string]interface{}, len(list))
	if len(list) > 0 {
		for i, v := range list {
			dictionarycategorysArray[i] = map[string]interface{}{
				"dictionarycategoryid": v.Dictionarycategoryid,                     //字典分类ID
				"categorykey":          v.Categorykey,                              //字典类别key
				"categoryvalue":        v.Categoryvalue,                            //字典类别value
				"adder":                v.Adder,                                    //创建人
				"addtime":              v.Addtime.Format(golibs.Time_TIMEStandard), //创建时间
				"moder":                v.Moder,                                    //修改人
				"modtime":              v.Modtime.Format(golibs.Time_TIMEStandard), //修改时间
				"deleteStatus":         v.DeleteStatus,                             //0:未知，1：未删除，2：已删除
			}
		}
	}
	//endregion

	//region 返回【字典分类】列表
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"total": total,
			"list":  dictionarycategorysArray,
		},
	})
	//endregion
}

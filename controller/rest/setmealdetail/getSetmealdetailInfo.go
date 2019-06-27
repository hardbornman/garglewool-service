package setmealdetail

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【套餐明细表】信息接口

// 请求
type getSetmealdetailInfoRequest struct {

	// 套餐明细ID
	setmealdetailid    string
	setmealdetailidInt int
}

// 方法
func GetSetmealdetailInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "setmealdetail.GetSetmealdetailInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getSetmealdetailInfoRequest
	//endregion

	//region 验证setmealdetailid参数
	request.setmealdetailid = c.Param("id")
	if golibs.Length(request.setmealdetailid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.GetSetmealdetailInfo.setmealdetailid is null",
			Message: "缺少【套餐明细ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.setmealdetailid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.GetSetmealdetailInfo.setmealdetailid is number",
			Message: "【套餐明细ID】参数格式不正确",
		})
		return
	}
	request.setmealdetailidInt, err = strconv.Atoi(request.setmealdetailid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.GetSetmealdetailInfo.setmealdetailid parse err",
			Message: "setmealdetailid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.setmealdetailidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.GetSetmealdetailInfo.setmealdetailid value err",
			Message: "【套餐明细ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【套餐明细表】信息
	setmealdetailInfo, err := service.GetSetmealdetailInfo(request.setmealdetailidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.GetSetmealdetailInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if setmealdetailInfo.Setmealdetailid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.GetSetmealdetailInfo.not found",
			Message: "没有找到【套餐明细表】信息",
		})
		return
	}
	if setmealdetailInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmealdetail.GetSetmealdetailInfo.has delete",
			Message: "【套餐明细表】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【套餐明细表】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"setmealdetailid": setmealdetailInfo.Setmealdetailid,                          //套餐明细ID
			"setmealid":       setmealdetailInfo.Setmealid,                                //套餐ID
			"name":            setmealdetailInfo.Name,                                     //商品名称
			"nums":            setmealdetailInfo.Nums,                                     //数量
			"price":           setmealdetailInfo.Price,                                    //单价（元）
			"adder":           setmealdetailInfo.Adder,                                    //创建人
			"addtime":         setmealdetailInfo.Addtime.Format(golibs.Time_TIMEStandard), //创建时间
			"moder":           setmealdetailInfo.Moder,                                    //修改人
			"modtime":         setmealdetailInfo.Modtime.Format(golibs.Time_TIMEStandard), //修改时间
			"deleteStatus":    setmealdetailInfo.DeleteStatus,                             //0:未知，1：未删除，2：已删除
		},
	})
	//endregion
}

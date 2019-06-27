package setmeal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【套餐表】信息接口

// 请求
type getSetmealInfoRequest struct {

	// 套餐ID
	setmealid    string
	setmealidInt int
}

// 方法
func GetSetmealInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "setmeal.GetSetmealInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getSetmealInfoRequest
	//endregion

	//region 验证setmealid参数
	request.setmealid = c.Param("id")
	if golibs.Length(request.setmealid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.GetSetmealInfo.setmealid is null",
			Message: "缺少【套餐ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.setmealid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.GetSetmealInfo.setmealid is number",
			Message: "【套餐ID】参数格式不正确",
		})
		return
	}
	request.setmealidInt, err = strconv.Atoi(request.setmealid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.GetSetmealInfo.setmealid parse err",
			Message: "setmealid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.setmealidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.GetSetmealInfo.setmealid value err",
			Message: "【套餐ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【套餐表】信息
	setmealInfo, err := service.GetSetmealInfo(request.setmealidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.GetSetmealInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if setmealInfo.Setmealid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.GetSetmealInfo.not found",
			Message: "没有找到【套餐表】信息",
		})
		return
	}
	if setmealInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "setmeal.GetSetmealInfo.has delete",
			Message: "【套餐表】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【套餐表】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"setmealid":    setmealInfo.Setmealid,                                   //套餐ID
			"shopid":       setmealInfo.Shopid,                                      //店铺ID
			"pkgcode":      setmealInfo.Pkgcode,                                     //套餐号
			"setmealtype":  setmealInfo.Setmealtype,                                 //套餐类别
			"title":        setmealInfo.Title,                                       //套餐标题
			"people":       setmealInfo.People,                                      //人数
			"isorder":      setmealInfo.Isorder,                                     //是否需要预约
			"isrefund":     setmealInfo.Isrefund,                                    //是否支持退款
			"isinhouse":    setmealInfo.Isinhouse,                                   //是否仅限堂食
			"isnew":        setmealInfo.Isnew,                                       //是否新品
			"isrecommend":  setmealInfo.Isrecommend,                                 //是否强烈推荐
			"validdays":    setmealInfo.Validdays,                                   //有效期（天）
			"watchers":     setmealInfo.Watchers,                                    //今日关注人数
			"createtime":   setmealInfo.Createtime.Format(golibs.Time_TIMEStandard), //上架日期
			"validtime":    setmealInfo.Validtime.Format(golibs.Time_TIMEStandard),  //生效日期
			"exittime":     setmealInfo.Exittime.Format(golibs.Time_TIMEStandard),   //下架日期
			"links":        setmealInfo.Links,                                       //更多优惠链接地址
			"info":         setmealInfo.Info,                                        //套餐说明
			"adder":        setmealInfo.Adder,                                       //创建人
			"addtime":      setmealInfo.Addtime.Format(golibs.Time_TIMEStandard),    //创建时间
			"moder":        setmealInfo.Moder,                                       //修改人
			"modtime":      setmealInfo.Modtime.Format(golibs.Time_TIMEStandard),    //修改时间
			"deleteStatus": setmealInfo.DeleteStatus,                                //0:未知，1：未删除，2：已删除
		},
	})
	//endregion
}

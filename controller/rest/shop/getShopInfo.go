package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【店铺表】信息接口

// 请求
type getShopInfoRequest struct {

	// 店铺ID
	shopid    string
	shopidInt int
}

// 方法
func GetShopInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "shop.GetShopInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getShopInfoRequest
	//endregion

	//region 验证shopid参数
	request.shopid = c.Param("id")
	if golibs.Length(request.shopid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.GetShopInfo.shopid is null",
			Message: "缺少【店铺ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.shopid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.GetShopInfo.shopid is number",
			Message: "【店铺ID】参数格式不正确",
		})
		return
	}
	request.shopidInt, err = strconv.Atoi(request.shopid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.GetShopInfo.shopid parse err",
			Message: "shopid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.shopidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.GetShopInfo.shopid value err",
			Message: "【店铺ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【店铺表】信息
	shopInfo, err := service.GetShopInfo(request.shopidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.GetShopInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if shopInfo.Shopid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.GetShopInfo.not found",
			Message: "没有找到【店铺表】信息",
		})
		return
	}
	if shopInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.GetShopInfo.has delete",
			Message: "【店铺表】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【店铺表】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"shopid":       shopInfo.Shopid,                                      //店铺ID
			"shopcode":     shopInfo.Shopcode,                                    //店铺号
			"shopname":     shopInfo.Shopname,                                    //店铺名称
			"province":     shopInfo.Province,                                    //店铺地址-省
			"city":         shopInfo.City,                                        //店铺地址-市
			"district":     shopInfo.District,                                    //店铺地址-区
			"address":      shopInfo.Address,                                     //店铺详细地址
			"phone":        shopInfo.Phone,                                       //电话
			"leaguetime":   shopInfo.Leaguetime.Format(golibs.Time_TIMEStandard), //加盟平台日期
			"exittime":     shopInfo.Exittime.Format(golibs.Time_TIMEStandard),   //退出平台日期
			"adder":        shopInfo.Adder,                                       //创建人
			"addtime":      shopInfo.Addtime.Format(golibs.Time_TIMEStandard),    //创建时间
			"moder":        shopInfo.Moder,                                       //修改人
			"modtime":      shopInfo.Modtime.Format(golibs.Time_TIMEStandard),    //修改时间
			"deleteStatus": shopInfo.DeleteStatus,                                //0:未知，1：未删除，2：已删除
			"merchantid":   shopInfo.Merchantid,                                  //商家ID
			"longtitude":   shopInfo.Longtitude,                                  //经度
			"latitude":     shopInfo.Latitude,                                    //纬度
		},
	})
	//endregion
}

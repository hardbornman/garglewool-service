package gwShop

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【店铺管理】信息接口

// 请求
type getGwShopInfoRequest struct {

	// 店铺ID
	shopid    string
	shopidInt int
}

// 方法
func GetGwShopInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwShop.GetGwShopInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getGwShopInfoRequest
	//endregion

	//region 验证shopid参数
	request.shopid = c.Param("id")
	if golibs.Length(request.shopid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.GetGwShopInfo.shopid is null",
			Message: "缺少【店铺ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.shopid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.GetGwShopInfo.shopid is number",
			Message: "【店铺ID】参数格式不正确",
		})
		return
	}
	request.shopidInt, err = strconv.Atoi(request.shopid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.GetGwShopInfo.shopid parse err",
			Message: "shopid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.shopidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.GetGwShopInfo.shopid value err",
			Message: "【店铺ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【店铺管理】信息
	gwShopInfo, err := service.GetGwShopInfo(request.shopidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.GetGwShopInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if gwShopInfo.Shopid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.GetGwShopInfo.not found",
			Message: "没有找到【店铺管理】信息",
		})
		return
	}
	if gwShopInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwShop.GetGwShopInfo.has delete",
			Message: "【店铺管理】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【店铺管理】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"shopid":          gwShopInfo.Shopid,                                          //店铺ID
			"shop_code":       gwShopInfo.ShopCode,                                        //店铺号
			"shop_name":       gwShopInfo.ShopName,                                        //店铺名称
			"shop_province":   gwShopInfo.ShopProvince,                                    //店铺地址-省
			"shop_city":       gwShopInfo.ShopCity,                                        //店铺地址-市
			"shop_district":   gwShopInfo.ShopDistrict,                                    //店铺地址-区
			"shop_address":    gwShopInfo.ShopAddress,                                     //店铺详细地址
			"shop_phone":      gwShopInfo.ShopPhone,                                       //加盟平台日期
			"shop_createtime": gwShopInfo.ShopCreatetime.Format(golibs.Time_TIMEStandard), //加盟平台日期
			"shop_exittime":   gwShopInfo.ShopExittime.Format(golibs.Time_TIMEStandard),   //退出平台日期
			"adder":           gwShopInfo.Adder,                                           //创建人
			"addtime":         gwShopInfo.Addtime.Format(golibs.Time_TIMEStandard),        //创建时间
			"moder":           gwShopInfo.Moder,                                           //修改人
			"modtime":         gwShopInfo.Modtime.Format(golibs.Time_TIMEStandard),        //修改时间
			"deleteStatus":    gwShopInfo.DeleteStatus,                                    //0:未知，1：未删除，2：已删除
		},
	})
	//endregion
}

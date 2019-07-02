package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
	"time"
)

// 获取【店铺表】列表接口

// 请求
type getShopListRequest struct {

	// 当前页码
	pageIndex    string
	pageIndexInt int

	// 每页记录数
	pageSize    string
	pageSizeInt int

	// 店铺ID
	shopid    string
	shopidInt int

	// 店铺名称
	shopname string

	// 店铺详细地址
	address string

	// 加盟平台日期
	leaguetime     string
	leaguetimeTime time.Time
}

// 方法
//func GetShopList(c *gin.Context) {
//	defer func() {
//		if err := recover(); err != nil {
//			c.JSON(http.StatusOK, model.Response{
//				Code:    "shop.GetShopList.ex",
//				Message: fmt.Sprintf("系统错误:%v", err),
//			})
//			return
//		}
//	}()
//
//	//region 解析请求参数
//	var err error
//	var request getShopListRequest
//	//endregion
//
//	//region 验证【当前页码】参数
//	request.pageIndexInt = 1
//	request.pageIndex = c.DefaultQuery("pageIndex", "")
//	if golibs.Length(request.pageIndex) > 0 {
//		if !golibs.IsNumber(request.pageIndex) {
//			c.JSON(http.StatusOK, model.Response{
//				Code:    "shop.GetShopList.pageIndex is number",
//				Message: "pageIndex参数格式不正确",
//			})
//			return
//		}
//		request.pageIndexInt, err = strconv.Atoi(request.pageIndex)
//		if err != nil {
//			c.JSON(http.StatusOK, model.Response{
//				Code:    "shop.GetShopList.pageIndex parse err",
//				Message: "pageIndex参数解析出错:" + err.Error(),
//			})
//			return
//		}
//		if request.pageIndexInt < 1 || request.pageIndexInt > 100000 {
//			c.JSON(http.StatusOK, model.Response{
//				Code:    "shop.GetShopList.pageIndex value err",
//				Message: "pageIndex参数值错误",
//			})
//			return
//		}
//	}
//	//endregion
//
//	//region 验证【每页记录数】参数
//	request.pageSizeInt = 15
//	request.pageSize = c.DefaultQuery("pageSize", "")
//	if golibs.Length(request.pageSize) > 0 {
//		if !golibs.IsNumber(request.pageSize) {
//			c.JSON(http.StatusOK, model.Response{
//				Code:    "shop.GetShopList.pageSize is number",
//				Message: "pageSize参数格式不正确",
//			})
//			return
//		}
//		request.pageSizeInt, err = strconv.Atoi(request.pageSize)
//		if err != nil {
//			c.JSON(http.StatusOK, model.Response{
//				Code:    "shop.GetShopList.pageSize parse err",
//				Message: "pageSize参数解析出错:" + err.Error(),
//			})
//			return
//		}
//		if request.pageSizeInt < 1 || request.pageSizeInt > 500 {
//			c.JSON(http.StatusOK, model.Response{
//				Code:    "shop.GetShopList.pageSize value err",
//				Message: "pageSize参数值错误",
//			})
//			return
//		}
//	}
//	//endregion
//
//	//region 查询【店铺表】列表
//	list, total, err := service.GetShopList(request.pageIndexInt, request.pageSizeInt)
//	if err != nil {
//		c.JSON(http.StatusOK, model.Response{
//			Code:    "shop.GetShopList.query err",
//			Message: "查询出错:" + err.Error(),
//		})
//		return
//	}
//	//endregion
//
//	//region 组装【店铺表】列表
//	shopsArray := make([]map[string]interface{}, len(list))
//	if len(list) > 0 {
//		for i, v := range list {
//			shopsArray[i] = map[string]interface{}{
//				"shopid":       v.Shopid,                                      //店铺ID
//				"shopcode":     v.Shopcode,                                    //店铺号
//				"shopname":     v.Shopname,                                    //店铺名称
//				"province":     v.Province,                                    //店铺地址-省
//				"city":         v.City,                                        //店铺地址-市
//				"district":     v.District,                                    //店铺地址-区
//				"address":      v.Address,                                     //店铺详细地址
//				"phone":        v.Phone,                                       //电话
//				"leaguetime":   v.Leaguetime.Format(golibs.Time_TIMEStandard), //加盟平台日期
//				"exittime":     v.Exittime.Format(golibs.Time_TIMEStandard),   //退出平台日期
//				"adder":        v.Adder,                                       //创建人
//				"addtime":      v.Addtime.Format(golibs.Time_TIMEStandard),    //创建时间
//				"moder":        v.Moder,                                       //修改人
//				"modtime":      v.Modtime.Format(golibs.Time_TIMEStandard),    //修改时间
//				"deleteStatus": v.DeleteStatus,                                //0:未知，1：未删除，2：已删除
//				"merchantid":   v.Merchantid,                                  //商家ID
//				"longtitude":   v.Longtitude,                                  //经度
//				"latitude":     v.Latitude,                                    //纬度
//			}
//		}
//	}
//	//endregion
//
//	//region 返回【店铺表】列表
//	c.JSON(http.StatusOK, model.Response{
//		Code:    "ok",
//		Message: "",
//		Data: map[string]interface{}{
//			"total": total,
//			"list":  shopsArray,
//		},
//	})
//	//endregion
//}
func GetShopList(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "shop.GetShopList.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getShopListRequest
	//endregion

	//region 验证【当前页码】参数
	request.pageIndexInt = 1
	request.pageIndex = c.DefaultQuery("pageIndex", "")
	if golibs.Length(request.pageIndex) > 0 {
		if !golibs.IsNumber(request.pageIndex) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "shop.GetShopList.pageIndex is number",
				Message: "pageIndex参数格式不正确",
			})
			return
		}
		request.pageIndexInt, err = strconv.Atoi(request.pageIndex)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "shop.GetShopList.pageIndex parse err",
				Message: "pageIndex参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageIndexInt < 1 || request.pageIndexInt > 100000 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "shop.GetShopList.pageIndex value err",
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
				Code:    "shop.GetShopList.pageSize is number",
				Message: "pageSize参数格式不正确",
			})
			return
		}
		request.pageSizeInt, err = strconv.Atoi(request.pageSize)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "shop.GetShopList.pageSize parse err",
				Message: "pageSize参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageSizeInt < 1 || request.pageSizeInt > 500 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "shop.GetShopList.pageSize value err",
				Message: "pageSize参数值错误",
			})
			return
		}
	}
	//endregion

	//region 验证【店铺ID】参数,可选
	request.shopid = c.DefaultQuery("shopid", "")
	if golibs.Length(request.shopid) > 0 {
		if !golibs.IsNumber(request.shopid) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "shop.GetShopList.shopid is not a number",
				Message: "【店铺ID】参数格式不正确",
			})
			return
		}
		request.shopidInt, err = strconv.Atoi(request.shopid)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "shop.GetShopList.shopid parse err",
				Message: "【店铺ID】参数解析出错:" + err.Error(),
			})
			return
		}
		if request.shopidInt <= 0 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "shop.GetShopList.shopid value err",
				Message: "【店铺ID】参数值错误",
			})
			return
		}
	}
	//endregion

	//region 验证【店铺名称】参数,可选
	request.shopname = c.DefaultQuery("shopname", "")
	if golibs.Length(request.shopname) > 0 {
		if !golibs.IsGeneralString(request.shopname) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "shop.GetShopList.shopname format err",
				Message: "【店铺名称】参数格式不正确",
			})
			return
		}
		if golibs.Length(request.shopname) > 255 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "shop.GetShopList.shopname length err",
				Message: "【店铺名称】参数长度不能超过255个字符",
			})
			return
		}
	}
	//endregion

	//region 验证【店铺详细地址】参数,可选
	request.address = c.DefaultQuery("address", "")
	if golibs.Length(request.address) > 0 {
		if !golibs.IsGeneralString(request.address) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "shop.GetShopList.address format err",
				Message: "【店铺详细地址】参数格式不正确",
			})
			return
		}
		if golibs.Length(request.address) > 255 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "shop.GetShopList.address length err",
				Message: "【店铺详细地址】参数长度不能超过255个字符",
			})
			return
		}
	}
	//endregion

	//region 验证【加盟平台日期】参数,可选
	request.leaguetime = c.DefaultQuery("leaguetime", "2018-01-01 00:00:00")
	if golibs.Length(request.leaguetime) <= 0 {
		//c.JSON(http.StatusOK, model.Response{
		//	Code:    "shop.GetShopList.leaguetime is null",
		//	Message: "【加盟平台日期】参数不能为空",
		//})
		//return
		request.leaguetime = "2018-01-01 00:00:00"
	}
	if !golibs.IsStandardTime(request.leaguetime) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.GetShopList.leaguetime format err",
			Message: "【加盟平台日期】参数格式不正确",
		})
		return
	}
	request.leaguetimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.leaguetime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.GetShopList.leaguetime parse err",
			Message: "【加盟平台日期】参数解析错误:" + err.Error(),
		})
		return
	}
	//endregion

	//region 查询【店铺表】列表
	list, total, err := service.GetShopList(request.shopidInt, request.shopname, request.address, request.leaguetimeTime, request.pageIndexInt, request.pageSizeInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "shop.GetShopList.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 组装【店铺表】列表
	shopsArray := make([]map[string]interface{}, len(list))
	if len(list) > 0 {
		for i, v := range list {
			shopsArray[i] = map[string]interface{}{
				"shopid":       v.Shopid,                                      //店铺ID
				"shopcode":     v.Shopcode,                                    //店铺号
				"shopname":     v.Shopname,                                    //店铺名称
				"province":     v.Province,                                    //店铺地址-省
				"city":         v.City,                                        //店铺地址-市
				"district":     v.District,                                    //店铺地址-区
				"address":      v.Address,                                     //店铺详细地址
				"phone":        v.Phone,                                       //电话
				"leaguetime":   v.Leaguetime.Format(golibs.Time_TIMEStandard), //加盟平台日期
				"exittime":     v.Exittime.Format(golibs.Time_TIMEStandard),   //退出平台日期
				"adder":        v.Adder,                                       //创建人
				"addtime":      v.Addtime.Format(golibs.Time_TIMEStandard),    //创建时间
				"moder":        v.Moder,                                       //修改人
				"modtime":      v.Modtime.Format(golibs.Time_TIMEStandard),    //修改时间
				"deleteStatus": v.DeleteStatus,                                //0:未知，1：未删除，2：已删除
				"merchantid":   v.Merchantid,                                  //商家ID
				"longtitude":   v.Longtitude,                                  //经度
				"latitude":     v.Latitude,                                    //纬度
			}
		}
	}
	//endregion

	//region 返回【店铺表】列表
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"total": total,
			"list":  shopsArray,
		},
	})
	//endregion
}

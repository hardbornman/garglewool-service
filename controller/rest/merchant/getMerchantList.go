package merchant

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【商家用户表】列表接口

// 请求
type getMerchantListRequest struct {

	// 当前页码
	pageIndex    string
	pageIndexInt int

	// 每页记录数
	pageSize    string
	pageSizeInt int
}

// 方法
func GetMerchantList(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "merchant.GetMerchantList.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getMerchantListRequest
	//endregion

	//region 验证【当前页码】参数
	request.pageIndexInt = 1
	request.pageIndex = c.DefaultQuery("pageIndex", "")
	if golibs.Length(request.pageIndex) > 0 {
		if !golibs.IsNumber(request.pageIndex) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "merchant.GetMerchantList.pageIndex is number",
				Message: "pageIndex参数格式不正确",
			})
			return
		}
		request.pageIndexInt, err = strconv.Atoi(request.pageIndex)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "merchant.GetMerchantList.pageIndex parse err",
				Message: "pageIndex参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageIndexInt < 1 || request.pageIndexInt > 100000 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "merchant.GetMerchantList.pageIndex value err",
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
				Code:    "merchant.GetMerchantList.pageSize is number",
				Message: "pageSize参数格式不正确",
			})
			return
		}
		request.pageSizeInt, err = strconv.Atoi(request.pageSize)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "merchant.GetMerchantList.pageSize parse err",
				Message: "pageSize参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageSizeInt < 1 || request.pageSizeInt > 500 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "merchant.GetMerchantList.pageSize value err",
				Message: "pageSize参数值错误",
			})
			return
		}
	}
	//endregion

	//region 查询【商家用户表】列表
	list, total, err := service.GetMerchantList(request.pageIndexInt, request.pageSizeInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.GetMerchantList.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 组装【商家用户表】列表
	merchantsArray := make([]map[string]interface{}, len(list))
	if len(list) > 0 {
		for i, v := range list {
			merchantsArray[i] = map[string]interface{}{
				"merchantid":    v.Merchantid,                                     //商家ID
				"merchantname":  v.Merchantname,                                   //商家名
				"phone":         v.Phone,                                          //手机号
				"createtime":    v.Createtime.Format(golibs.Time_TIMEStandard),    //创建时间
				"updatetime":    v.Updatetime.Format(golibs.Time_TIMEStandard),    //修改时间
				"deleteStatus":  v.DeleteStatus,                                   //0:未知，1：未删除，2：已删除
				"userroleid":    v.Userroleid,                                     //用户角色ID
				"loginpwd":      v.Loginpwd,                                       //登录密码
				"loginaccount":  v.Loginaccount,                                   //登录账户
				"nick":          v.Nick,                                           //昵称
				"wechataccount": v.Wechataccount,                                  //微信账号
				"wechatsign":    v.Wechatsign,                                     //微信签名
				"remark":        v.Remark,                                         //备注
				"lastlogintime": v.Lastlogintime.Format(golibs.Time_TIMEStandard), //最近登录时间
				"lastloginaddr": v.Lastloginaddr,                                  //最近登录地址
				"rigstertime":   v.Rigstertime.Format(golibs.Time_TIMEStandard),   //注册时间
				"enable":        v.Enable,                                         //是否启用
				"addr":          v.Addr,                                           //地址
			}
		}
	}
	//endregion

	//region 返回【商家用户表】列表
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"total": total,
			"list":  merchantsArray,
		},
	})
	//endregion
}

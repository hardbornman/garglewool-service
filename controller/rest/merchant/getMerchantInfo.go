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

// 获取【商家用户表】信息接口

// 请求
type getMerchantInfoRequest struct {

	// 商家ID
	merchantid    string
	merchantidInt int
}

// 方法
func GetMerchantInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "merchant.GetMerchantInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getMerchantInfoRequest
	//endregion

	//region 验证merchantid参数
	request.merchantid = c.Param("id")
	if golibs.Length(request.merchantid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.GetMerchantInfo.merchantid is null",
			Message: "缺少【商家ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.merchantid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.GetMerchantInfo.merchantid is number",
			Message: "【商家ID】参数格式不正确",
		})
		return
	}
	request.merchantidInt, err = strconv.Atoi(request.merchantid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.GetMerchantInfo.merchantid parse err",
			Message: "merchantid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.merchantidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.GetMerchantInfo.merchantid value err",
			Message: "【商家ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【商家用户表】信息
	merchantInfo, err := service.GetMerchantInfo(request.merchantidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.GetMerchantInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if merchantInfo.Merchantid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.GetMerchantInfo.not found",
			Message: "没有找到【商家用户表】信息",
		})
		return
	}
	if merchantInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "merchant.GetMerchantInfo.has delete",
			Message: "【商家用户表】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【商家用户表】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"merchantid":    merchantInfo.Merchantid,                                     //商家ID
			"merchantname":  merchantInfo.Merchantname,                                   //商家名
			"phone":         merchantInfo.Phone,                                          //手机号
			"createtime":    merchantInfo.Createtime.Format(golibs.Time_TIMEStandard),    //创建时间
			"updatetime":    merchantInfo.Updatetime.Format(golibs.Time_TIMEStandard),    //修改时间
			"deleteStatus":  merchantInfo.DeleteStatus,                                   //0:未知，1：未删除，2：已删除
			"userroleid":    merchantInfo.Userroleid,                                     //用户角色ID
			"loginpwd":      merchantInfo.Loginpwd,                                       //登录密码
			"loginaccount":  merchantInfo.Loginaccount,                                   //登录账户
			"nick":          merchantInfo.Nick,                                           //昵称
			"wechataccount": merchantInfo.Wechataccount,                                  //微信账号
			"wechatsign":    merchantInfo.Wechatsign,                                     //微信签名
			"remark":        merchantInfo.Remark,                                         //备注
			"lastlogintime": merchantInfo.Lastlogintime.Format(golibs.Time_TIMEStandard), //最近登录时间
			"lastloginaddr": merchantInfo.Lastloginaddr,                                  //最近登录地址
			"rigstertime":   merchantInfo.Rigstertime.Format(golibs.Time_TIMEStandard),   //注册时间
			"enable":        merchantInfo.Enable,                                         //是否启用
			"addr":          merchantInfo.Addr,                                           //地址
		},
	})
	//endregion
}

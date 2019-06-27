package guest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【买家客户表】信息接口

// 请求
type getGuestInfoRequest struct {

	// 用户ID
	guestid    string
	guestidInt int
}

// 方法
func GetGuestInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "guest.GetGuestInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getGuestInfoRequest
	//endregion

	//region 验证guestid参数
	request.guestid = c.Param("id")
	if golibs.Length(request.guestid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.GetGuestInfo.guestid is null",
			Message: "缺少【用户ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.guestid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.GetGuestInfo.guestid is number",
			Message: "【用户ID】参数格式不正确",
		})
		return
	}
	request.guestidInt, err = strconv.Atoi(request.guestid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.GetGuestInfo.guestid parse err",
			Message: "guestid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.guestidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.GetGuestInfo.guestid value err",
			Message: "【用户ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【买家客户表】信息
	guestInfo, err := service.GetGuestInfo(request.guestidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.GetGuestInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if guestInfo.Guestid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.GetGuestInfo.not found",
			Message: "没有找到【买家客户表】信息",
		})
		return
	}
	if guestInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "guest.GetGuestInfo.has delete",
			Message: "【买家客户表】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【买家客户表】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"guestid":      guestInfo.Guestid,                                  //用户ID
			"name":         guestInfo.Name,                                     //用户名
			"password":     guestInfo.Password,                                 //用户密码
			"phone":        guestInfo.Phone,                                    //手机号
			"golds":        guestInfo.Golds,                                    //用户金币数量
			"adder":        guestInfo.Adder,                                    //创建人
			"addtime":      guestInfo.Addtime.Format(golibs.Time_TIMEStandard), //创建时间
			"moder":        guestInfo.Moder,                                    //修改人
			"modtime":      guestInfo.Modtime.Format(golibs.Time_TIMEStandard), //修改时间
			"deleteStatus": guestInfo.DeleteStatus,                             //0:未知，1：未删除，2：已删除
		},
	})
	//endregion
}

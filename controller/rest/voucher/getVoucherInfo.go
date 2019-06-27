package voucher

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【抵用券管理】信息接口

// 请求
type getVoucherInfoRequest struct {

	// 抵用券ID
	voucherid    string
	voucheridInt int
}

// 方法
func GetVoucherInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "voucher.GetVoucherInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getVoucherInfoRequest
	//endregion

	//region 验证voucherid参数
	request.voucherid = c.Param("id")
	if golibs.Length(request.voucherid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.GetVoucherInfo.voucherid is null",
			Message: "缺少【抵用券ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.voucherid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.GetVoucherInfo.voucherid is number",
			Message: "【抵用券ID】参数格式不正确",
		})
		return
	}
	request.voucheridInt, err = strconv.Atoi(request.voucherid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.GetVoucherInfo.voucherid parse err",
			Message: "voucherid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.voucheridInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.GetVoucherInfo.voucherid value err",
			Message: "【抵用券ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【抵用券管理】信息
	voucherInfo, err := service.GetVoucherInfo(request.voucheridInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.GetVoucherInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if voucherInfo.Voucherid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.GetVoucherInfo.not found",
			Message: "没有找到【抵用券管理】信息",
		})
		return
	}
	if voucherInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "voucher.GetVoucherInfo.has delete",
			Message: "【抵用券管理】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【抵用券管理】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"voucherid":    voucherInfo.Voucherid,                                   //抵用券ID
			"code":         voucherInfo.Code,                                        //抵用券号
			"guestid":      voucherInfo.Guestid,                                     //用户id
			"quota":        voucherInfo.Quota,                                       //额度
			"reduce":       voucherInfo.Reduce,                                      //抵消
			"createtime":   voucherInfo.Createtime.Format(golibs.Time_TIMEStandard), //创建时间
			"validdays":    voucherInfo.Validdays,                                   //有效期（天）
			"isinvalid":    voucherInfo.Isinvalid,                                   //是否失效
			"adder":        voucherInfo.Adder,                                       //创建人
			"addtime":      voucherInfo.Addtime.Format(golibs.Time_TIMEStandard),    //创建时间
			"moder":        voucherInfo.Moder,                                       //修改人
			"modtime":      voucherInfo.Modtime.Format(golibs.Time_TIMEStandard),    //修改时间
			"deleteStatus": voucherInfo.DeleteStatus,                                //0:未知，1：未删除，2：已删除
		},
	})
	//endregion
}

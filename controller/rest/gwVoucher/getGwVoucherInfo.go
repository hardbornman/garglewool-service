package gwVoucher

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
type getGwVoucherInfoRequest struct {

	// 抵用券ID
	voucherid    string
	voucheridInt int
}

// 方法
func GetGwVoucherInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwVoucher.GetGwVoucherInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getGwVoucherInfoRequest
	//endregion

	//region 验证voucherid参数
	request.voucherid = c.Param("id")
	if golibs.Length(request.voucherid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.GetGwVoucherInfo.voucherid is null",
			Message: "缺少【抵用券ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.voucherid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.GetGwVoucherInfo.voucherid is number",
			Message: "【抵用券ID】参数格式不正确",
		})
		return
	}
	request.voucheridInt, err = strconv.Atoi(request.voucherid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.GetGwVoucherInfo.voucherid parse err",
			Message: "voucherid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.voucheridInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.GetGwVoucherInfo.voucherid value err",
			Message: "【抵用券ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【抵用券管理】信息
	gwVoucherInfo, err := service.GetGwVoucherInfo(request.voucheridInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.GetGwVoucherInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if gwVoucherInfo.Voucherid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.GetGwVoucherInfo.not found",
			Message: "没有找到【抵用券管理】信息",
		})
		return
	}
	if gwVoucherInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwVoucher.GetGwVoucherInfo.has delete",
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
			"voucherid":          gwVoucherInfo.Voucherid,                                          //抵用券ID
			"voucher_code":       gwVoucherInfo.VoucherCode,                                        //抵用券号
			"voucher_userid":     gwVoucherInfo.VoucherUserid,                                      //用户id
			"voucher_quota":      gwVoucherInfo.VoucherQuota,                                       //额度
			"voucher_reduce":     gwVoucherInfo.VoucherReduce,                                      //抵消
			"voucher_createtime": gwVoucherInfo.VoucherCreatetime.Format(golibs.Time_TIMEStandard), //创建时间
			"voucher_validdays":  gwVoucherInfo.VoucherValiddays,                                   //有效期（天）
			"voucher_isinvalid":  gwVoucherInfo.VoucherIsinvalid,                                   //是否失效
			"adder":              gwVoucherInfo.Adder,                                              //创建人
			"addtime":            gwVoucherInfo.Addtime.Format(golibs.Time_TIMEStandard),           //创建时间
			"moder":              gwVoucherInfo.Moder,                                              //修改人
			"modtime":            gwVoucherInfo.Modtime.Format(golibs.Time_TIMEStandard),           //修改时间
			"deleteStatus":       gwVoucherInfo.DeleteStatus,                                       //0:未知，1：未删除，2：已删除
		},
	})
	//endregion
}

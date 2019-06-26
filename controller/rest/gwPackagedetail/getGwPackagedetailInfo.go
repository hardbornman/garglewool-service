package gwPackagedetail

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【套餐明细管理】信息接口

// 请求
type getGwPackagedetailInfoRequest struct {

	// 套餐明细ID
	packagedetailid    string
	packagedetailidInt int
}

// 方法
func GetGwPackagedetailInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwPackagedetail.GetGwPackagedetailInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getGwPackagedetailInfoRequest
	//endregion

	//region 验证packagedetailid参数
	request.packagedetailid = c.Param("id")
	if golibs.Length(request.packagedetailid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.GetGwPackagedetailInfo.packagedetailid is null",
			Message: "缺少【套餐明细ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.packagedetailid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.GetGwPackagedetailInfo.packagedetailid is number",
			Message: "【套餐明细ID】参数格式不正确",
		})
		return
	}
	request.packagedetailidInt, err = strconv.Atoi(request.packagedetailid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.GetGwPackagedetailInfo.packagedetailid parse err",
			Message: "packagedetailid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.packagedetailidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.GetGwPackagedetailInfo.packagedetailid value err",
			Message: "【套餐明细ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【套餐明细管理】信息
	gwPackagedetailInfo, err := service.GetGwPackagedetailInfo(request.packagedetailidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.GetGwPackagedetailInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if gwPackagedetailInfo.Packagedetailid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.GetGwPackagedetailInfo.not found",
			Message: "没有找到【套餐明细管理】信息",
		})
		return
	}
	if gwPackagedetailInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.GetGwPackagedetailInfo.has delete",
			Message: "【套餐明细管理】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【套餐明细管理】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"packagedetailid": gwPackagedetailInfo.Packagedetailid,                          //套餐明细ID
			"pkg_code":        gwPackagedetailInfo.PkgCode,                                  //套餐号
			"pkgdetail_name":  gwPackagedetailInfo.PkgdetailName,                            //商品名称
			"pkgdetail_nums":  gwPackagedetailInfo.PkgdetailNums,                            //数量
			"pkgdetail_price": gwPackagedetailInfo.PkgdetailPrice,                           //单价（元）
			"adder":           gwPackagedetailInfo.Adder,                                    //创建人
			"addtime":         gwPackagedetailInfo.Addtime.Format(golibs.Time_TIMEStandard), //创建时间
			"moder":           gwPackagedetailInfo.Moder,                                    //修改人
			"modtime":         gwPackagedetailInfo.Modtime.Format(golibs.Time_TIMEStandard), //修改时间
			"deleteStatus":    gwPackagedetailInfo.DeleteStatus,                             //0:未知，1：未删除，2：已删除
		},
	})
	//endregion
}

package gwPackage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【套餐管理】信息接口

// 请求
type getGwPackageInfoRequest struct {

	// 套餐ID
	packageid    string
	packageidInt int
}

// 方法
func GetGwPackageInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwPackage.GetGwPackageInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getGwPackageInfoRequest
	//endregion

	//region 验证packageid参数
	request.packageid = c.Param("id")
	if golibs.Length(request.packageid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.GetGwPackageInfo.packageid is null",
			Message: "缺少【套餐ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.packageid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.GetGwPackageInfo.packageid is number",
			Message: "【套餐ID】参数格式不正确",
		})
		return
	}
	request.packageidInt, err = strconv.Atoi(request.packageid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.GetGwPackageInfo.packageid parse err",
			Message: "packageid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.packageidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.GetGwPackageInfo.packageid value err",
			Message: "【套餐ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【套餐管理】信息
	gwPackageInfo, err := service.GetGwPackageInfo(request.packageidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.GetGwPackageInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if gwPackageInfo.Packageid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.GetGwPackageInfo.not found",
			Message: "没有找到【套餐管理】信息",
		})
		return
	}
	if gwPackageInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.GetGwPackageInfo.has delete",
			Message: "【套餐管理】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【套餐管理】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"packageid":       gwPackageInfo.Packageid,                                      //套餐ID
			"shop_code":       gwPackageInfo.ShopCode,                                       //店铺号
			"pkg_code":        gwPackageInfo.PkgCode,                                        //套餐号
			"pkg_type":        gwPackageInfo.PkgType,                                        //套餐类别
			"pkg_title":       gwPackageInfo.PkgTitle,                                       //套餐标题
			"pkg_people":      gwPackageInfo.PkgPeople,                                      //人数
			"pkg_isorder":     gwPackageInfo.PkgIsorder,                                     //是否需要预约
			"pkg_isrefund":    gwPackageInfo.PkgIsrefund,                                    //是否支持退款
			"pkg_isinhouse":   gwPackageInfo.PkgIsinhouse,                                   //是否仅限堂食
			"pkg_isnew":       gwPackageInfo.PkgIsnew,                                       //是否新品
			"pkg_isrecommend": gwPackageInfo.PkgIsrecommend,                                 //是否强烈推荐
			"pkg_validdays":   gwPackageInfo.PkgValiddays,                                   //有效期（天）
			"pkg_follows":     gwPackageInfo.PkgFollows,                                     //今日关注人数
			"pkg_createtime":  gwPackageInfo.PkgCreatetime.Format(golibs.Time_TIMEStandard), //上架日期
			"pkg_validtime":   gwPackageInfo.PkgValidtime.Format(golibs.Time_TIMEStandard),  //生效日期
			"pkg_exittime":    gwPackageInfo.PkgExittime.Format(golibs.Time_TIMEStandard),   //下架日期
			"pkg_links":       gwPackageInfo.PkgLinks,                                       //更多优惠链接地址
			"pkg_info":        gwPackageInfo.PkgInfo,                                        //套餐说明
			"adder":           gwPackageInfo.Adder,                                          //创建人
			"addtime":         gwPackageInfo.Addtime.Format(golibs.Time_TIMEStandard),       //创建时间
			"moder":           gwPackageInfo.Moder,                                          //修改人
			"modtime":         gwPackageInfo.Modtime.Format(golibs.Time_TIMEStandard),       //修改时间
			"deleteStatus":    gwPackageInfo.DeleteStatus,                                   //0:未知，1：未删除，2：已删除
		},
	})
	//endregion
}

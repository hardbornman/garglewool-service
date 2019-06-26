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

// 获取【套餐管理】列表接口

// 请求
type getGwPackageListRequest struct {

	// 当前页码
	pageIndex    string
	pageIndexInt int

	// 每页记录数
	pageSize    string
	pageSizeInt int
}

// 方法
func GetGwPackageList(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwPackage.GetGwPackageList.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getGwPackageListRequest
	//endregion

	//region 验证【当前页码】参数
	request.pageIndexInt = 1
	request.pageIndex = c.DefaultQuery("pageIndex", "")
	if golibs.Length(request.pageIndex) > 0 {
		if !golibs.IsNumber(request.pageIndex) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwPackage.GetGwPackageList.pageIndex is number",
				Message: "pageIndex参数格式不正确",
			})
			return
		}
		request.pageIndexInt, err = strconv.Atoi(request.pageIndex)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwPackage.GetGwPackageList.pageIndex parse err",
				Message: "pageIndex参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageIndexInt < 1 || request.pageIndexInt > 100000 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwPackage.GetGwPackageList.pageIndex value err",
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
				Code:    "gwPackage.GetGwPackageList.pageSize is number",
				Message: "pageSize参数格式不正确",
			})
			return
		}
		request.pageSizeInt, err = strconv.Atoi(request.pageSize)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwPackage.GetGwPackageList.pageSize parse err",
				Message: "pageSize参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageSizeInt < 1 || request.pageSizeInt > 500 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwPackage.GetGwPackageList.pageSize value err",
				Message: "pageSize参数值错误",
			})
			return
		}
	}
	//endregion

	//region 查询【套餐管理】列表
	list, total, err := service.GetGwPackageList(request.pageIndexInt, request.pageSizeInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.GetGwPackageList.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 组装【套餐管理】列表
	gwPackagesArray := make([]map[string]interface{}, len(list))
	if len(list) > 0 {
		for i, v := range list {
			gwPackagesArray[i] = map[string]interface{}{
				"packageid":       v.Packageid,                                      //套餐ID
				"shop_code":       v.ShopCode,                                       //店铺号
				"pkg_code":        v.PkgCode,                                        //套餐号
				"pkg_type":        v.PkgType,                                        //套餐类别
				"pkg_title":       v.PkgTitle,                                       //套餐标题
				"pkg_people":      v.PkgPeople,                                      //人数
				"pkg_isorder":     v.PkgIsorder,                                     //是否需要预约
				"pkg_isrefund":    v.PkgIsrefund,                                    //是否支持退款
				"pkg_isinhouse":   v.PkgIsinhouse,                                   //是否仅限堂食
				"pkg_isnew":       v.PkgIsnew,                                       //是否新品
				"pkg_isrecommend": v.PkgIsrecommend,                                 //是否强烈推荐
				"pkg_validdays":   v.PkgValiddays,                                   //有效期（天）
				"pkg_follows":     v.PkgFollows,                                     //今日关注人数
				"pkg_createtime":  v.PkgCreatetime.Format(golibs.Time_TIMEStandard), //上架日期
				"pkg_validtime":   v.PkgValidtime.Format(golibs.Time_TIMEStandard),  //生效日期
				"pkg_exittime":    v.PkgExittime.Format(golibs.Time_TIMEStandard),   //下架日期
				"pkg_links":       v.PkgLinks,                                       //更多优惠链接地址
				"pkg_info":        v.PkgInfo,                                        //套餐说明
				"adder":           v.Adder,                                          //创建人
				"addtime":         v.Addtime.Format(golibs.Time_TIMEStandard),       //创建时间
				"moder":           v.Moder,                                          //修改人
				"modtime":         v.Modtime.Format(golibs.Time_TIMEStandard),       //修改时间
				"deleteStatus":    v.DeleteStatus,                                   //0:未知，1：未删除，2：已删除
			}
		}
	}
	//endregion

	//region 返回【套餐管理】列表
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"total": total,
			"list":  gwPackagesArray,
		},
	})
	//endregion
}

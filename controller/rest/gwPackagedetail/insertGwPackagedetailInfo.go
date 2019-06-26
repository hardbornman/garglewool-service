package gwPackagedetail

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"time"
)

// 插入【套餐明细管理】信息接口

// 请求
type insertGwPackagedetailInfoRequest struct {

	// 套餐号
	PkgCode string `form:"pkg_code"`

	// 商品名称
	PkgdetailName string `form:"pkgdetail_name"`

	// 数量
	PkgdetailNums int `form:"pkgdetail_nums"`

	// 单价（元）
	PkgdetailPrice float64 `form:"pkgdetail_price"`

	// 创建人
	Adder int `form:"adder"`

	// 创建时间
	Addtime     string `form:"addtime"`
	addtimeTime time.Time

	// 修改人
	Moder int `form:"moder"`

	// 修改时间
	Modtime     string `form:"modtime"`
	modtimeTime time.Time
}

// 方法
func InsertGwPackagedetailInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwPackagedetail.InsertGwPackagedetailInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request insertGwPackagedetailInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证pkg_code参数
	if golibs.Length(request.PkgCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.pkg_code is null",
			Message: "缺少【套餐号】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.PkgCode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.pkg_code format err",
			Message: "【套餐号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.PkgCode) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.pkg_code length err",
			Message: "【套餐号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证pkgdetail_name参数
	if golibs.Length(request.PkgdetailName) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.pkgdetail_name is null",
			Message: "缺少【商品名称】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.PkgdetailName) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.pkgdetail_name format err",
			Message: "【商品名称】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.PkgdetailName) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.pkgdetail_name length err",
			Message: "【商品名称】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证pkgdetail_nums参数
	if request.PkgdetailNums <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.pkgdetail_nums value err",
			Message: "pkgdetail_nums参数值错误",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.adder value err",
			Message: "adder参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.moder value err",
			Message: "moder参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 插入【套餐明细管理】信息
	packagedetailid, err := service.InsertGwPackagedetailInfo(request.PkgCode, request.PkgdetailName, request.PkgdetailNums, request.PkgdetailPrice, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.insert err",
			Message: "插入出错:" + err.Error(),
		})
		return
	}
	if packagedetailid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackagedetail.InsertGwPackagedetailInfo.insert failure",
			Message: "插入失败",
		})
		return
	}
	//endregion

	//region 返回插入【套餐明细管理】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result":          true,
			"packagedetailid": packagedetailid,
		},
	})
	//endregion
}

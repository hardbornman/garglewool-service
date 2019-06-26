package gwPackage

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

// 修改【套餐管理】信息接口

// 请求
type updateGwPackageInfoRequest struct {

	// 套餐ID
	packageid    string
	packageidInt int

	// 店铺号
	ShopCode string `form:"shop_code"`

	// 套餐号
	PkgCode string `form:"pkg_code"`

	// 套餐类别
	PkgType string `form:"pkg_type"`

	// 套餐标题
	PkgTitle string `form:"pkg_title"`

	// 人数
	PkgPeople int `form:"pkg_people"`

	// 是否需要预约
	PkgIsorder int `form:"pkg_isorder"`

	// 是否支持退款
	PkgIsrefund int `form:"pkg_isrefund"`

	// 是否仅限堂食
	PkgIsinhouse int `form:"pkg_isinhouse"`

	// 是否新品
	PkgIsnew int `form:"pkg_isnew"`

	// 是否强烈推荐
	PkgIsrecommend int `form:"pkg_isrecommend"`

	// 有效期（天）
	PkgValiddays int `form:"pkg_validdays"`

	// 今日关注人数
	PkgFollows int `form:"pkg_follows"`

	// 上架日期
	PkgCreatetime      string `form:"pkg_createtime"`
	pkg_createtimeTime time.Time

	// 生效日期
	PkgValidtime      string `form:"pkg_validtime"`
	pkg_validtimeTime time.Time

	// 下架日期
	PkgExittime      string `form:"pkg_exittime"`
	pkg_exittimeTime time.Time

	// 更多优惠链接地址
	PkgLinks string `form:"pkg_links"`

	// 套餐说明
	PkgInfo string `form:"pkg_info"`

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
func UpdateGwPackageInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwPackage.UpdateGwPackageInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateGwPackageInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证packageid参数
	request.packageid = c.Param("id")
	if golibs.Length(request.packageid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.packageid is null",
			Message: "缺少【套餐ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.packageid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.packageid is number",
			Message: "【套餐ID】参数格式不正确",
		})
		return
	}
	request.packageidInt, err = strconv.Atoi(request.packageid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.packageid parse err",
			Message: "【套餐ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.packageidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.packageid value err",
			Message: "【套餐ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证shop_code参数
	if golibs.Length(request.ShopCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.shop_code is null",
			Message: "缺少【店铺号】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.ShopCode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.shop_code format err",
			Message: "【店铺号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.ShopCode) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.shop_code length err",
			Message: "【店铺号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证pkg_code参数
	if golibs.Length(request.PkgCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_code is null",
			Message: "缺少【套餐号】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.PkgCode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_code format err",
			Message: "【套餐号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.PkgCode) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_code length err",
			Message: "【套餐号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证pkg_type参数
	if golibs.Length(request.PkgType) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_type is null",
			Message: "缺少【套餐类别】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.PkgType) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_type format err",
			Message: "【套餐类别】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.PkgType) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_type length err",
			Message: "【套餐类别】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证pkg_title参数
	if golibs.Length(request.PkgTitle) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_title is null",
			Message: "缺少【套餐标题】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.PkgTitle) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_title format err",
			Message: "【套餐标题】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.PkgTitle) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_title length err",
			Message: "【套餐标题】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证pkg_people参数
	if request.PkgPeople <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_people value err",
			Message: "【人数】参数值错误",
		})
		return
	}
	//endregion

	//region 验证pkg_isorder参数
	if request.PkgIsorder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_isorder value err",
			Message: "【是否需要预约】参数值错误",
		})
		return
	}
	//endregion

	//region 验证pkg_isrefund参数
	if request.PkgIsrefund <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_isrefund value err",
			Message: "【是否支持退款】参数值错误",
		})
		return
	}
	//endregion

	//region 验证pkg_isinhouse参数
	if request.PkgIsinhouse <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_isinhouse value err",
			Message: "【是否仅限堂食】参数值错误",
		})
		return
	}
	//endregion

	//region 验证pkg_isnew参数
	if request.PkgIsnew <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_isnew value err",
			Message: "【是否新品】参数值错误",
		})
		return
	}
	//endregion

	//region 验证pkg_isrecommend参数
	if request.PkgIsrecommend <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_isrecommend value err",
			Message: "【是否强烈推荐】参数值错误",
		})
		return
	}
	//endregion

	//region 验证pkg_validdays参数
	if request.PkgValiddays <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_validdays value err",
			Message: "【有效期（天）】参数值错误",
		})
		return
	}
	//endregion

	//region 验证pkg_follows参数
	if request.PkgFollows <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_follows value err",
			Message: "【今日关注人数】参数值错误",
		})
		return
	}
	//endregion

	//region 验证pkg_createtime参数
	if golibs.Length(request.PkgCreatetime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_createtime length err",
			Message: "缺少【上架日期】参数",
		})
		return
	}
	request.pkg_createtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.PkgCreatetime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_createtime parse err",
			Message: "【上架日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.pkg_createtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_createtime value err",
			Message: "【上架日期】参数值错误:" + request.PkgCreatetime,
		})
		return
	}
	//endregion

	//region 验证pkg_validtime参数
	if golibs.Length(request.PkgValidtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_validtime length err",
			Message: "缺少【生效日期】参数",
		})
		return
	}
	request.pkg_validtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.PkgValidtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_validtime parse err",
			Message: "【生效日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.pkg_validtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_validtime value err",
			Message: "【生效日期】参数值错误:" + request.PkgValidtime,
		})
		return
	}
	//endregion

	//region 验证pkg_exittime参数
	if golibs.Length(request.PkgExittime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_exittime length err",
			Message: "缺少【下架日期】参数",
		})
		return
	}
	request.pkg_exittimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.PkgExittime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_exittime parse err",
			Message: "【下架日期】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.pkg_exittimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_exittime value err",
			Message: "【下架日期】参数值错误:" + request.PkgExittime,
		})
		return
	}
	//endregion

	//region 验证pkg_links参数
	if golibs.Length(request.PkgLinks) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_links is null",
			Message: "缺少【更多优惠链接地址】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.PkgLinks) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_links format err",
			Message: "【更多优惠链接地址】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.PkgLinks) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_links length err",
			Message: "【更多优惠链接地址】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证pkg_info参数
	if golibs.Length(request.PkgInfo) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_info is null",
			Message: "缺少【套餐说明】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.PkgInfo) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_info format err",
			Message: "【套餐说明】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.PkgInfo) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.pkg_info length err",
			Message: "【套餐说明】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.adder value err",
			Message: "【创建人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.moder value err",
			Message: "【修改人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 修改【套餐管理】信息
	isSuccess, err := service.UpdateGwPackageInfo(request.packageidInt, request.ShopCode, request.PkgCode, request.PkgType, request.PkgTitle, request.PkgPeople, int8(request.PkgIsorder), int8(request.PkgIsrefund), int8(request.PkgIsinhouse), int8(request.PkgIsnew), int8(request.PkgIsrecommend), request.PkgValiddays, request.PkgFollows, request.pkg_createtimeTime, request.pkg_validtimeTime, request.pkg_exittimeTime, request.PkgLinks, request.PkgInfo, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwPackage.UpdateGwPackageInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【套餐管理】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

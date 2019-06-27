package comment

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

// 修改【评论表】信息接口

// 请求
type updateCommentInfoRequest struct {

	// 评论ID
	id    string
	idInt int

	// 订单号
	Ordercode string `form:"ordercode"`

	// 用户id
	Commentor int `form:"commentor"`

	// 评价星级
	Star int `form:"star"`

	// 评价内容
	Info string `form:"info"`

	// 评价图片
	Images string `form:"images"`

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

	// 店铺ID
	Shopid int `form:"shopid"`
}

// 方法
func UpdateCommentInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "comment.UpdateCommentInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateCommentInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证id参数
	request.id = c.Param("id")
	if golibs.Length(request.id) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.id is null",
			Message: "缺少【评论ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.id) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.id is number",
			Message: "【评论ID】参数格式不正确",
		})
		return
	}
	request.idInt, err = strconv.Atoi(request.id)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.id parse err",
			Message: "【评论ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.idInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.id value err",
			Message: "【评论ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证ordercode参数
	if golibs.Length(request.Ordercode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.ordercode is null",
			Message: "缺少【订单号】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.Ordercode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.ordercode format err",
			Message: "【订单号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Ordercode) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.ordercode length err",
			Message: "【订单号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证commentor参数
	if request.Commentor <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.commentor value err",
			Message: "【用户id】参数值错误",
		})
		return
	}
	//endregion

	//region 验证star参数
	if request.Star <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.star value err",
			Message: "【评价星级】参数值错误",
		})
		return
	}
	//endregion

	//region 验证info参数
	if golibs.Length(request.Info) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.info is null",
			Message: "缺少【评价内容】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Info) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.info format err",
			Message: "【评价内容】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Info) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.info length err",
			Message: "【评价内容】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证images参数
	if golibs.Length(request.Images) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.images is null",
			Message: "缺少【评价图片】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Images) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.images format err",
			Message: "【评价图片】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Images) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.images length err",
			Message: "【评价图片】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.adder value err",
			Message: "【创建人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.moder value err",
			Message: "【修改人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 验证shopid参数
	if request.Shopid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.shopid value err",
			Message: "【店铺ID】参数值错误",
		})
		return
	}
	//endregion

	//region 修改【评论表】信息
	isSuccess, err := service.UpdateCommentInfo(request.idInt, request.Ordercode, request.Commentor, request.Star, request.Info, request.Images, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime, request.Shopid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.UpdateCommentInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【评论表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

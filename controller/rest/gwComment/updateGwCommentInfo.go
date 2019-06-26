package gwComment

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

// 修改【评论管理】信息接口

// 请求
type updateGwCommentInfoRequest struct {

	// 评论ID
	commentid    string
	commentidInt int

	// 订单号
	OrderCode string `form:"order_code"`

	// 用户id
	OrderCommentor int `form:"order_commentor"`

	// 评价星级
	CommentStar int `form:"comment_star"`

	// 评价内容
	CommentInfo string `form:"comment_info"`

	// 评价图片
	CommentImages string `form:"comment_images"`

	// 评价时间
	CommentCreatetime      string `form:"comment_createtime"`
	comment_createtimeTime time.Time

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
func UpdateGwCommentInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwComment.UpdateGwCommentInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateGwCommentInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证commentid参数
	request.commentid = c.Param("id")
	if golibs.Length(request.commentid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.commentid is null",
			Message: "缺少【评论ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.commentid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.commentid is number",
			Message: "【评论ID】参数格式不正确",
		})
		return
	}
	request.commentidInt, err = strconv.Atoi(request.commentid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.commentid parse err",
			Message: "【评论ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.commentidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.commentid value err",
			Message: "【评论ID】参数值错误",
		})
		return
	}
	//endregion

	//region 验证order_code参数
	if golibs.Length(request.OrderCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.order_code is null",
			Message: "缺少【订单号】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.OrderCode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.order_code format err",
			Message: "【订单号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.OrderCode) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.order_code length err",
			Message: "【订单号】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证order_commentor参数
	if request.OrderCommentor <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.order_commentor value err",
			Message: "【用户id】参数值错误",
		})
		return
	}
	//endregion

	//region 验证comment_star参数
	if request.CommentStar <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.comment_star value err",
			Message: "【评价星级】参数值错误",
		})
		return
	}
	//endregion

	//region 验证comment_info参数
	if golibs.Length(request.CommentInfo) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.comment_info is null",
			Message: "缺少【评价内容】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.CommentInfo) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.comment_info format err",
			Message: "【评价内容】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.CommentInfo) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.comment_info length err",
			Message: "【评价内容】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证comment_images参数
	if golibs.Length(request.CommentImages) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.comment_images is null",
			Message: "缺少【评价图片】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.CommentImages) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.comment_images format err",
			Message: "【评价图片】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.CommentImages) > 255 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.comment_images length err",
			Message: "【评价图片】参数长度不能超过255个字符",
		})
		return
	}
	//endregion

	//region 验证comment_createtime参数
	if golibs.Length(request.CommentCreatetime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.comment_createtime length err",
			Message: "缺少【评价时间】参数",
		})
		return
	}
	request.comment_createtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.CommentCreatetime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.comment_createtime parse err",
			Message: "【评价时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.comment_createtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.comment_createtime value err",
			Message: "【评价时间】参数值错误:" + request.CommentCreatetime,
		})
		return
	}
	//endregion

	//region 验证adder参数
	if request.Adder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.adder value err",
			Message: "【创建人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证addtime参数
	if golibs.Length(request.Addtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.addtime length err",
			Message: "缺少【创建时间】参数",
		})
		return
	}
	request.addtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Addtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.addtime parse err",
			Message: "【创建时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.addtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.addtime value err",
			Message: "【创建时间】参数值错误:" + request.Addtime,
		})
		return
	}
	//endregion

	//region 验证moder参数
	if request.Moder <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.moder value err",
			Message: "【修改人】参数值错误",
		})
		return
	}
	//endregion

	//region 验证modtime参数
	if golibs.Length(request.Modtime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.modtime length err",
			Message: "缺少【修改时间】参数",
		})
		return
	}
	request.modtimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.Modtime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.modtime parse err",
			Message: "【修改时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.modtimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.modtime value err",
			Message: "【修改时间】参数值错误:" + request.Modtime,
		})
		return
	}
	//endregion

	//region 修改【评论管理】信息
	isSuccess, err := service.UpdateGwCommentInfo(request.commentidInt, request.OrderCode, request.OrderCommentor, request.CommentStar, request.CommentInfo, request.CommentImages, request.comment_createtimeTime, request.Adder, request.addtimeTime, request.Moder, request.modtimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.UpdateGwCommentInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【评论管理】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

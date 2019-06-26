package gwComment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 获取【评论管理】信息接口

// 请求
type getGwCommentInfoRequest struct {

	// 评论ID
	commentid    string
	commentidInt int
}

// 方法
func GetGwCommentInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwComment.GetGwCommentInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getGwCommentInfoRequest
	//endregion

	//region 验证commentid参数
	request.commentid = c.Param("id")
	if golibs.Length(request.commentid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.GetGwCommentInfo.commentid is null",
			Message: "缺少【评论ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.commentid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.GetGwCommentInfo.commentid is number",
			Message: "【评论ID】参数格式不正确",
		})
		return
	}
	request.commentidInt, err = strconv.Atoi(request.commentid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.GetGwCommentInfo.commentid parse err",
			Message: "commentid参数解析出错:" + err.Error(),
		})
		return
	}
	if request.commentidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.GetGwCommentInfo.commentid value err",
			Message: "【评论ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【评论管理】信息
	gwCommentInfo, err := service.GetGwCommentInfo(request.commentidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.GetGwCommentInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if gwCommentInfo.Commentid <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.GetGwCommentInfo.not found",
			Message: "没有找到【评论管理】信息",
		})
		return
	}
	if gwCommentInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.GetGwCommentInfo.has delete",
			Message: "【评论管理】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【评论管理】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"commentid":          gwCommentInfo.Commentid,                                          //评论ID
			"order_code":         gwCommentInfo.OrderCode,                                          //订单号
			"order_commentor":    gwCommentInfo.OrderCommentor,                                     //用户id
			"comment_star":       gwCommentInfo.CommentStar,                                        //评价星级
			"comment_info":       gwCommentInfo.CommentInfo,                                        //评价内容
			"comment_images":     gwCommentInfo.CommentImages,                                      //评价图片
			"comment_createtime": gwCommentInfo.CommentCreatetime.Format(golibs.Time_TIMEStandard), //评价时间
			"adder":              gwCommentInfo.Adder,                                              //创建人
			"addtime":            gwCommentInfo.Addtime.Format(golibs.Time_TIMEStandard),           //创建时间
			"moder":              gwCommentInfo.Moder,                                              //修改人
			"modtime":            gwCommentInfo.Modtime.Format(golibs.Time_TIMEStandard),           //修改时间
			"deleteStatus":       gwCommentInfo.DeleteStatus,                                       //0:未知，1：未删除，2：已删除
		},
	})
	//endregion
}

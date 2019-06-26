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

// 删除【评论管理】信息接口

// 请求
type deleteGwCommentInfoRequest struct {

	// 评论ID
	commentid    string
	commentidInt int
}

// 方法
func DeleteGwCommentInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwComment.DeleteGwCommentInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteGwCommentInfoRequest
	//endregion

	//region 验证commentid参数
	request.commentid = c.Param("id")
	if golibs.Length(request.commentid) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.DeleteGwCommentInfo.commentid is null",
			Message: "缺少【评论ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.commentid) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.DeleteGwCommentInfo.commentid is number",
			Message: "【评论ID】参数格式不正确",
		})
		return
	}
	request.commentidInt, err = strconv.Atoi(request.commentid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.DeleteGwCommentInfo.commentid parse err",
			Message: "【评论ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.commentidInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.DeleteGwCommentInfo.commentid value err",
			Message: "【评论ID】参数值错误",
		})
		return
	}
	//endregion

	//region 删除【评论管理】信息
	isSuccess, err := service.DeleteGwCommentInfo(request.commentidInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.DeleteGwCommentInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.DeleteGwCommentInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【评论管理】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

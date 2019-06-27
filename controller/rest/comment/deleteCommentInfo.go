package comment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/hardbornman/garglewool-service/service"
	"net/http"
	"strconv"
)

// 删除【评论表】信息接口

// 请求
type deleteCommentInfoRequest struct {

	// 评论ID
	id    string
	idInt int
}

// 方法
func DeleteCommentInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "comment.DeleteCommentInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteCommentInfoRequest
	//endregion

	//region 验证id参数
	request.id = c.Param("id")
	if golibs.Length(request.id) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.DeleteCommentInfo.id is null",
			Message: "缺少【评论ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.id) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.DeleteCommentInfo.id is number",
			Message: "【评论ID】参数格式不正确",
		})
		return
	}
	request.idInt, err = strconv.Atoi(request.id)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.DeleteCommentInfo.id parse err",
			Message: "【评论ID】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.idInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.DeleteCommentInfo.id value err",
			Message: "【评论ID】参数值错误",
		})
		return
	}
	//endregion

	//region 删除【评论表】信息
	isSuccess, err := service.DeleteCommentInfo(request.idInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.DeleteCommentInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.DeleteCommentInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【评论表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}

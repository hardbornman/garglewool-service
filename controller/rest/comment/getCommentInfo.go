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

// 获取【评论表】信息接口

// 请求
type getCommentInfoRequest struct {

	// 评论ID
	id    string
	idInt int
}

// 方法
func GetCommentInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "comment.GetCommentInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getCommentInfoRequest
	//endregion

	//region 验证id参数
	request.id = c.Param("id")
	if golibs.Length(request.id) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.GetCommentInfo.id is null",
			Message: "缺少【评论ID】参数",
		})
		return
	}
	if !golibs.IsNumber(request.id) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.GetCommentInfo.id is number",
			Message: "【评论ID】参数格式不正确",
		})
		return
	}
	request.idInt, err = strconv.Atoi(request.id)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.GetCommentInfo.id parse err",
			Message: "id参数解析出错:" + err.Error(),
		})
		return
	}
	if request.idInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.GetCommentInfo.id value err",
			Message: "【评论ID】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【评论表】信息
	commentInfo, err := service.GetCommentInfo(request.idInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.GetCommentInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if commentInfo.Id <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.GetCommentInfo.not found",
			Message: "没有找到【评论表】信息",
		})
		return
	}
	if commentInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "comment.GetCommentInfo.has delete",
			Message: "【评论表】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【评论表】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"id":           commentInfo.Id,                                       //评论ID
			"ordercode":    commentInfo.Ordercode,                                //订单号
			"commentor":    commentInfo.Commentor,                                //用户id
			"star":         commentInfo.Star,                                     //评价星级
			"info":         commentInfo.Info,                                     //评价内容
			"images":       commentInfo.Images,                                   //评价图片
			"adder":        commentInfo.Adder,                                    //创建人
			"addtime":      commentInfo.Addtime.Format(golibs.Time_TIMEStandard), //创建时间
			"moder":        commentInfo.Moder,                                    //修改人
			"modtime":      commentInfo.Modtime.Format(golibs.Time_TIMEStandard), //修改时间
			"deleteStatus": commentInfo.DeleteStatus,                             //0:未知，1：未删除，2：已删除
			"shopid":       commentInfo.Shopid,                                   //店铺ID
		},
	})
	//endregion
}

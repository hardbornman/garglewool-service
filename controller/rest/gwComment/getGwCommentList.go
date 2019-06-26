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

// 获取【评论管理】列表接口

// 请求
type getGwCommentListRequest struct {

	// 当前页码
	pageIndex    string
	pageIndexInt int

	// 每页记录数
	pageSize    string
	pageSizeInt int
}

// 方法
func GetGwCommentList(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwComment.GetGwCommentList.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getGwCommentListRequest
	//endregion

	//region 验证【当前页码】参数
	request.pageIndexInt = 1
	request.pageIndex = c.DefaultQuery("pageIndex", "")
	if golibs.Length(request.pageIndex) > 0 {
		if !golibs.IsNumber(request.pageIndex) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwComment.GetGwCommentList.pageIndex is number",
				Message: "pageIndex参数格式不正确",
			})
			return
		}
		request.pageIndexInt, err = strconv.Atoi(request.pageIndex)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwComment.GetGwCommentList.pageIndex parse err",
				Message: "pageIndex参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageIndexInt < 1 || request.pageIndexInt > 100000 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwComment.GetGwCommentList.pageIndex value err",
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
				Code:    "gwComment.GetGwCommentList.pageSize is number",
				Message: "pageSize参数格式不正确",
			})
			return
		}
		request.pageSizeInt, err = strconv.Atoi(request.pageSize)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwComment.GetGwCommentList.pageSize parse err",
				Message: "pageSize参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageSizeInt < 1 || request.pageSizeInt > 500 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "gwComment.GetGwCommentList.pageSize value err",
				Message: "pageSize参数值错误",
			})
			return
		}
	}
	//endregion

	//region 查询【评论管理】列表
	list, total, err := service.GetGwCommentList(request.pageIndexInt, request.pageSizeInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "gwComment.GetGwCommentList.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 组装【评论管理】列表
	gwCommentsArray := make([]map[string]interface{}, len(list))
	if len(list) > 0 {
		for i, v := range list {
			gwCommentsArray[i] = map[string]interface{}{
				"commentid":          v.Commentid,                                          //评论ID
				"order_code":         v.OrderCode,                                          //订单号
				"order_commentor":    v.OrderCommentor,                                     //用户id
				"comment_star":       v.CommentStar,                                        //评价星级
				"comment_info":       v.CommentInfo,                                        //评价内容
				"comment_images":     v.CommentImages,                                      //评价图片
				"comment_createtime": v.CommentCreatetime.Format(golibs.Time_TIMEStandard), //评价时间
				"adder":              v.Adder,                                              //创建人
				"addtime":            v.Addtime.Format(golibs.Time_TIMEStandard),           //创建时间
				"moder":              v.Moder,                                              //修改人
				"modtime":            v.Modtime.Format(golibs.Time_TIMEStandard),           //修改时间
				"deleteStatus":       v.DeleteStatus,                                       //0:未知，1：未删除，2：已删除
			}
		}
	}
	//endregion

	//region 返回【评论管理】列表
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"total": total,
			"list":  gwCommentsArray,
		},
	})
	//endregion
}

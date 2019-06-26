package model

import (
	"time"
)

// 评论管理
type GwComment struct {
	Commentid         int       //评论ID
	OrderCode         string    //订单号
	OrderCommentor    int       //用户id
	CommentStar       int       //评价星级
	CommentInfo       string    //评价内容
	CommentImages     string    //评价图片
	CommentCreatetime time.Time //评价时间
	Adder             int       //创建人
	Addtime           time.Time //创建时间
	Moder             int       //修改人
	Modtime           time.Time //修改时间
	DeleteStatus      int8      //0:未知，1：未删除，2：已删除
}

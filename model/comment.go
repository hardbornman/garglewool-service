package model

import (
	"time"
)

// 评论表
type Comment struct {
	Id           int       //评论ID
	Ordercode    string    //订单号
	Commentor    int       //用户id
	Star         int       //评价星级
	Info         string    //评价内容
	Images       string    //评价图片
	Adder        int       //创建人
	Addtime      time.Time //创建时间
	Moder        int       //修改人
	Modtime      time.Time //修改时间
	DeleteStatus int8      //0:未知，1：未删除，2：已删除
	Shopid       int       //店铺ID
	Shopname     string    //店铺名称【外键表:店铺表】
}

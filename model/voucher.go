package model

import (
	"time"
)

// 抵用券管理
type Voucher struct {
	Voucherid    int       //抵用券ID
	Code         string    //抵用券号
	Guestid      int       //用户id
	Name         string    //用户名【外键表:买家客户表】
	Quota        float64   //额度
	Reduce       float64   //抵消
	Createtime   time.Time //创建时间
	Validdays    int       //有效期（天）
	Isinvalid    bool      //是否失效
	Adder        int       //创建人
	Addtime      time.Time //创建时间
	Moder        int       //修改人
	Modtime      time.Time //修改时间
	DeleteStatus int8      //0:未知，1：未删除，2：已删除
}

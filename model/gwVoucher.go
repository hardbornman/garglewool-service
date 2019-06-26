package model

import (
	"time"
)

// 抵用券管理
type GwVoucher struct {
	Voucherid         int       //抵用券ID
	VoucherCode       string    //抵用券号
	VoucherUserid     int       //用户id
	VoucherQuota      float64   //额度
	VoucherReduce     float64   //抵消
	VoucherCreatetime time.Time //创建时间
	VoucherValiddays  int       //有效期（天）
	VoucherIsinvalid  int8      //是否失效
	Adder             int       //创建人
	Addtime           time.Time //创建时间
	Moder             int       //修改人
	Modtime           time.Time //修改时间
	DeleteStatus      int8      //0:未知，1：未删除，2：已删除
}

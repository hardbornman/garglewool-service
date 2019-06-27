package model

import (
	"time"
)

// 订单表
type Order struct {
	Orderid      int       //订单ID
	Pkgcode      string    //套餐号
	Ordercode    string    //订单号
	Buyer        int       //用户id
	Paytype      int       //支付方式
	Totalprice   float64   //订单总价（元）
	Payprice     float64   //购买总价（元）
	Paytime      time.Time //购买日期
	Isinvalid    bool      //是否失效
	Isused       bool      //是否使用
	Isrefund     bool      //是否退款
	Refundprice  float64   //退款金额
	Refundtime   time.Time //退款日期
	Remark       string    //订单备注
	Adder        int       //创建人
	Addtime      time.Time //创建时间
	Moder        int       //修改人
	Modtime      time.Time //修改时间
	DeleteStatus int8      //0:未知，1：未删除，2：已删除
}

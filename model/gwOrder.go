package model

import (
	"time"
)

// 订单管理
type GwOrder struct {
	Orderid          int       //订单ID
	PkgCode          string    //套餐号
	OrderCode        string    //订单号
	OrderBuyer       int       //用户id
	OrderPaytype     int       //支付方式
	OrderTotalprice  float64   //订单总价（元）
	OrderPayprice    float64   //购买总价（元）
	OrderPaytime     time.Time //购买日期
	OrderIsinvalid   int8      //是否失效
	OrderIsused      int8      //是否使用
	OrderIsrefund    int8      //是否退款
	OrderRefundprice float64   //退款金额
	OrderRefundtime  time.Time //退款日期
	OrderRemark      string    //订单备注
	Adder            int       //创建人
	Addtime          time.Time //创建时间
	Moder            int       //修改人
	Modtime          time.Time //修改时间
	DeleteStatus     int8      //0:未知，1：未删除，2：已删除
}

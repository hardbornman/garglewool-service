package model

import (
	"time"
)

// 店铺表
type Shop struct {
	Shopid       int       //店铺ID
	Shopcode     string    //店铺号
	Shopname     string    //店铺名称
	Province     string    //店铺地址-省
	City         string    //店铺地址-市
	District     string    //店铺地址-区
	Address      string    //店铺详细地址
	Phone        string    //电话
	Leaguetime   time.Time //加盟平台日期
	Exittime     time.Time //退出平台日期
	Adder        int       //创建人
	Addtime      time.Time //创建时间
	Moder        int       //修改人
	Modtime      time.Time //修改时间
	DeleteStatus int8      //0:未知，1：未删除，2：已删除
	Merchantid   int       //商家ID
	Merchantname string    //商家名【外键表:商家用户表】
	Longtitude   string    //经度
	Latitude     string    //纬度
}

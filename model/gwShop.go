package model

import (
	"time"
)

// 店铺管理
type GwShop struct {
	Shopid         int       //店铺ID
	ShopCode       string    //店铺号
	ShopName       string    //店铺名称
	ShopProvince   string    //店铺地址-省
	ShopCity       string    //店铺地址-市
	ShopDistrict   string    //店铺地址-区
	ShopAddress    string    //店铺详细地址
	ShopPhone      string    //加盟平台日期
	ShopCreatetime time.Time //加盟平台日期
	ShopExittime   time.Time //退出平台日期
	Adder          int       //创建人
	Addtime        time.Time //创建时间
	Moder          int       //修改人
	Modtime        time.Time //修改时间
	DeleteStatus   int8      //0:未知，1：未删除，2：已删除
}

package model

import (
	"time"
)

// 买家客户表
type Guest struct {
	Guestid      int       //用户ID
	Name         string    //用户名
	Password     string    //用户密码
	Phone        string    //手机号
	Golds        int       //用户金币数量
	Adder        int       //创建人
	Addtime      time.Time //创建时间
	Moder        int       //修改人
	Modtime      time.Time //修改时间
	DeleteStatus int8      //0:未知，1：未删除，2：已删除
}

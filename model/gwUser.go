package model

import (
	"time"
)

// 用户管理
type GwUser struct {
	Userid       int       //用户ID
	UserName     string    //用户名
	UserPhone    string    //手机号
	UserGolds    int       //用户金币数量
	Adder        int       //创建人
	Addtime      time.Time //创建时间
	Moder        int       //修改人
	Modtime      time.Time //修改时间
	DeleteStatus int8      //0:未知，1：未删除，2：已删除
}

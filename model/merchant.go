package model

import (
	"time"
)

// 商家用户表
type Merchant struct {
	Merchantid    int       //商家ID
	Merchantname  string    //商家名
	Phone         string    //手机号
	Createtime    time.Time //创建时间
	Updatetime    time.Time //修改时间
	DeleteStatus  int8      //0:未知，1：未删除，2：已删除
	Userroleid    int       //用户角色ID
	Rolename      string    //角色名称【外键表:用户角色表】
	Loginpwd      string    //登录密码
	Loginaccount  string    //登录账户
	Nick          string    //昵称
	Wechataccount string    //微信账号
	Wechatsign    string    //微信签名
	Remark        string    //备注
	Lastlogintime time.Time //最近登录时间
	Lastloginaddr string    //最近登录地址
	Rigstertime   time.Time //注册时间
	Enable        bool      //是否启用
	Addr          string    //地址
}

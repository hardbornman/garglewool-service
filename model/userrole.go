package model

import (
	"time"
)

// 用户角色表
type Userrole struct {
	Userroleid   int       //用户角色ID
	Rolename     string    //角色名称
	Desc         string    //描述
	Remark       string    //备注
	DeleteStatus int8      //0:未知，1：未删除，2：已删除
	Createtime   time.Time //创建时间
	Updatetime   time.Time //更新时间
}

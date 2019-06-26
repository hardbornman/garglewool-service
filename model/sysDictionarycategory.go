package model

import (
	"time"
)

// 字典分类
type SysDictionarycategory struct {
	Dictionarycategoryid int       //字典分类ID
	Categorykey          string    //字典类别key
	Categoryvalue        string    //字典类别value
	Adder                int       //创建人
	Addtime              time.Time //创建时间
	Moder                int       //修改人
	Modtime              time.Time //修改时间
	DeleteStatus         int8      //0:未知，1：未删除，2：已删除
}

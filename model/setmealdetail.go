package model

import (
	"time"
)

// 套餐明细表
type Setmealdetail struct {
	Setmealdetailid int       //套餐明细ID
	Setmealid       int       //套餐ID
	Name            string    //商品名称
	Nums            int       //数量
	Price           float64   //单价（元）
	Adder           int       //创建人
	Addtime         time.Time //创建时间
	Moder           int       //修改人
	Modtime         time.Time //修改时间
	DeleteStatus    int8      //0:未知，1：未删除，2：已删除
}

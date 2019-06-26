package model

import (
	"time"
)

// 套餐明细管理
type GwPackagedetail struct {
	Packagedetailid int       //套餐明细ID
	PkgCode         string    //套餐号
	PkgdetailName   string    //商品名称
	PkgdetailNums   int       //数量
	PkgdetailPrice  float64   //单价（元）
	Adder           int       //创建人
	Addtime         time.Time //创建时间
	Moder           int       //修改人
	Modtime         time.Time //修改时间
	DeleteStatus    int8      //0:未知，1：未删除，2：已删除
}

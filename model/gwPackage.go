package model

import (
	"time"
)

// 套餐管理
type GwPackage struct {
	Packageid      int       //套餐ID
	ShopCode       string    //店铺号
	PkgCode        string    //套餐号
	PkgType        string    //套餐类别
	PkgTitle       string    //套餐标题
	PkgPeople      int       //人数
	PkgIsorder     int8      //是否需要预约
	PkgIsrefund    int8      //是否支持退款
	PkgIsinhouse   int8      //是否仅限堂食
	PkgIsnew       int8      //是否新品
	PkgIsrecommend int8      //是否强烈推荐
	PkgValiddays   int       //有效期（天）
	PkgFollows     int       //今日关注人数
	PkgCreatetime  time.Time //上架日期
	PkgValidtime   time.Time //生效日期
	PkgExittime    time.Time //下架日期
	PkgLinks       string    //更多优惠链接地址
	PkgInfo        string    //套餐说明
	Adder          int       //创建人
	Addtime        time.Time //创建时间
	Moder          int       //修改人
	Modtime        time.Time //修改时间
	DeleteStatus   int8      //0:未知，1：未删除，2：已删除
}

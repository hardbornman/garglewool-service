package model

import (
	"time"
)

// 套餐表
type Setmeal struct {
	Setmealid    int       //套餐ID
	Shopid       int       //店铺ID
	Shopname     string    //店铺名称【外键表:店铺表】
	Pkgcode      string    //套餐号
	Setmealtype  string    //套餐类别
	Title        string    //套餐标题
	People       int       //人数
	Isorder      bool      //是否需要预约
	Isrefund     bool      //是否支持退款
	Isinhouse    bool      //是否仅限堂食
	Isnew        bool      //是否新品
	Isrecommend  bool      //是否强烈推荐
	Validdays    int       //有效期（天）
	Watchers     int       //今日关注人数
	Createtime   time.Time //上架日期
	Validtime    time.Time //生效日期
	Exittime     time.Time //下架日期
	Links        string    //更多优惠链接地址
	Info         string    //套餐说明
	Adder        int       //创建人
	Addtime      time.Time //创建时间
	Moder        int       //修改人
	Modtime      time.Time //修改时间
	DeleteStatus int8      //0:未知，1：未删除，2：已删除
}

package model

import (
	"time"
)

// 省，直辖市
type SProvince struct {
	ProvinceId   int64     //省ID
	ProvinceName string    //省名称
	DateCreated  time.Time //创建日期
	DateUpdated  time.Time //更新日期
}

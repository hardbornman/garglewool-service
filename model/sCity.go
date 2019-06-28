package model

import (
	"time"
)

// 市
type SCity struct {
	CityId       int64     //市ID
	CityName     string    //市名称
	ZipCode      string    //编码
	ProvinceId   int64     //省ID
	ProvinceName string    //省名称【外键表:省，直辖市】
	DateCreated  time.Time //创建日期
	DateUpdated  time.Time //更新日期
}

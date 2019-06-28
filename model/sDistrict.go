package model

import (
	"time"
)

// 区县
type SDistrict struct {
	DistrictId   int64     //区ID
	DistrictName string    //区名称
	CityId       int64     //市ID
	CityName     string    //市名称【外键表:市】
	DateCreated  time.Time //创建日期
	DateUpdated  time.Time //更新日期
}

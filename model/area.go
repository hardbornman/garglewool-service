package model

// 区域表
type Area struct {
	Areaid     int    //序号
	RegionName string //区域名称
	RegionCode string //区域编码
	ParentCode string //父级编码
	Longitude  string //经度
	Latitude   string //纬度
}

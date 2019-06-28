package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
)

// 获取【区域表】信息
func GetAreaInfo(areaid int) (area model.Area, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【区域表】信息
	area, err = dao.AreaDao.Get(areaid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【序号】批量获取【区域表】列表
func GetInAreaList(areaids []int) (areas []model.Area, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	areas, err = dao.AreaDao.GetIn(areaids)
	if err != nil {
		return
	}
	count = len(areas)
	//endregion

	return
}

// 分页查询【区域表】列表
func GetAreaList(regionName string, longitude string, latitude string, pageIndex, pageSize int) (areas []model.Area, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.AreaDao.GetRowCount(regionName, longitude, latitude)
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	areas, err = dao.AreaDao.GetRowList(regionName, longitude, latitude, pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【区域表】信息
func InsertAreaInfo(areaid int, regionName string, regionCode string, parentCode string, longitude string, latitude string) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	var err error
	isExist := false
	var areaInfo model.Area

	//region 根据【序号】查询【区域表】信息
	{
		areaInfo, err = dao.AreaDao.Get(areaid)
		if err != nil {
			return false, err
		}
		if areaInfo.Areaid > 0 {
			return false, errors.New("【序号】已存在,不能重复添加")
		}
	}
	//endregion

	//region 构造【区域表】信息
	areaInfo.Areaid = areaid
	areaInfo.RegionName = regionName
	areaInfo.RegionCode = regionCode
	areaInfo.ParentCode = parentCode
	areaInfo.Longitude = longitude
	areaInfo.Latitude = latitude
	//endregion

	if isExist {

		//region 修改【区域表】信息
		isSuccess, err := dao.AreaDao.Update(&areaInfo)
		if err != nil {
			return false, errors.New("插入【区域表】出错:" + err.Error())
		}
		if !isSuccess {
			return false, errors.New("插入【区域表】失败")
		}
		return isSuccess, nil
		//endregion

	} else {

		//region 插入数据库
		isSuccess, err := dao.AreaDao.Insert(&areaInfo)
		if err != nil {
			return false, errors.New("插入【区域表】出错:" + err.Error())
		}
		if !isSuccess {
			return false, errors.New("插入【区域表】失败")
		}
		return isSuccess, nil
		//endregion

	}

}

//修改【区域表】信息
func UpdateAreaInfo(areaid int, regionName string, regionCode string, parentCode string, longitude string, latitude string) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【区域表】信息
	areaInfo, err := dao.AreaDao.Get(areaid)
	if err != nil {
		return false, errors.New("查询【区域表】信息出错:" + err.Error())
	}
	if areaInfo.Areaid <= 0 {
		return false, errors.New("【区域表】信息不存在")
	}
	//endregion

	//region 验证是否需要执行修改
	if areaInfo.Areaid == areaid &&
		areaInfo.RegionName == regionName &&
		areaInfo.RegionCode == regionCode &&
		areaInfo.ParentCode == parentCode &&
		areaInfo.Longitude == longitude &&
		areaInfo.Latitude == latitude {
		return true, nil
	}
	//endregion

	//region 构造【区域表】信息
	areaInfo.RegionName = regionName
	areaInfo.RegionCode = regionCode
	areaInfo.ParentCode = parentCode
	areaInfo.Longitude = longitude
	areaInfo.Latitude = latitude
	//endregion

	//region 修改【区域表】信息
	return dao.AreaDao.Update(&areaInfo)
	//endregion

}

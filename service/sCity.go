package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【市】信息
func GetSCityInfo(cityId int64) (sCity model.SCity, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【市】信息
	sCity, err = dao.SCityDao.Get(cityId)
	if err != nil {
		return
	}
	//endregion

	//region 获取【省，直辖市】信息
	if sCity.ProvinceId > 0 {
		var sProvince model.SProvince
		sProvince, err = dao.SProvinceDao.Get(sCity.ProvinceId)
		if err != nil {
			return
		}
		sCity.ProvinceName = sProvince.ProvinceName
	}
	//endregion

	return
}

// 根据【市ID】批量获取【市】列表
func GetInSCityList(cityIds []int64) (sCitys []model.SCity, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	sCitys, err = dao.SCityDao.GetIn(cityIds)
	if err != nil {
		return
	}
	count = len(sCitys)
	if len(sCitys) > 0 && sCitys[0].CityId > 0 {

		//region 查询外键表【省，直辖市】列表
		provinceIds := make([]int64, 0)
		var sProvinces []model.SProvince
		for _, sCity := range sCitys {
			if sCity.ProvinceId > 0 {
				provinceIds = append(provinceIds, sCity.ProvinceId)
			}
		}
		if len(provinceIds) > 0 {
			sProvinces, err = dao.SProvinceDao.GetIn(provinceIds)
			if err != nil {
				return
			}
		}
		//endregion

		//region 处理外键表
		for i, sCity := range sCitys {

			//region 处理【省，直辖市】列表
			if len(sProvinces) > 0 {
				for _, sProvince := range sProvinces {
					if sCity.ProvinceId == sProvince.ProvinceId {
						sCitys[i].ProvinceName = sProvince.ProvinceName
						break
					}
				}
			}
			//endregion

		}
		//endregion

	}
	//endregion

	return
}

// 分页查询【市】列表
func GetSCityList(cityName string, provinceId int64, pageIndex, pageSize int) (sCitys []model.SCity, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.SCityDao.GetRowCount(cityName, provinceId)
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	sCitys, err = dao.SCityDao.GetRowList(cityName, provinceId, pageIndex, pageSize)
	if err != nil {
		return
	}
	if len(sCitys) > 0 && sCitys[0].CityId > 0 {

		//region 查询外键表【省，直辖市】列表
		provinceIds := make([]int64, 0)
		var sProvinces []model.SProvince
		for _, sCity := range sCitys {
			if sCity.ProvinceId > 0 {
				provinceIds = append(provinceIds, sCity.ProvinceId)
			}
		}
		if len(provinceIds) > 0 {
			sProvinces, err = dao.SProvinceDao.GetIn(provinceIds)
			if err != nil {
				return
			}
		}
		//endregion

		//region 处理外键表
		for i, sCity := range sCitys {

			//region 处理【省，直辖市】列表
			if len(sProvinces) > 0 {
				for _, sProvince := range sProvinces {
					if sCity.ProvinceId == sProvince.ProvinceId {
						sCitys[i].ProvinceName = sProvince.ProvinceName
						break
					}
				}
			}
			//endregion

		}
		//endregion

	}
	//endregion

	return
}

// 插入【市】信息
func InsertSCityInfo(cityId int64, cityName string, zipCode string, provinceId int64, dateCreated time.Time, dateUpdated time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	var err error
	isExist := false
	var sCityInfo model.SCity

	//region 根据【市ID】查询【市】信息
	{
		sCityInfo, err = dao.SCityDao.Get(cityId)
		if err != nil {
			return false, err
		}
		if sCityInfo.CityId > 0 {
			return false, errors.New("【市ID】已存在,不能重复添加")
		}
	}
	//endregion

	//region 判断provinceId值是否存在
	s_province, err := dao.SProvinceDao.Get(provinceId)
	if err != nil {
		return false, err
	}
	if s_province.ProvinceId <= 0 {
		return false, errors.New("provinceId值不存在")
	}
	//if s_province.DeleteStatus != 1 {
	//	return false, errors.New("provinceId值已删除")
	//}
	//endregion

	//region 构造【市】信息
	sCityInfo.CityId = cityId
	sCityInfo.CityName = cityName
	sCityInfo.ZipCode = zipCode
	sCityInfo.ProvinceId = provinceId
	sCityInfo.DateCreated = dateCreated
	sCityInfo.DateUpdated = dateUpdated
	//endregion

	if isExist {

		//region 修改【市】信息
		isSuccess, err := dao.SCityDao.Update(&sCityInfo)
		if err != nil {
			return false, errors.New("插入【市】出错:" + err.Error())
		}
		if !isSuccess {
			return false, errors.New("插入【市】失败")
		}
		return isSuccess, nil
		//endregion

	} else {

		//region 插入数据库
		isSuccess, err := dao.SCityDao.Insert(&sCityInfo)
		if err != nil {
			return false, errors.New("插入【市】出错:" + err.Error())
		}
		if !isSuccess {
			return false, errors.New("插入【市】失败")
		}
		return isSuccess, nil
		//endregion

	}

}

//修改【市】信息
func UpdateSCityInfo(cityId int64, cityName string, zipCode string, provinceId int64, dateCreated time.Time, dateUpdated time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【市】信息
	sCityInfo, err := dao.SCityDao.Get(cityId)
	if err != nil {
		return false, errors.New("查询【市】信息出错:" + err.Error())
	}
	if sCityInfo.CityId <= 0 {
		return false, errors.New("【市】信息不存在")
	}
	//endregion

	//region 判断provinceId值是否存在
	s_province, err := dao.SProvinceDao.Get(provinceId)
	if err != nil {
		return false, err
	}
	if s_province.ProvinceId <= 0 {
		return false, errors.New("provinceId值不存在")
	}
	//if s_province.DeleteStatus != 1 {
	//	return false, errors.New("provinceId值已删除")
	//}
	//endregion

	//region 验证是否需要执行修改
	if sCityInfo.CityId == cityId &&
		sCityInfo.CityName == cityName &&
		sCityInfo.ZipCode == zipCode &&
		sCityInfo.ProvinceId == provinceId &&
		sCityInfo.DateCreated == dateCreated &&
		sCityInfo.DateUpdated == dateUpdated {
		return true, nil
	}
	//endregion

	//region 构造【市】信息
	sCityInfo.CityName = cityName
	sCityInfo.ZipCode = zipCode
	sCityInfo.ProvinceId = provinceId
	sCityInfo.DateCreated = dateCreated
	sCityInfo.DateUpdated = dateUpdated
	//endregion

	//region 修改【市】信息
	return dao.SCityDao.Update(&sCityInfo)
	//endregion

}

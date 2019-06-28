package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【区县】信息
func GetSDistrictInfo(districtId int64) (sDistrict model.SDistrict, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【区县】信息
	sDistrict, err = dao.SDistrictDao.Get(districtId)
	if err != nil {
		return
	}
	//endregion

	//region 获取【市】信息
	if sDistrict.CityId > 0 {
		var sCity model.SCity
		sCity, err = dao.SCityDao.Get(sDistrict.CityId)
		if err != nil {
			return
		}
		sDistrict.CityName = sCity.CityName
	}
	//endregion

	return
}

// 根据【区ID】批量获取【区县】列表
func GetInSDistrictList(districtIds []int64) (sDistricts []model.SDistrict, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	sDistricts, err = dao.SDistrictDao.GetIn(districtIds)
	if err != nil {
		return
	}
	count = len(sDistricts)
	if len(sDistricts) > 0 && sDistricts[0].DistrictId > 0 {

		//region 查询外键表【市】列表
		cityIds := make([]int64, 0)
		var sCitys []model.SCity
		for _, sDistrict := range sDistricts {
			if sDistrict.CityId > 0 {
				cityIds = append(cityIds, sDistrict.CityId)
			}
		}
		if len(cityIds) > 0 {
			sCitys, err = dao.SCityDao.GetIn(cityIds)
			if err != nil {
				return
			}
		}
		//endregion

		//region 处理外键表
		for i, sDistrict := range sDistricts {

			//region 处理【市】列表
			if len(sCitys) > 0 {
				for _, sCity := range sCitys {
					if sDistrict.CityId == sCity.CityId {
						sDistricts[i].CityName = sCity.CityName
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

// 分页查询【区县】列表
func GetSDistrictList(districtName string, cityId int64, pageIndex, pageSize int) (sDistricts []model.SDistrict, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.SDistrictDao.GetRowCount(districtName, cityId)
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	sDistricts, err = dao.SDistrictDao.GetRowList(districtName, cityId, pageIndex, pageSize)
	if err != nil {
		return
	}
	if len(sDistricts) > 0 && sDistricts[0].DistrictId > 0 {

		//region 查询外键表【市】列表
		cityIds := make([]int64, 0)
		var sCitys []model.SCity
		for _, sDistrict := range sDistricts {
			if sDistrict.CityId > 0 {
				cityIds = append(cityIds, sDistrict.CityId)
			}
		}
		if len(cityIds) > 0 {
			sCitys, err = dao.SCityDao.GetIn(cityIds)
			if err != nil {
				return
			}
		}
		//endregion

		//region 处理外键表
		for i, sDistrict := range sDistricts {

			//region 处理【市】列表
			if len(sCitys) > 0 {
				for _, sCity := range sCitys {
					if sDistrict.CityId == sCity.CityId {
						sDistricts[i].CityName = sCity.CityName
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

// 插入【区县】信息
func InsertSDistrictInfo(districtId int64, districtName string, cityId int64, dateCreated time.Time, dateUpdated time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	var err error
	isExist := false
	var sDistrictInfo model.SDistrict

	//region 根据【区ID】查询【区县】信息
	{
		sDistrictInfo, err = dao.SDistrictDao.Get(districtId)
		if err != nil {
			return false, err
		}
		if sDistrictInfo.DistrictId > 0 {
			return false, errors.New("【区ID】已存在,不能重复添加")
		}
	}
	//endregion

	//region 判断cityId值是否存在
	s_city, err := dao.SCityDao.Get(cityId)
	if err != nil {
		return false, err
	}
	if s_city.CityId <= 0 {
		return false, errors.New("cityId值不存在")
	}
	//if s_city.DeleteStatus != 1 {
	//	return false, errors.New("cityId值已删除")
	//}
	//endregion

	//region 构造【区县】信息
	sDistrictInfo.DistrictId = districtId
	sDistrictInfo.DistrictName = districtName
	sDistrictInfo.CityId = cityId
	sDistrictInfo.DateCreated = dateCreated
	sDistrictInfo.DateUpdated = dateUpdated
	//endregion

	if isExist {

		//region 修改【区县】信息
		isSuccess, err := dao.SDistrictDao.Update(&sDistrictInfo)
		if err != nil {
			return false, errors.New("插入【区县】出错:" + err.Error())
		}
		if !isSuccess {
			return false, errors.New("插入【区县】失败")
		}
		return isSuccess, nil
		//endregion

	} else {

		//region 插入数据库
		isSuccess, err := dao.SDistrictDao.Insert(&sDistrictInfo)
		if err != nil {
			return false, errors.New("插入【区县】出错:" + err.Error())
		}
		if !isSuccess {
			return false, errors.New("插入【区县】失败")
		}
		return isSuccess, nil
		//endregion

	}

}

//修改【区县】信息
func UpdateSDistrictInfo(districtId int64, districtName string, cityId int64, dateCreated time.Time, dateUpdated time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【区县】信息
	sDistrictInfo, err := dao.SDistrictDao.Get(districtId)
	if err != nil {
		return false, errors.New("查询【区县】信息出错:" + err.Error())
	}
	if sDistrictInfo.DistrictId <= 0 {
		return false, errors.New("【区县】信息不存在")
	}
	//endregion

	//region 判断cityId值是否存在
	s_city, err := dao.SCityDao.Get(cityId)
	if err != nil {
		return false, err
	}
	if s_city.CityId <= 0 {
		return false, errors.New("cityId值不存在")
	}
	//if s_city.DeleteStatus != 1 {
	//	return false, errors.New("cityId值已删除")
	//}
	//endregion

	//region 验证是否需要执行修改
	if sDistrictInfo.DistrictId == districtId &&
		sDistrictInfo.DistrictName == districtName &&
		sDistrictInfo.CityId == cityId &&
		sDistrictInfo.DateCreated == dateCreated &&
		sDistrictInfo.DateUpdated == dateUpdated {
		return true, nil
	}
	//endregion

	//region 构造【区县】信息
	sDistrictInfo.DistrictName = districtName
	sDistrictInfo.CityId = cityId
	sDistrictInfo.DateCreated = dateCreated
	sDistrictInfo.DateUpdated = dateUpdated
	//endregion

	//region 修改【区县】信息
	return dao.SDistrictDao.Update(&sDistrictInfo)
	//endregion

}

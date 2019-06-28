package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【省，直辖市】信息
func GetSProvinceInfo(provinceId int64) (sProvince model.SProvince, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【省，直辖市】信息
	sProvince, err = dao.SProvinceDao.Get(provinceId)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【省ID】批量获取【省，直辖市】列表
func GetInSProvinceList(provinceIds []int64) (sProvinces []model.SProvince, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	sProvinces, err = dao.SProvinceDao.GetIn(provinceIds)
	if err != nil {
		return
	}
	count = len(sProvinces)
	//endregion

	return
}

// 分页查询【省，直辖市】列表
func GetSProvinceList(provinceName string, pageIndex, pageSize int) (sProvinces []model.SProvince, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.SProvinceDao.GetRowCount(provinceName)
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	sProvinces, err = dao.SProvinceDao.GetRowList(provinceName, pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【省，直辖市】信息
func InsertSProvinceInfo(provinceId int64, provinceName string, dateCreated time.Time, dateUpdated time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	var err error
	isExist := false
	var sProvinceInfo model.SProvince

	//region 根据【省ID】查询【省，直辖市】信息
	{
		sProvinceInfo, err = dao.SProvinceDao.Get(provinceId)
		if err != nil {
			return false, err
		}
		if sProvinceInfo.ProvinceId > 0 {
			return false, errors.New("【省ID】已存在,不能重复添加")
		}
	}
	//endregion

	//region 构造【省，直辖市】信息
	sProvinceInfo.ProvinceId = provinceId
	sProvinceInfo.ProvinceName = provinceName
	sProvinceInfo.DateCreated = dateCreated
	sProvinceInfo.DateUpdated = dateUpdated
	//endregion

	if isExist {

		//region 修改【省，直辖市】信息
		isSuccess, err := dao.SProvinceDao.Update(&sProvinceInfo)
		if err != nil {
			return false, errors.New("插入【省，直辖市】出错:" + err.Error())
		}
		if !isSuccess {
			return false, errors.New("插入【省，直辖市】失败")
		}
		return isSuccess, nil
		//endregion

	} else {

		//region 插入数据库
		isSuccess, err := dao.SProvinceDao.Insert(&sProvinceInfo)
		if err != nil {
			return false, errors.New("插入【省，直辖市】出错:" + err.Error())
		}
		if !isSuccess {
			return false, errors.New("插入【省，直辖市】失败")
		}
		return isSuccess, nil
		//endregion

	}

}

//修改【省，直辖市】信息
func UpdateSProvinceInfo(provinceId int64, provinceName string, dateCreated time.Time, dateUpdated time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【省，直辖市】信息
	sProvinceInfo, err := dao.SProvinceDao.Get(provinceId)
	if err != nil {
		return false, errors.New("查询【省，直辖市】信息出错:" + err.Error())
	}
	if sProvinceInfo.ProvinceId <= 0 {
		return false, errors.New("【省，直辖市】信息不存在")
	}
	//endregion

	//region 验证是否需要执行修改
	if sProvinceInfo.ProvinceId == provinceId &&
		sProvinceInfo.ProvinceName == provinceName &&
		sProvinceInfo.DateCreated == dateCreated &&
		sProvinceInfo.DateUpdated == dateUpdated {
		return true, nil
	}
	//endregion

	//region 构造【省，直辖市】信息
	sProvinceInfo.ProvinceName = provinceName
	sProvinceInfo.DateCreated = dateCreated
	sProvinceInfo.DateUpdated = dateUpdated
	//endregion

	//region 修改【省，直辖市】信息
	return dao.SProvinceDao.Update(&sProvinceInfo)
	//endregion

}

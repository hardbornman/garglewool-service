package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【字典分类】信息
func GetDictionarycategoryInfo(dictionarycategoryid int) (dictionarycategory model.Dictionarycategory, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【字典分类】信息
	dictionarycategory, err = dao.DictionarycategoryDao.Get(dictionarycategoryid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【字典分类ID】批量获取【字典分类】列表
func GetInDictionarycategoryList(dictionarycategoryids []int) (dictionarycategorys []model.Dictionarycategory, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	dictionarycategorys, err = dao.DictionarycategoryDao.GetIn(dictionarycategoryids)
	if err != nil {
		return
	}
	count = len(dictionarycategorys)
	//endregion

	return
}

// 分页查询【字典分类】列表
func GetDictionarycategoryList(pageIndex, pageSize int) (dictionarycategorys []model.Dictionarycategory, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.DictionarycategoryDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	dictionarycategorys, err = dao.DictionarycategoryDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【字典分类】信息
func InsertDictionarycategoryInfo(categorykey string, categoryvalue string, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//var err error
	isExist := false
	var dictionarycategoryInfo model.Dictionarycategory

	//region 构造【字典分类】信息
	dictionarycategoryInfo.Categorykey = categorykey
	dictionarycategoryInfo.Categoryvalue = categoryvalue
	dictionarycategoryInfo.Adder = adder
	dictionarycategoryInfo.Addtime = addtime
	dictionarycategoryInfo.Moder = moder
	dictionarycategoryInfo.Modtime = modtime
	dictionarycategoryInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【字典分类】信息
		isSuccess, err := dao.DictionarycategoryDao.Update(&dictionarycategoryInfo)
		if err != nil {
			return -1, errors.New("插入【字典分类】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【字典分类】失败")
		}
		return dictionarycategoryInfo.Dictionarycategoryid, nil
		//endregion

	} else {

		//region 插入数据库
		dictionarycategoryid, err := dao.DictionarycategoryDao.Insert(&dictionarycategoryInfo)
		if err != nil {
			return -1, errors.New("插入【字典分类】出错:" + err.Error())
		}
		if dictionarycategoryid <= 0 {
			return -1, errors.New("插入【字典分类】失败")
		}
		return int(dictionarycategoryid), nil
		//endregion

	}

}

//修改【字典分类】信息
func UpdateDictionarycategoryInfo(dictionarycategoryid int, categorykey string, categoryvalue string, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【字典分类】信息
	dictionarycategoryInfo, err := dao.DictionarycategoryDao.Get(dictionarycategoryid)
	if err != nil {
		return false, errors.New("查询【字典分类】信息出错:" + err.Error())
	}
	if dictionarycategoryInfo.Dictionarycategoryid <= 0 {
		return false, errors.New("【字典分类】信息不存在")
	}
	if dictionarycategoryInfo.DeleteStatus != 1 {
		return false, errors.New("【字典分类】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if dictionarycategoryInfo.Dictionarycategoryid == dictionarycategoryid &&
		dictionarycategoryInfo.Categorykey == categorykey &&
		dictionarycategoryInfo.Categoryvalue == categoryvalue &&
		dictionarycategoryInfo.Adder == adder &&
		dictionarycategoryInfo.Addtime == addtime &&
		dictionarycategoryInfo.Moder == moder &&
		dictionarycategoryInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【字典分类】信息
	dictionarycategoryInfo.Categorykey = categorykey
	dictionarycategoryInfo.Categoryvalue = categoryvalue
	dictionarycategoryInfo.Adder = adder
	dictionarycategoryInfo.Addtime = addtime
	dictionarycategoryInfo.Moder = moder
	dictionarycategoryInfo.Modtime = modtime
	dictionarycategoryInfo.DeleteStatus = 1
	//endregion

	//region 修改【字典分类】信息
	return dao.DictionarycategoryDao.Update(&dictionarycategoryInfo)
	//endregion

}

//删除【字典分类】信息
func DeleteDictionarycategoryInfo(dictionarycategoryid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【字典分类】信息
	{
		dictionarycategoryInfo, err := dao.DictionarycategoryDao.Get(dictionarycategoryid)
		if err != nil {
			return false, errors.New("查询【字典分类】信息出错:" + err.Error())
		}
		if dictionarycategoryInfo.Dictionarycategoryid <= 0 {
			return false, errors.New("没有找到【字典分类】信息")
		}
		if dictionarycategoryInfo.DeleteStatus != 1 {
			return false, errors.New("【字典分类】信息已被删除")
		}
	}
	//endregion

	//region 删除【字典分类】信息
	return dao.DictionarycategoryDao.Delete(dictionarycategoryid)
	//endregion

}

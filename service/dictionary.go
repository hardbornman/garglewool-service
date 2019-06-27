package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【字典表】信息
func GetDictionaryInfo(dictionaryid int) (dictionary model.Dictionary, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【字典表】信息
	dictionary, err = dao.DictionaryDao.Get(dictionaryid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【字典ID】批量获取【字典表】列表
func GetInDictionaryList(dictionaryids []int) (dictionarys []model.Dictionary, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	dictionarys, err = dao.DictionaryDao.GetIn(dictionaryids)
	if err != nil {
		return
	}
	count = len(dictionarys)
	//endregion

	return
}

// 分页查询【字典表】列表
func GetDictionaryList(pageIndex, pageSize int) (dictionarys []model.Dictionary, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.DictionaryDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	dictionarys, err = dao.DictionaryDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【字典表】信息
func InsertDictionaryInfo(categorykey string, dictkey string, dictvalue string, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//var err error
	isExist := false
	var dictionaryInfo model.Dictionary

	//region 构造【字典表】信息
	dictionaryInfo.Categorykey = categorykey
	dictionaryInfo.Dictkey = dictkey
	dictionaryInfo.Dictvalue = dictvalue
	dictionaryInfo.Adder = adder
	dictionaryInfo.Addtime = addtime
	dictionaryInfo.Moder = moder
	dictionaryInfo.Modtime = modtime
	dictionaryInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【字典表】信息
		isSuccess, err := dao.DictionaryDao.Update(&dictionaryInfo)
		if err != nil {
			return -1, errors.New("插入【字典表】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【字典表】失败")
		}
		return dictionaryInfo.Dictionaryid, nil
		//endregion

	} else {

		//region 插入数据库
		dictionaryid, err := dao.DictionaryDao.Insert(&dictionaryInfo)
		if err != nil {
			return -1, errors.New("插入【字典表】出错:" + err.Error())
		}
		if dictionaryid <= 0 {
			return -1, errors.New("插入【字典表】失败")
		}
		return int(dictionaryid), nil
		//endregion

	}

}

//修改【字典表】信息
func UpdateDictionaryInfo(dictionaryid int, categorykey string, dictkey string, dictvalue string, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【字典表】信息
	dictionaryInfo, err := dao.DictionaryDao.Get(dictionaryid)
	if err != nil {
		return false, errors.New("查询【字典表】信息出错:" + err.Error())
	}
	if dictionaryInfo.Dictionaryid <= 0 {
		return false, errors.New("【字典表】信息不存在")
	}
	if dictionaryInfo.DeleteStatus != 1 {
		return false, errors.New("【字典表】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if dictionaryInfo.Dictionaryid == dictionaryid &&
		dictionaryInfo.Categorykey == categorykey &&
		dictionaryInfo.Dictkey == dictkey &&
		dictionaryInfo.Dictvalue == dictvalue &&
		dictionaryInfo.Adder == adder &&
		dictionaryInfo.Addtime == addtime &&
		dictionaryInfo.Moder == moder &&
		dictionaryInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【字典表】信息
	dictionaryInfo.Categorykey = categorykey
	dictionaryInfo.Dictkey = dictkey
	dictionaryInfo.Dictvalue = dictvalue
	dictionaryInfo.Adder = adder
	dictionaryInfo.Addtime = addtime
	dictionaryInfo.Moder = moder
	dictionaryInfo.Modtime = modtime
	dictionaryInfo.DeleteStatus = 1
	//endregion

	//region 修改【字典表】信息
	return dao.DictionaryDao.Update(&dictionaryInfo)
	//endregion

}

//删除【字典表】信息
func DeleteDictionaryInfo(dictionaryid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【字典表】信息
	{
		dictionaryInfo, err := dao.DictionaryDao.Get(dictionaryid)
		if err != nil {
			return false, errors.New("查询【字典表】信息出错:" + err.Error())
		}
		if dictionaryInfo.Dictionaryid <= 0 {
			return false, errors.New("没有找到【字典表】信息")
		}
		if dictionaryInfo.DeleteStatus != 1 {
			return false, errors.New("【字典表】信息已被删除")
		}
	}
	//endregion

	//region 删除【字典表】信息
	return dao.DictionaryDao.Delete(dictionaryid)
	//endregion

}

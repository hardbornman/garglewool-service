package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【字典表】信息
func GetSysDictionaryInfo(dictionaryid int) (sysDictionary model.SysDictionary, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【字典表】信息
	sysDictionary, err = dao.SysDictionaryDao.Get(dictionaryid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【字典ID】批量获取【字典表】列表
func GetInSysDictionaryList(dictionaryids []int) (sysDictionarys []model.SysDictionary, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	sysDictionarys, err = dao.SysDictionaryDao.GetIn(dictionaryids)
	if err != nil {
		return
	}
	count = len(sysDictionarys)
	//endregion

	return
}

// 分页查询【字典表】列表
func GetSysDictionaryList(pageIndex, pageSize int) (sysDictionarys []model.SysDictionary, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.SysDictionaryDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	sysDictionarys, err = dao.SysDictionaryDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【字典表】信息
func InsertSysDictionaryInfo(categorykey string, dictkey string, dictvalue string, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//var err error
	isExist := false
	var sysDictionaryInfo model.SysDictionary

	//region 构造【字典表】信息
	sysDictionaryInfo.Categorykey = categorykey
	sysDictionaryInfo.Dictkey = dictkey
	sysDictionaryInfo.Dictvalue = dictvalue
	sysDictionaryInfo.Adder = adder
	sysDictionaryInfo.Addtime = addtime
	sysDictionaryInfo.Moder = moder
	sysDictionaryInfo.Modtime = modtime
	sysDictionaryInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【字典表】信息
		isSuccess, err := dao.SysDictionaryDao.Update(&sysDictionaryInfo)
		if err != nil {
			return -1, errors.New("插入【字典表】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【字典表】失败")
		}
		return sysDictionaryInfo.Dictionaryid, nil
		//endregion

	} else {

		//region 插入数据库
		dictionaryid, err := dao.SysDictionaryDao.Insert(&sysDictionaryInfo)
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
func UpdateSysDictionaryInfo(dictionaryid int, categorykey string, dictkey string, dictvalue string, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【字典表】信息
	sysDictionaryInfo, err := dao.SysDictionaryDao.Get(dictionaryid)
	if err != nil {
		return false, errors.New("查询【字典表】信息出错:" + err.Error())
	}
	if sysDictionaryInfo.Dictionaryid <= 0 {
		return false, errors.New("【字典表】信息不存在")
	}
	if sysDictionaryInfo.DeleteStatus != 1 {
		return false, errors.New("【字典表】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if sysDictionaryInfo.Dictionaryid == dictionaryid &&
		sysDictionaryInfo.Categorykey == categorykey &&
		sysDictionaryInfo.Dictkey == dictkey &&
		sysDictionaryInfo.Dictvalue == dictvalue &&
		sysDictionaryInfo.Adder == adder &&
		sysDictionaryInfo.Addtime == addtime &&
		sysDictionaryInfo.Moder == moder &&
		sysDictionaryInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【字典表】信息
	sysDictionaryInfo.Categorykey = categorykey
	sysDictionaryInfo.Dictkey = dictkey
	sysDictionaryInfo.Dictvalue = dictvalue
	sysDictionaryInfo.Adder = adder
	sysDictionaryInfo.Addtime = addtime
	sysDictionaryInfo.Moder = moder
	sysDictionaryInfo.Modtime = modtime
	sysDictionaryInfo.DeleteStatus = 1
	//endregion

	//region 修改【字典表】信息
	return dao.SysDictionaryDao.Update(&sysDictionaryInfo)
	//endregion

}

//删除【字典表】信息
func DeleteSysDictionaryInfo(dictionaryid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【字典表】信息
	{
		sysDictionaryInfo, err := dao.SysDictionaryDao.Get(dictionaryid)
		if err != nil {
			return false, errors.New("查询【字典表】信息出错:" + err.Error())
		}
		if sysDictionaryInfo.Dictionaryid <= 0 {
			return false, errors.New("没有找到【字典表】信息")
		}
		if sysDictionaryInfo.DeleteStatus != 1 {
			return false, errors.New("【字典表】信息已被删除")
		}
	}
	//endregion

	//region 删除【字典表】信息
	return dao.SysDictionaryDao.Delete(dictionaryid)
	//endregion

}

package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【字典分类】信息
func GetSysDictionarycategoryInfo(dictionarycategoryid int) (sysDictionarycategory model.SysDictionarycategory, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【字典分类】信息
	sysDictionarycategory, err = dao.SysDictionarycategoryDao.Get(dictionarycategoryid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【字典分类ID】批量获取【字典分类】列表
func GetInSysDictionarycategoryList(dictionarycategoryids []int) (sysDictionarycategorys []model.SysDictionarycategory, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	sysDictionarycategorys, err = dao.SysDictionarycategoryDao.GetIn(dictionarycategoryids)
	if err != nil {
		return
	}
	count = len(sysDictionarycategorys)
	//endregion

	return
}

// 分页查询【字典分类】列表
func GetSysDictionarycategoryList(pageIndex, pageSize int) (sysDictionarycategorys []model.SysDictionarycategory, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.SysDictionarycategoryDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	sysDictionarycategorys, err = dao.SysDictionarycategoryDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【字典分类】信息
func InsertSysDictionarycategoryInfo(categorykey string, categoryvalue string, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//var err error
	isExist := false
	var sysDictionarycategoryInfo model.SysDictionarycategory

	//region 构造【字典分类】信息
	sysDictionarycategoryInfo.Categorykey = categorykey
	sysDictionarycategoryInfo.Categoryvalue = categoryvalue
	sysDictionarycategoryInfo.Adder = adder
	sysDictionarycategoryInfo.Addtime = addtime
	sysDictionarycategoryInfo.Moder = moder
	sysDictionarycategoryInfo.Modtime = modtime
	sysDictionarycategoryInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【字典分类】信息
		isSuccess, err := dao.SysDictionarycategoryDao.Update(&sysDictionarycategoryInfo)
		if err != nil {
			return -1, errors.New("插入【字典分类】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【字典分类】失败")
		}
		return sysDictionarycategoryInfo.Dictionarycategoryid, nil
		//endregion

	} else {

		//region 插入数据库
		dictionarycategoryid, err := dao.SysDictionarycategoryDao.Insert(&sysDictionarycategoryInfo)
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
func UpdateSysDictionarycategoryInfo(dictionarycategoryid int, categorykey string, categoryvalue string, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【字典分类】信息
	sysDictionarycategoryInfo, err := dao.SysDictionarycategoryDao.Get(dictionarycategoryid)
	if err != nil {
		return false, errors.New("查询【字典分类】信息出错:" + err.Error())
	}
	if sysDictionarycategoryInfo.Dictionarycategoryid <= 0 {
		return false, errors.New("【字典分类】信息不存在")
	}
	if sysDictionarycategoryInfo.DeleteStatus != 1 {
		return false, errors.New("【字典分类】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if sysDictionarycategoryInfo.Dictionarycategoryid == dictionarycategoryid &&
		sysDictionarycategoryInfo.Categorykey == categorykey &&
		sysDictionarycategoryInfo.Categoryvalue == categoryvalue &&
		sysDictionarycategoryInfo.Adder == adder &&
		sysDictionarycategoryInfo.Addtime == addtime &&
		sysDictionarycategoryInfo.Moder == moder &&
		sysDictionarycategoryInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【字典分类】信息
	sysDictionarycategoryInfo.Categorykey = categorykey
	sysDictionarycategoryInfo.Categoryvalue = categoryvalue
	sysDictionarycategoryInfo.Adder = adder
	sysDictionarycategoryInfo.Addtime = addtime
	sysDictionarycategoryInfo.Moder = moder
	sysDictionarycategoryInfo.Modtime = modtime
	sysDictionarycategoryInfo.DeleteStatus = 1
	//endregion

	//region 修改【字典分类】信息
	return dao.SysDictionarycategoryDao.Update(&sysDictionarycategoryInfo)
	//endregion

}

//删除【字典分类】信息
func DeleteSysDictionarycategoryInfo(dictionarycategoryid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【字典分类】信息
	{
		sysDictionarycategoryInfo, err := dao.SysDictionarycategoryDao.Get(dictionarycategoryid)
		if err != nil {
			return false, errors.New("查询【字典分类】信息出错:" + err.Error())
		}
		if sysDictionarycategoryInfo.Dictionarycategoryid <= 0 {
			return false, errors.New("没有找到【字典分类】信息")
		}
		if sysDictionarycategoryInfo.DeleteStatus != 1 {
			return false, errors.New("【字典分类】信息已被删除")
		}
	}
	//endregion

	//region 删除【字典分类】信息
	return dao.SysDictionarycategoryDao.Delete(dictionarycategoryid)
	//endregion

}

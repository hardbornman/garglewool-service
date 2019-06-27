package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【套餐明细表】信息
func GetSetmealdetailInfo(setmealdetailid int) (setmealdetail model.Setmealdetail, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【套餐明细表】信息
	setmealdetail, err = dao.SetmealdetailDao.Get(setmealdetailid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【套餐明细ID】批量获取【套餐明细表】列表
func GetInSetmealdetailList(setmealdetailids []int) (setmealdetails []model.Setmealdetail, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	setmealdetails, err = dao.SetmealdetailDao.GetIn(setmealdetailids)
	if err != nil {
		return
	}
	count = len(setmealdetails)
	//endregion

	return
}

// 分页查询【套餐明细表】列表
func GetSetmealdetailList(pageIndex, pageSize int) (setmealdetails []model.Setmealdetail, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.SetmealdetailDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	setmealdetails, err = dao.SetmealdetailDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【套餐明细表】信息
func InsertSetmealdetailInfo(setmealid int, name string, nums int, price float64, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	var err error
	isExist := false
	var setmealdetailInfo model.Setmealdetail

	//region 判断setmealid值是否存在
	setmeal, err := dao.SetmealDao.Get(setmealid)
	if err != nil {
		return -1, err
	}
	if setmeal.Setmealid <= 0 {
		return -1, errors.New("setmealid值不存在")
	}
	if setmeal.DeleteStatus != 1 {
		return -1, errors.New("setmealid值已删除")
	}
	//endregion

	//region 构造【套餐明细表】信息
	setmealdetailInfo.Setmealid = setmealid
	setmealdetailInfo.Name = name
	setmealdetailInfo.Nums = nums
	setmealdetailInfo.Price = price
	setmealdetailInfo.Adder = adder
	setmealdetailInfo.Addtime = addtime
	setmealdetailInfo.Moder = moder
	setmealdetailInfo.Modtime = modtime
	setmealdetailInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【套餐明细表】信息
		isSuccess, err := dao.SetmealdetailDao.Update(&setmealdetailInfo)
		if err != nil {
			return -1, errors.New("插入【套餐明细表】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【套餐明细表】失败")
		}
		return setmealdetailInfo.Setmealdetailid, nil
		//endregion

	} else {

		//region 插入数据库
		setmealdetailid, err := dao.SetmealdetailDao.Insert(&setmealdetailInfo)
		if err != nil {
			return -1, errors.New("插入【套餐明细表】出错:" + err.Error())
		}
		if setmealdetailid <= 0 {
			return -1, errors.New("插入【套餐明细表】失败")
		}
		return int(setmealdetailid), nil
		//endregion

	}

}

//修改【套餐明细表】信息
func UpdateSetmealdetailInfo(setmealdetailid int, setmealid int, name string, nums int, price float64, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【套餐明细表】信息
	setmealdetailInfo, err := dao.SetmealdetailDao.Get(setmealdetailid)
	if err != nil {
		return false, errors.New("查询【套餐明细表】信息出错:" + err.Error())
	}
	if setmealdetailInfo.Setmealdetailid <= 0 {
		return false, errors.New("【套餐明细表】信息不存在")
	}
	if setmealdetailInfo.DeleteStatus != 1 {
		return false, errors.New("【套餐明细表】信息已被删除")
	}
	//endregion

	//region 判断setmealid值是否存在
	setmeal, err := dao.SetmealDao.Get(setmealid)
	if err != nil {
		return false, err
	}
	if setmeal.Setmealid <= 0 {
		return false, errors.New("setmealid值不存在")
	}
	if setmeal.DeleteStatus != 1 {
		return false, errors.New("setmealid值已删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if setmealdetailInfo.Setmealdetailid == setmealdetailid &&
		setmealdetailInfo.Setmealid == setmealid &&
		setmealdetailInfo.Name == name &&
		setmealdetailInfo.Nums == nums &&
		setmealdetailInfo.Price == price &&
		setmealdetailInfo.Adder == adder &&
		setmealdetailInfo.Addtime == addtime &&
		setmealdetailInfo.Moder == moder &&
		setmealdetailInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【套餐明细表】信息
	setmealdetailInfo.Setmealid = setmealid
	setmealdetailInfo.Name = name
	setmealdetailInfo.Nums = nums
	setmealdetailInfo.Price = price
	setmealdetailInfo.Adder = adder
	setmealdetailInfo.Addtime = addtime
	setmealdetailInfo.Moder = moder
	setmealdetailInfo.Modtime = modtime
	setmealdetailInfo.DeleteStatus = 1
	//endregion

	//region 修改【套餐明细表】信息
	return dao.SetmealdetailDao.Update(&setmealdetailInfo)
	//endregion

}

//删除【套餐明细表】信息
func DeleteSetmealdetailInfo(setmealdetailid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【套餐明细表】信息
	{
		setmealdetailInfo, err := dao.SetmealdetailDao.Get(setmealdetailid)
		if err != nil {
			return false, errors.New("查询【套餐明细表】信息出错:" + err.Error())
		}
		if setmealdetailInfo.Setmealdetailid <= 0 {
			return false, errors.New("没有找到【套餐明细表】信息")
		}
		if setmealdetailInfo.DeleteStatus != 1 {
			return false, errors.New("【套餐明细表】信息已被删除")
		}
	}
	//endregion

	//region 删除【套餐明细表】信息
	return dao.SetmealdetailDao.Delete(setmealdetailid)
	//endregion

}

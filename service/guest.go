package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【买家客户表】信息
func GetGuestInfo(guestid int) (guest model.Guest, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【买家客户表】信息
	guest, err = dao.GuestDao.Get(guestid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【用户ID】批量获取【买家客户表】列表
func GetInGuestList(guestids []int) (guests []model.Guest, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	guests, err = dao.GuestDao.GetIn(guestids)
	if err != nil {
		return
	}
	count = len(guests)
	//endregion

	return
}

// 分页查询【买家客户表】列表
func GetGuestList(pageIndex, pageSize int) (guests []model.Guest, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.GuestDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	guests, err = dao.GuestDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【买家客户表】信息
func InsertGuestInfo(name string, password string, phone string, golds int, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//var err error
	isExist := false
	var guestInfo model.Guest

	//region 构造【买家客户表】信息
	guestInfo.Name = name
	guestInfo.Password = password
	guestInfo.Phone = phone
	guestInfo.Golds = golds
	guestInfo.Adder = adder
	guestInfo.Addtime = addtime
	guestInfo.Moder = moder
	guestInfo.Modtime = modtime
	guestInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【买家客户表】信息
		isSuccess, err := dao.GuestDao.Update(&guestInfo)
		if err != nil {
			return -1, errors.New("插入【买家客户表】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【买家客户表】失败")
		}
		return guestInfo.Guestid, nil
		//endregion

	} else {

		//region 插入数据库
		guestid, err := dao.GuestDao.Insert(&guestInfo)
		if err != nil {
			return -1, errors.New("插入【买家客户表】出错:" + err.Error())
		}
		if guestid <= 0 {
			return -1, errors.New("插入【买家客户表】失败")
		}
		return int(guestid), nil
		//endregion

	}

}

//修改【买家客户表】信息
func UpdateGuestInfo(guestid int, name string, password string, phone string, golds int, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【买家客户表】信息
	guestInfo, err := dao.GuestDao.Get(guestid)
	if err != nil {
		return false, errors.New("查询【买家客户表】信息出错:" + err.Error())
	}
	if guestInfo.Guestid <= 0 {
		return false, errors.New("【买家客户表】信息不存在")
	}
	if guestInfo.DeleteStatus != 1 {
		return false, errors.New("【买家客户表】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if guestInfo.Guestid == guestid &&
		guestInfo.Name == name &&
		guestInfo.Password == password &&
		guestInfo.Phone == phone &&
		guestInfo.Golds == golds &&
		guestInfo.Adder == adder &&
		guestInfo.Addtime == addtime &&
		guestInfo.Moder == moder &&
		guestInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【买家客户表】信息
	guestInfo.Name = name
	guestInfo.Password = password
	guestInfo.Phone = phone
	guestInfo.Golds = golds
	guestInfo.Adder = adder
	guestInfo.Addtime = addtime
	guestInfo.Moder = moder
	guestInfo.Modtime = modtime
	guestInfo.DeleteStatus = 1
	//endregion

	//region 修改【买家客户表】信息
	return dao.GuestDao.Update(&guestInfo)
	//endregion

}

//删除【买家客户表】信息
func DeleteGuestInfo(guestid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 根据【用户id】查询【抵用券管理】总数
	{
		count, err := dao.VoucherDao.GetRowCountByGuestid(guestid)
		if err != nil {
			return false, errors.New("查询【抵用券管理】总数出错:" + err.Error())
		}
		if count > 0 {
			return false, errors.New("【抵用券管理】存在相关记录")
		}
	}
	//endregion

	//region 查询【买家客户表】信息
	{
		guestInfo, err := dao.GuestDao.Get(guestid)
		if err != nil {
			return false, errors.New("查询【买家客户表】信息出错:" + err.Error())
		}
		if guestInfo.Guestid <= 0 {
			return false, errors.New("没有找到【买家客户表】信息")
		}
		if guestInfo.DeleteStatus != 1 {
			return false, errors.New("【买家客户表】信息已被删除")
		}
	}
	//endregion

	//region 删除【买家客户表】信息
	return dao.GuestDao.Delete(guestid)
	//endregion

}

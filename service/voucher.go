package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【抵用券管理】信息
func GetVoucherInfo(voucherid int) (voucher model.Voucher, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【抵用券管理】信息
	voucher, err = dao.VoucherDao.Get(voucherid)
	if err != nil {
		return
	}
	//endregion

	//region 获取【买家客户表】信息
	if voucher.Guestid > 0 {
		var guest model.Guest
		guest, err = dao.GuestDao.Get(voucher.Guestid)
		if err != nil {
			return
		}
		voucher.Name = guest.Name
	}
	//endregion

	return
}

// 根据【抵用券ID】批量获取【抵用券管理】列表
func GetInVoucherList(voucherids []int) (vouchers []model.Voucher, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	vouchers, err = dao.VoucherDao.GetIn(voucherids)
	if err != nil {
		return
	}
	count = len(vouchers)
	if len(vouchers) > 0 && vouchers[0].Voucherid > 0 {

		//region 查询外键表【买家客户表】列表
		guestids := make([]int, 0)
		var guests []model.Guest
		for _, voucher := range vouchers {
			if voucher.Guestid > 0 {
				guestids = append(guestids, voucher.Guestid)
			}
		}
		if len(guestids) > 0 {
			guests, err = dao.GuestDao.GetIn(guestids)
			if err != nil {
				return
			}
		}
		//endregion

		//region 处理外键表
		for i, voucher := range vouchers {

			//region 处理【买家客户表】列表
			if len(guests) > 0 {
				for _, guest := range guests {
					if voucher.Guestid == guest.Guestid {
						vouchers[i].Name = guest.Name
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

// 分页查询【抵用券管理】列表
func GetVoucherList(pageIndex, pageSize int) (vouchers []model.Voucher, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.VoucherDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	vouchers, err = dao.VoucherDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	if len(vouchers) > 0 && vouchers[0].Voucherid > 0 {

		//region 查询外键表【买家客户表】列表
		guestids := make([]int, 0)
		var guests []model.Guest
		for _, voucher := range vouchers {
			if voucher.Guestid > 0 {
				guestids = append(guestids, voucher.Guestid)
			}
		}
		if len(guestids) > 0 {
			guests, err = dao.GuestDao.GetIn(guestids)
			if err != nil {
				return
			}
		}
		//endregion

		//region 处理外键表
		for i, voucher := range vouchers {

			//region 处理【买家客户表】列表
			if len(guests) > 0 {
				for _, guest := range guests {
					if voucher.Guestid == guest.Guestid {
						vouchers[i].Name = guest.Name
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

// 插入【抵用券管理】信息
func InsertVoucherInfo(code string, guestid int, quota float64, reduce float64, validdays int, isinvalid bool, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	var err error
	isExist := false
	var voucherInfo model.Voucher

	//region 判断guestid值是否存在
	guest, err := dao.GuestDao.Get(guestid)
	if err != nil {
		return -1, err
	}
	if guest.Guestid <= 0 {
		return -1, errors.New("guestid值不存在")
	}
	if guest.DeleteStatus != 1 {
		return -1, errors.New("guestid值已删除")
	}
	//endregion

	//region 构造【抵用券管理】信息
	voucherInfo.Code = code
	voucherInfo.Guestid = guestid
	voucherInfo.Quota = quota
	voucherInfo.Reduce = reduce
	voucherInfo.Validdays = validdays
	voucherInfo.Isinvalid = isinvalid
	voucherInfo.Adder = adder
	voucherInfo.Addtime = addtime
	voucherInfo.Moder = moder
	voucherInfo.Modtime = modtime
	voucherInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【抵用券管理】信息
		isSuccess, err := dao.VoucherDao.Update(&voucherInfo)
		if err != nil {
			return -1, errors.New("插入【抵用券管理】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【抵用券管理】失败")
		}
		return voucherInfo.Voucherid, nil
		//endregion

	} else {

		//region 插入数据库
		voucherid, err := dao.VoucherDao.Insert(&voucherInfo)
		if err != nil {
			return -1, errors.New("插入【抵用券管理】出错:" + err.Error())
		}
		if voucherid <= 0 {
			return -1, errors.New("插入【抵用券管理】失败")
		}
		return int(voucherid), nil
		//endregion

	}

}

//修改【抵用券管理】信息
func UpdateVoucherInfo(voucherid int, code string, guestid int, quota float64, reduce float64, validdays int, isinvalid bool, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【抵用券管理】信息
	voucherInfo, err := dao.VoucherDao.Get(voucherid)
	if err != nil {
		return false, errors.New("查询【抵用券管理】信息出错:" + err.Error())
	}
	if voucherInfo.Voucherid <= 0 {
		return false, errors.New("【抵用券管理】信息不存在")
	}
	if voucherInfo.DeleteStatus != 1 {
		return false, errors.New("【抵用券管理】信息已被删除")
	}
	//endregion

	//region 判断guestid值是否存在
	guest, err := dao.GuestDao.Get(guestid)
	if err != nil {
		return false, err
	}
	if guest.Guestid <= 0 {
		return false, errors.New("guestid值不存在")
	}
	if guest.DeleteStatus != 1 {
		return false, errors.New("guestid值已删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if voucherInfo.Voucherid == voucherid &&
		voucherInfo.Code == code &&
		voucherInfo.Guestid == guestid &&
		voucherInfo.Quota == quota &&
		voucherInfo.Reduce == reduce &&
		voucherInfo.Validdays == validdays &&
		voucherInfo.Isinvalid == isinvalid &&
		voucherInfo.Adder == adder &&
		voucherInfo.Addtime == addtime &&
		voucherInfo.Moder == moder &&
		voucherInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【抵用券管理】信息
	voucherInfo.Code = code
	voucherInfo.Guestid = guestid
	voucherInfo.Quota = quota
	voucherInfo.Reduce = reduce
	voucherInfo.Validdays = validdays
	voucherInfo.Isinvalid = isinvalid
	voucherInfo.Adder = adder
	voucherInfo.Addtime = addtime
	voucherInfo.Moder = moder
	voucherInfo.Modtime = modtime
	voucherInfo.DeleteStatus = 1
	//endregion

	//region 修改【抵用券管理】信息
	return dao.VoucherDao.Update(&voucherInfo)
	//endregion

}

//删除【抵用券管理】信息
func DeleteVoucherInfo(voucherid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【抵用券管理】信息
	{
		voucherInfo, err := dao.VoucherDao.Get(voucherid)
		if err != nil {
			return false, errors.New("查询【抵用券管理】信息出错:" + err.Error())
		}
		if voucherInfo.Voucherid <= 0 {
			return false, errors.New("没有找到【抵用券管理】信息")
		}
		if voucherInfo.DeleteStatus != 1 {
			return false, errors.New("【抵用券管理】信息已被删除")
		}
	}
	//endregion

	//region 删除【抵用券管理】信息
	return dao.VoucherDao.Delete(voucherid)
	//endregion

}

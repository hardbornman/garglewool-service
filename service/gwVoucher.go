package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【抵用券管理】信息
func GetGwVoucherInfo(voucherid int) (gwVoucher model.GwVoucher, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【抵用券管理】信息
	gwVoucher, err = dao.GwVoucherDao.Get(voucherid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【抵用券ID】批量获取【抵用券管理】列表
func GetInGwVoucherList(voucherids []int) (gwVouchers []model.GwVoucher, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	gwVouchers, err = dao.GwVoucherDao.GetIn(voucherids)
	if err != nil {
		return
	}
	count = len(gwVouchers)
	//endregion

	return
}

// 分页查询【抵用券管理】列表
func GetGwVoucherList(pageIndex, pageSize int) (gwVouchers []model.GwVoucher, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.GwVoucherDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	gwVouchers, err = dao.GwVoucherDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【抵用券管理】信息
func InsertGwVoucherInfo(voucher_code string, voucher_userid int, voucher_quota float64, voucher_reduce float64, voucher_createtime time.Time, voucher_validdays int, voucher_isinvalid int8, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//var err error
	isExist := false
	var gwVoucherInfo model.GwVoucher

	//region 构造【抵用券管理】信息
	gwVoucherInfo.VoucherCode = voucher_code
	gwVoucherInfo.VoucherUserid = voucher_userid
	gwVoucherInfo.VoucherQuota = voucher_quota
	gwVoucherInfo.VoucherReduce = voucher_reduce
	gwVoucherInfo.VoucherCreatetime = voucher_createtime
	gwVoucherInfo.VoucherValiddays = voucher_validdays
	gwVoucherInfo.VoucherIsinvalid = voucher_isinvalid
	gwVoucherInfo.Adder = adder
	gwVoucherInfo.Addtime = addtime
	gwVoucherInfo.Moder = moder
	gwVoucherInfo.Modtime = modtime
	gwVoucherInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【抵用券管理】信息
		isSuccess, err := dao.GwVoucherDao.Update(&gwVoucherInfo)
		if err != nil {
			return -1, errors.New("插入【抵用券管理】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【抵用券管理】失败")
		}
		return gwVoucherInfo.Voucherid, nil
		//endregion

	} else {

		//region 插入数据库
		voucherid, err := dao.GwVoucherDao.Insert(&gwVoucherInfo)
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
func UpdateGwVoucherInfo(voucherid int, voucher_code string, voucher_userid int, voucher_quota float64, voucher_reduce float64, voucher_createtime time.Time, voucher_validdays int, voucher_isinvalid int8, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【抵用券管理】信息
	gwVoucherInfo, err := dao.GwVoucherDao.Get(voucherid)
	if err != nil {
		return false, errors.New("查询【抵用券管理】信息出错:" + err.Error())
	}
	if gwVoucherInfo.Voucherid <= 0 {
		return false, errors.New("【抵用券管理】信息不存在")
	}
	if gwVoucherInfo.DeleteStatus != 1 {
		return false, errors.New("【抵用券管理】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if gwVoucherInfo.Voucherid == voucherid &&
		gwVoucherInfo.VoucherCode == voucher_code &&
		gwVoucherInfo.VoucherUserid == voucher_userid &&
		gwVoucherInfo.VoucherQuota == voucher_quota &&
		gwVoucherInfo.VoucherReduce == voucher_reduce &&
		gwVoucherInfo.VoucherCreatetime == voucher_createtime &&
		gwVoucherInfo.VoucherValiddays == voucher_validdays &&
		gwVoucherInfo.VoucherIsinvalid == voucher_isinvalid &&
		gwVoucherInfo.Adder == adder &&
		gwVoucherInfo.Addtime == addtime &&
		gwVoucherInfo.Moder == moder &&
		gwVoucherInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【抵用券管理】信息
	gwVoucherInfo.VoucherCode = voucher_code
	gwVoucherInfo.VoucherUserid = voucher_userid
	gwVoucherInfo.VoucherQuota = voucher_quota
	gwVoucherInfo.VoucherReduce = voucher_reduce
	gwVoucherInfo.VoucherCreatetime = voucher_createtime
	gwVoucherInfo.VoucherValiddays = voucher_validdays
	gwVoucherInfo.VoucherIsinvalid = voucher_isinvalid
	gwVoucherInfo.Adder = adder
	gwVoucherInfo.Addtime = addtime
	gwVoucherInfo.Moder = moder
	gwVoucherInfo.Modtime = modtime
	gwVoucherInfo.DeleteStatus = 1
	//endregion

	//region 修改【抵用券管理】信息
	return dao.GwVoucherDao.Update(&gwVoucherInfo)
	//endregion

}

//删除【抵用券管理】信息
func DeleteGwVoucherInfo(voucherid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【抵用券管理】信息
	{
		gwVoucherInfo, err := dao.GwVoucherDao.Get(voucherid)
		if err != nil {
			return false, errors.New("查询【抵用券管理】信息出错:" + err.Error())
		}
		if gwVoucherInfo.Voucherid <= 0 {
			return false, errors.New("没有找到【抵用券管理】信息")
		}
		if gwVoucherInfo.DeleteStatus != 1 {
			return false, errors.New("【抵用券管理】信息已被删除")
		}
	}
	//endregion

	//region 删除【抵用券管理】信息
	return dao.GwVoucherDao.Delete(voucherid)
	//endregion

}

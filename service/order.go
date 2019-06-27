package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【订单表】信息
func GetOrderInfo(orderid int) (order model.Order, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【订单表】信息
	order, err = dao.OrderDao.Get(orderid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【订单ID】批量获取【订单表】列表
func GetInOrderList(orderids []int) (orders []model.Order, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	orders, err = dao.OrderDao.GetIn(orderids)
	if err != nil {
		return
	}
	count = len(orders)
	//endregion

	return
}

// 分页查询【订单表】列表
func GetOrderList(pageIndex, pageSize int) (orders []model.Order, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.OrderDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	orders, err = dao.OrderDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【订单表】信息
func InsertOrderInfo(pkgcode string, ordercode string, buyer int, paytype int, totalprice float64, payprice float64, paytime time.Time, isinvalid bool, isused bool, isrefund bool, refundprice float64, refundtime time.Time, remark string, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//var err error
	isExist := false
	var orderInfo model.Order

	//region 构造【订单表】信息
	orderInfo.Pkgcode = pkgcode
	orderInfo.Ordercode = ordercode
	orderInfo.Buyer = buyer
	orderInfo.Paytype = paytype
	orderInfo.Totalprice = totalprice
	orderInfo.Payprice = payprice
	orderInfo.Paytime = paytime
	orderInfo.Isinvalid = isinvalid
	orderInfo.Isused = isused
	orderInfo.Isrefund = isrefund
	orderInfo.Refundprice = refundprice
	orderInfo.Refundtime = refundtime
	orderInfo.Remark = remark
	orderInfo.Adder = adder
	orderInfo.Addtime = addtime
	orderInfo.Moder = moder
	orderInfo.Modtime = modtime
	orderInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【订单表】信息
		isSuccess, err := dao.OrderDao.Update(&orderInfo)
		if err != nil {
			return -1, errors.New("插入【订单表】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【订单表】失败")
		}
		return orderInfo.Orderid, nil
		//endregion

	} else {

		//region 插入数据库
		orderid, err := dao.OrderDao.Insert(&orderInfo)
		if err != nil {
			return -1, errors.New("插入【订单表】出错:" + err.Error())
		}
		if orderid <= 0 {
			return -1, errors.New("插入【订单表】失败")
		}
		return int(orderid), nil
		//endregion

	}

}

//修改【订单表】信息
func UpdateOrderInfo(orderid int, pkgcode string, ordercode string, buyer int, paytype int, totalprice float64, payprice float64, paytime time.Time, isinvalid bool, isused bool, isrefund bool, refundprice float64, refundtime time.Time, remark string, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【订单表】信息
	orderInfo, err := dao.OrderDao.Get(orderid)
	if err != nil {
		return false, errors.New("查询【订单表】信息出错:" + err.Error())
	}
	if orderInfo.Orderid <= 0 {
		return false, errors.New("【订单表】信息不存在")
	}
	if orderInfo.DeleteStatus != 1 {
		return false, errors.New("【订单表】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if orderInfo.Orderid == orderid &&
		orderInfo.Pkgcode == pkgcode &&
		orderInfo.Ordercode == ordercode &&
		orderInfo.Buyer == buyer &&
		orderInfo.Paytype == paytype &&
		orderInfo.Totalprice == totalprice &&
		orderInfo.Payprice == payprice &&
		orderInfo.Paytime == paytime &&
		orderInfo.Isinvalid == isinvalid &&
		orderInfo.Isused == isused &&
		orderInfo.Isrefund == isrefund &&
		orderInfo.Refundprice == refundprice &&
		orderInfo.Refundtime == refundtime &&
		orderInfo.Remark == remark &&
		orderInfo.Adder == adder &&
		orderInfo.Addtime == addtime &&
		orderInfo.Moder == moder &&
		orderInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【订单表】信息
	orderInfo.Pkgcode = pkgcode
	orderInfo.Ordercode = ordercode
	orderInfo.Buyer = buyer
	orderInfo.Paytype = paytype
	orderInfo.Totalprice = totalprice
	orderInfo.Payprice = payprice
	orderInfo.Paytime = paytime
	orderInfo.Isinvalid = isinvalid
	orderInfo.Isused = isused
	orderInfo.Isrefund = isrefund
	orderInfo.Refundprice = refundprice
	orderInfo.Refundtime = refundtime
	orderInfo.Remark = remark
	orderInfo.Adder = adder
	orderInfo.Addtime = addtime
	orderInfo.Moder = moder
	orderInfo.Modtime = modtime
	orderInfo.DeleteStatus = 1
	//endregion

	//region 修改【订单表】信息
	return dao.OrderDao.Update(&orderInfo)
	//endregion

}

//删除【订单表】信息
func DeleteOrderInfo(orderid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【订单表】信息
	{
		orderInfo, err := dao.OrderDao.Get(orderid)
		if err != nil {
			return false, errors.New("查询【订单表】信息出错:" + err.Error())
		}
		if orderInfo.Orderid <= 0 {
			return false, errors.New("没有找到【订单表】信息")
		}
		if orderInfo.DeleteStatus != 1 {
			return false, errors.New("【订单表】信息已被删除")
		}
	}
	//endregion

	//region 删除【订单表】信息
	return dao.OrderDao.Delete(orderid)
	//endregion

}

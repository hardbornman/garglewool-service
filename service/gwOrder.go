package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【订单管理】信息
func GetGwOrderInfo(orderid int) (gwOrder model.GwOrder, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【订单管理】信息
	gwOrder, err = dao.GwOrderDao.Get(orderid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【订单ID】批量获取【订单管理】列表
func GetInGwOrderList(orderids []int) (gwOrders []model.GwOrder, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	gwOrders, err = dao.GwOrderDao.GetIn(orderids)
	if err != nil {
		return
	}
	count = len(gwOrders)
	//endregion

	return
}

// 分页查询【订单管理】列表
func GetGwOrderList(pageIndex, pageSize int) (gwOrders []model.GwOrder, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.GwOrderDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	gwOrders, err = dao.GwOrderDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【订单管理】信息
func InsertGwOrderInfo(pkg_code string, order_code string, order_buyer int, order_paytype int, order_totalprice float64, order_payprice float64, order_paytime time.Time, order_isinvalid int8, order_isused int8, order_isrefund int8, order_refundprice float64, order_refundtime time.Time, order_remark string, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//var err error
	isExist := false
	var gwOrderInfo model.GwOrder

	//region 构造【订单管理】信息
	gwOrderInfo.PkgCode = pkg_code
	gwOrderInfo.OrderCode = order_code
	gwOrderInfo.OrderBuyer = order_buyer
	gwOrderInfo.OrderPaytype = order_paytype
	gwOrderInfo.OrderTotalprice = order_totalprice
	gwOrderInfo.OrderPayprice = order_payprice
	gwOrderInfo.OrderPaytime = order_paytime
	gwOrderInfo.OrderIsinvalid = order_isinvalid
	gwOrderInfo.OrderIsused = order_isused
	gwOrderInfo.OrderIsrefund = order_isrefund
	gwOrderInfo.OrderRefundprice = order_refundprice
	gwOrderInfo.OrderRefundtime = order_refundtime
	gwOrderInfo.OrderRemark = order_remark
	gwOrderInfo.Adder = adder
	gwOrderInfo.Addtime = addtime
	gwOrderInfo.Moder = moder
	gwOrderInfo.Modtime = modtime
	gwOrderInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【订单管理】信息
		isSuccess, err := dao.GwOrderDao.Update(&gwOrderInfo)
		if err != nil {
			return -1, errors.New("插入【订单管理】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【订单管理】失败")
		}
		return gwOrderInfo.Orderid, nil
		//endregion

	} else {

		//region 插入数据库
		orderid, err := dao.GwOrderDao.Insert(&gwOrderInfo)
		if err != nil {
			return -1, errors.New("插入【订单管理】出错:" + err.Error())
		}
		if orderid <= 0 {
			return -1, errors.New("插入【订单管理】失败")
		}
		return int(orderid), nil
		//endregion

	}

}

//修改【订单管理】信息
func UpdateGwOrderInfo(orderid int, pkg_code string, order_code string, order_buyer int, order_paytype int, order_totalprice float64, order_payprice float64, order_paytime time.Time, order_isinvalid int8, order_isused int8, order_isrefund int8, order_refundprice float64, order_refundtime time.Time, order_remark string, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【订单管理】信息
	gwOrderInfo, err := dao.GwOrderDao.Get(orderid)
	if err != nil {
		return false, errors.New("查询【订单管理】信息出错:" + err.Error())
	}
	if gwOrderInfo.Orderid <= 0 {
		return false, errors.New("【订单管理】信息不存在")
	}
	if gwOrderInfo.DeleteStatus != 1 {
		return false, errors.New("【订单管理】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if gwOrderInfo.Orderid == orderid &&
		gwOrderInfo.PkgCode == pkg_code &&
		gwOrderInfo.OrderCode == order_code &&
		gwOrderInfo.OrderBuyer == order_buyer &&
		gwOrderInfo.OrderPaytype == order_paytype &&
		gwOrderInfo.OrderTotalprice == order_totalprice &&
		gwOrderInfo.OrderPayprice == order_payprice &&
		gwOrderInfo.OrderPaytime == order_paytime &&
		gwOrderInfo.OrderIsinvalid == order_isinvalid &&
		gwOrderInfo.OrderIsused == order_isused &&
		gwOrderInfo.OrderIsrefund == order_isrefund &&
		gwOrderInfo.OrderRefundprice == order_refundprice &&
		gwOrderInfo.OrderRefundtime == order_refundtime &&
		gwOrderInfo.OrderRemark == order_remark &&
		gwOrderInfo.Adder == adder &&
		gwOrderInfo.Addtime == addtime &&
		gwOrderInfo.Moder == moder &&
		gwOrderInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【订单管理】信息
	gwOrderInfo.PkgCode = pkg_code
	gwOrderInfo.OrderCode = order_code
	gwOrderInfo.OrderBuyer = order_buyer
	gwOrderInfo.OrderPaytype = order_paytype
	gwOrderInfo.OrderTotalprice = order_totalprice
	gwOrderInfo.OrderPayprice = order_payprice
	gwOrderInfo.OrderPaytime = order_paytime
	gwOrderInfo.OrderIsinvalid = order_isinvalid
	gwOrderInfo.OrderIsused = order_isused
	gwOrderInfo.OrderIsrefund = order_isrefund
	gwOrderInfo.OrderRefundprice = order_refundprice
	gwOrderInfo.OrderRefundtime = order_refundtime
	gwOrderInfo.OrderRemark = order_remark
	gwOrderInfo.Adder = adder
	gwOrderInfo.Addtime = addtime
	gwOrderInfo.Moder = moder
	gwOrderInfo.Modtime = modtime
	gwOrderInfo.DeleteStatus = 1
	//endregion

	//region 修改【订单管理】信息
	return dao.GwOrderDao.Update(&gwOrderInfo)
	//endregion

}

//删除【订单管理】信息
func DeleteGwOrderInfo(orderid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【订单管理】信息
	{
		gwOrderInfo, err := dao.GwOrderDao.Get(orderid)
		if err != nil {
			return false, errors.New("查询【订单管理】信息出错:" + err.Error())
		}
		if gwOrderInfo.Orderid <= 0 {
			return false, errors.New("没有找到【订单管理】信息")
		}
		if gwOrderInfo.DeleteStatus != 1 {
			return false, errors.New("【订单管理】信息已被删除")
		}
	}
	//endregion

	//region 删除【订单管理】信息
	return dao.GwOrderDao.Delete(orderid)
	//endregion

}

package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【店铺管理】信息
func GetGwShopInfo(shopid int) (gwShop model.GwShop, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【店铺管理】信息
	gwShop, err = dao.GwShopDao.Get(shopid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【店铺ID】批量获取【店铺管理】列表
func GetInGwShopList(shopids []int) (gwShops []model.GwShop, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	gwShops, err = dao.GwShopDao.GetIn(shopids)
	if err != nil {
		return
	}
	count = len(gwShops)
	//endregion

	return
}

// 分页查询【店铺管理】列表
func GetGwShopList(pageIndex, pageSize int) (gwShops []model.GwShop, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.GwShopDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	gwShops, err = dao.GwShopDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【店铺管理】信息
func InsertGwShopInfo(shop_code string, shop_name string, shop_province string, shop_city string, shop_district string, shop_address string, shop_phone string, shop_createtime time.Time, shop_exittime time.Time, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//var err error
	isExist := false
	var gwShopInfo model.GwShop

	//region 构造【店铺管理】信息
	gwShopInfo.ShopCode = shop_code
	gwShopInfo.ShopName = shop_name
	gwShopInfo.ShopProvince = shop_province
	gwShopInfo.ShopCity = shop_city
	gwShopInfo.ShopDistrict = shop_district
	gwShopInfo.ShopAddress = shop_address
	gwShopInfo.ShopPhone = shop_phone
	gwShopInfo.ShopCreatetime = shop_createtime
	gwShopInfo.ShopExittime = shop_exittime
	gwShopInfo.Adder = adder
	gwShopInfo.Addtime = addtime
	gwShopInfo.Moder = moder
	gwShopInfo.Modtime = modtime
	gwShopInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【店铺管理】信息
		isSuccess, err := dao.GwShopDao.Update(&gwShopInfo)
		if err != nil {
			return -1, errors.New("插入【店铺管理】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【店铺管理】失败")
		}
		return gwShopInfo.Shopid, nil
		//endregion

	} else {

		//region 插入数据库
		shopid, err := dao.GwShopDao.Insert(&gwShopInfo)
		if err != nil {
			return -1, errors.New("插入【店铺管理】出错:" + err.Error())
		}
		if shopid <= 0 {
			return -1, errors.New("插入【店铺管理】失败")
		}
		return int(shopid), nil
		//endregion

	}

}

//修改【店铺管理】信息
func UpdateGwShopInfo(shopid int, shop_code string, shop_name string, shop_province string, shop_city string, shop_district string, shop_address string, shop_phone string, shop_createtime time.Time, shop_exittime time.Time, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【店铺管理】信息
	gwShopInfo, err := dao.GwShopDao.Get(shopid)
	if err != nil {
		return false, errors.New("查询【店铺管理】信息出错:" + err.Error())
	}
	if gwShopInfo.Shopid <= 0 {
		return false, errors.New("【店铺管理】信息不存在")
	}
	if gwShopInfo.DeleteStatus != 1 {
		return false, errors.New("【店铺管理】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if gwShopInfo.Shopid == shopid &&
		gwShopInfo.ShopCode == shop_code &&
		gwShopInfo.ShopName == shop_name &&
		gwShopInfo.ShopProvince == shop_province &&
		gwShopInfo.ShopCity == shop_city &&
		gwShopInfo.ShopDistrict == shop_district &&
		gwShopInfo.ShopAddress == shop_address &&
		gwShopInfo.ShopPhone == shop_phone &&
		gwShopInfo.ShopCreatetime == shop_createtime &&
		gwShopInfo.ShopExittime == shop_exittime &&
		gwShopInfo.Adder == adder &&
		gwShopInfo.Addtime == addtime &&
		gwShopInfo.Moder == moder &&
		gwShopInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【店铺管理】信息
	gwShopInfo.ShopCode = shop_code
	gwShopInfo.ShopName = shop_name
	gwShopInfo.ShopProvince = shop_province
	gwShopInfo.ShopCity = shop_city
	gwShopInfo.ShopDistrict = shop_district
	gwShopInfo.ShopAddress = shop_address
	gwShopInfo.ShopPhone = shop_phone
	gwShopInfo.ShopCreatetime = shop_createtime
	gwShopInfo.ShopExittime = shop_exittime
	gwShopInfo.Adder = adder
	gwShopInfo.Addtime = addtime
	gwShopInfo.Moder = moder
	gwShopInfo.Modtime = modtime
	gwShopInfo.DeleteStatus = 1
	//endregion

	//region 修改【店铺管理】信息
	return dao.GwShopDao.Update(&gwShopInfo)
	//endregion

}

//删除【店铺管理】信息
func DeleteGwShopInfo(shopid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【店铺管理】信息
	{
		gwShopInfo, err := dao.GwShopDao.Get(shopid)
		if err != nil {
			return false, errors.New("查询【店铺管理】信息出错:" + err.Error())
		}
		if gwShopInfo.Shopid <= 0 {
			return false, errors.New("没有找到【店铺管理】信息")
		}
		if gwShopInfo.DeleteStatus != 1 {
			return false, errors.New("【店铺管理】信息已被删除")
		}
	}
	//endregion

	//region 删除【店铺管理】信息
	return dao.GwShopDao.Delete(shopid)
	//endregion

}

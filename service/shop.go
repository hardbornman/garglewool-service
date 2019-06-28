package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【店铺表】信息
func GetShopInfo(shopid int) (shop model.Shop, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【店铺表】信息
	shop, err = dao.ShopDao.Get(shopid)
	if err != nil {
		return
	}
	//endregion

	//region 获取【商家用户表】信息
	if shop.Merchantid > 0 {
		var merchant model.Merchant
		merchant, err = dao.MerchantDao.Get(shop.Merchantid)
		if err != nil {
			return
		}
		shop.Merchantname = merchant.Merchantname
	}
	//endregion

	return
}

// 根据【店铺ID】批量获取【店铺表】列表
func GetInShopList(shopids []int) (shops []model.Shop, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	shops, err = dao.ShopDao.GetIn(shopids)
	if err != nil {
		return
	}
	count = len(shops)
	if len(shops) > 0 && shops[0].Shopid > 0 {

		//region 查询外键表【商家用户表】列表
		merchantids := make([]int, 0)
		var merchants []model.Merchant
		for _, shop := range shops {
			if shop.Merchantid > 0 {
				merchantids = append(merchantids, shop.Merchantid)
			}
		}
		if len(merchantids) > 0 {
			merchants, err = dao.MerchantDao.GetIn(merchantids)
			if err != nil {
				return
			}
		}
		//endregion

		//region 处理外键表
		for i, shop := range shops {

			//region 处理【商家用户表】列表
			if len(merchants) > 0 {
				for _, merchant := range merchants {
					if shop.Merchantid == merchant.Merchantid {
						shops[i].Merchantname = merchant.Merchantname
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

// 分页查询【店铺表】列表
func GetShopList(pageIndex, pageSize int) (shops []model.Shop, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.ShopDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	shops, err = dao.ShopDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	if len(shops) > 0 && shops[0].Shopid > 0 {

		//region 查询外键表【商家用户表】列表
		merchantids := make([]int, 0)
		var merchants []model.Merchant
		for _, shop := range shops {
			if shop.Merchantid > 0 {
				merchantids = append(merchantids, shop.Merchantid)
			}
		}
		if len(merchantids) > 0 {
			merchants, err = dao.MerchantDao.GetIn(merchantids)
			if err != nil {
				return
			}
		}
		//endregion

		//region 处理外键表
		for i, shop := range shops {

			//region 处理【商家用户表】列表
			if len(merchants) > 0 {
				for _, merchant := range merchants {
					if shop.Merchantid == merchant.Merchantid {
						shops[i].Merchantname = merchant.Merchantname
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

// 插入【店铺表】信息
func InsertShopInfo(shopcode string, shopname string, province string, city string, district string, address string, phone string, leaguetime time.Time, exittime time.Time, adder int, addtime time.Time, moder int, modtime time.Time, merchantid int, longtitude string, latitude string) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	var err error
	isExist := false
	var shopInfo model.Shop

	//region 判断merchantid值是否存在
	merchant, err := dao.MerchantDao.Get(merchantid)
	if err != nil {
		return -1, err
	}
	if merchant.Merchantid <= 0 {
		return -1, errors.New("merchantid值不存在")
	}
	if merchant.DeleteStatus != 1 {
		return -1, errors.New("merchantid值已删除")
	}
	//endregion

	//region 构造【店铺表】信息
	shopInfo.Shopcode = shopcode
	shopInfo.Shopname = shopname
	shopInfo.Province = province
	shopInfo.City = city
	shopInfo.District = district
	shopInfo.Address = address
	shopInfo.Phone = phone
	shopInfo.Leaguetime = leaguetime
	shopInfo.Exittime = exittime
	shopInfo.Adder = adder
	shopInfo.Addtime = addtime
	shopInfo.Moder = moder
	shopInfo.Modtime = modtime
	shopInfo.Merchantid = merchantid
	shopInfo.Longtitude = longtitude
	shopInfo.Latitude = latitude
	shopInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【店铺表】信息
		isSuccess, err := dao.ShopDao.Update(&shopInfo)
		if err != nil {
			return -1, errors.New("插入【店铺表】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【店铺表】失败")
		}
		return shopInfo.Shopid, nil
		//endregion

	} else {

		//region 插入数据库
		shopid, err := dao.ShopDao.Insert(&shopInfo)
		if err != nil {
			return -1, errors.New("插入【店铺表】出错:" + err.Error())
		}
		if shopid <= 0 {
			return -1, errors.New("插入【店铺表】失败")
		}
		return int(shopid), nil
		//endregion

	}

}

//修改【店铺表】信息
func UpdateShopInfo(shopid int, shopcode string, shopname string, province string, city string, district string, address string, phone string, leaguetime time.Time, exittime time.Time, adder int, addtime time.Time, moder int, modtime time.Time, merchantid int, longtitude string, latitude string) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【店铺表】信息
	shopInfo, err := dao.ShopDao.Get(shopid)
	if err != nil {
		return false, errors.New("查询【店铺表】信息出错:" + err.Error())
	}
	if shopInfo.Shopid <= 0 {
		return false, errors.New("【店铺表】信息不存在")
	}
	if shopInfo.DeleteStatus != 1 {
		return false, errors.New("【店铺表】信息已被删除")
	}
	//endregion

	//region 判断merchantid值是否存在
	merchant, err := dao.MerchantDao.Get(merchantid)
	if err != nil {
		return false, err
	}
	if merchant.Merchantid <= 0 {
		return false, errors.New("merchantid值不存在")
	}
	if merchant.DeleteStatus != 1 {
		return false, errors.New("merchantid值已删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if shopInfo.Shopid == shopid &&
		shopInfo.Shopcode == shopcode &&
		shopInfo.Shopname == shopname &&
		shopInfo.Province == province &&
		shopInfo.City == city &&
		shopInfo.District == district &&
		shopInfo.Address == address &&
		shopInfo.Phone == phone &&
		shopInfo.Leaguetime == leaguetime &&
		shopInfo.Exittime == exittime &&
		shopInfo.Adder == adder &&
		shopInfo.Addtime == addtime &&
		shopInfo.Moder == moder &&
		shopInfo.Modtime == modtime &&
		shopInfo.Merchantid == merchantid &&
		shopInfo.Longtitude == longtitude &&
		shopInfo.Latitude == latitude {
		return true, nil
	}
	//endregion

	//region 构造【店铺表】信息
	shopInfo.Shopcode = shopcode
	shopInfo.Shopname = shopname
	shopInfo.Province = province
	shopInfo.City = city
	shopInfo.District = district
	shopInfo.Address = address
	shopInfo.Phone = phone
	shopInfo.Leaguetime = leaguetime
	shopInfo.Exittime = exittime
	shopInfo.Adder = adder
	shopInfo.Addtime = addtime
	shopInfo.Moder = moder
	shopInfo.Modtime = modtime
	shopInfo.Merchantid = merchantid
	shopInfo.Longtitude = longtitude
	shopInfo.Latitude = latitude
	shopInfo.DeleteStatus = 1
	//endregion

	//region 修改【店铺表】信息
	return dao.ShopDao.Update(&shopInfo)
	//endregion

}

//删除【店铺表】信息
func DeleteShopInfo(shopid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 根据【店铺ID】查询【评论表】总数
	{
		count, err := dao.CommentDao.GetRowCountByShopid(shopid)
		if err != nil {
			return false, errors.New("查询【评论表】总数出错:" + err.Error())
		}
		if count > 0 {
			return false, errors.New("【评论表】存在相关记录")
		}
	}
	//endregion

	//region 根据【店铺ID】查询【套餐表】总数
	{
		count, err := dao.SetmealDao.GetRowCountByShopid(shopid)
		if err != nil {
			return false, errors.New("查询【套餐表】总数出错:" + err.Error())
		}
		if count > 0 {
			return false, errors.New("【套餐表】存在相关记录")
		}
	}
	//endregion

	//region 查询【店铺表】信息
	{
		shopInfo, err := dao.ShopDao.Get(shopid)
		if err != nil {
			return false, errors.New("查询【店铺表】信息出错:" + err.Error())
		}
		if shopInfo.Shopid <= 0 {
			return false, errors.New("没有找到【店铺表】信息")
		}
		if shopInfo.DeleteStatus != 1 {
			return false, errors.New("【店铺表】信息已被删除")
		}
	}
	//endregion

	//region 删除【店铺表】信息
	return dao.ShopDao.Delete(shopid)
	//endregion

}

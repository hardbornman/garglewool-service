package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【套餐表】信息
func GetSetmealInfo(setmealid int) (setmeal model.Setmeal, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【套餐表】信息
	setmeal, err = dao.SetmealDao.Get(setmealid)
	if err != nil {
		return
	}
	//endregion

	//region 获取【店铺表】信息
	if setmeal.Shopid > 0 {
		var shop model.Shop
		shop, err = dao.ShopDao.Get(setmeal.Shopid)
		if err != nil {
			return
		}
		setmeal.Shopname = shop.Shopname
	}
	//endregion

	return
}

// 根据【套餐ID】批量获取【套餐表】列表
func GetInSetmealList(setmealids []int) (setmeals []model.Setmeal, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	setmeals, err = dao.SetmealDao.GetIn(setmealids)
	if err != nil {
		return
	}
	count = len(setmeals)
	if len(setmeals) > 0 && setmeals[0].Setmealid > 0 {

		//region 查询外键表【店铺表】列表
		shopids := make([]int, 0)
		var shops []model.Shop
		for _, setmeal := range setmeals {
			if setmeal.Shopid > 0 {
				shopids = append(shopids, setmeal.Shopid)
			}
		}
		if len(shopids) > 0 {
			shops, err = dao.ShopDao.GetIn(shopids)
			if err != nil {
				return
			}
		}
		//endregion

		//region 处理外键表
		for i, setmeal := range setmeals {

			//region 处理【店铺表】列表
			if len(shops) > 0 {
				for _, shop := range shops {
					if setmeal.Shopid == shop.Shopid {
						setmeals[i].Shopname = shop.Shopname
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

// 分页查询【套餐表】列表
func GetSetmealList(pageIndex, pageSize int) (setmeals []model.Setmeal, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.SetmealDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	setmeals, err = dao.SetmealDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	if len(setmeals) > 0 && setmeals[0].Setmealid > 0 {

		//region 查询外键表【店铺表】列表
		shopids := make([]int, 0)
		var shops []model.Shop
		for _, setmeal := range setmeals {
			if setmeal.Shopid > 0 {
				shopids = append(shopids, setmeal.Shopid)
			}
		}
		if len(shopids) > 0 {
			shops, err = dao.ShopDao.GetIn(shopids)
			if err != nil {
				return
			}
		}
		//endregion

		//region 处理外键表
		for i, setmeal := range setmeals {

			//region 处理【店铺表】列表
			if len(shops) > 0 {
				for _, shop := range shops {
					if setmeal.Shopid == shop.Shopid {
						setmeals[i].Shopname = shop.Shopname
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

// 插入【套餐表】信息
func InsertSetmealInfo(shopid int, pkgcode string, setmealtype string, title string, people int, isorder bool, isrefund bool, isinhouse bool, isnew bool, isrecommend bool, validdays int, watchers int, validtime time.Time, exittime time.Time, links string, info string, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	var err error
	isExist := false
	var setmealInfo model.Setmeal

	//region 判断shopid值是否存在
	shop, err := dao.ShopDao.Get(shopid)
	if err != nil {
		return -1, err
	}
	if shop.Shopid <= 0 {
		return -1, errors.New("shopid值不存在")
	}
	if shop.DeleteStatus != 1 {
		return -1, errors.New("shopid值已删除")
	}
	//endregion

	//region 构造【套餐表】信息
	setmealInfo.Shopid = shopid
	setmealInfo.Pkgcode = pkgcode
	setmealInfo.Setmealtype = setmealtype
	setmealInfo.Title = title
	setmealInfo.People = people
	setmealInfo.Isorder = isorder
	setmealInfo.Isrefund = isrefund
	setmealInfo.Isinhouse = isinhouse
	setmealInfo.Isnew = isnew
	setmealInfo.Isrecommend = isrecommend
	setmealInfo.Validdays = validdays
	setmealInfo.Watchers = watchers
	setmealInfo.Validtime = validtime
	setmealInfo.Exittime = exittime
	setmealInfo.Links = links
	setmealInfo.Info = info
	setmealInfo.Adder = adder
	setmealInfo.Addtime = addtime
	setmealInfo.Moder = moder
	setmealInfo.Modtime = modtime
	setmealInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【套餐表】信息
		isSuccess, err := dao.SetmealDao.Update(&setmealInfo)
		if err != nil {
			return -1, errors.New("插入【套餐表】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【套餐表】失败")
		}
		return setmealInfo.Setmealid, nil
		//endregion

	} else {

		//region 插入数据库
		setmealid, err := dao.SetmealDao.Insert(&setmealInfo)
		if err != nil {
			return -1, errors.New("插入【套餐表】出错:" + err.Error())
		}
		if setmealid <= 0 {
			return -1, errors.New("插入【套餐表】失败")
		}
		return int(setmealid), nil
		//endregion

	}

}

//修改【套餐表】信息
func UpdateSetmealInfo(setmealid int, shopid int, pkgcode string, setmealtype string, title string, people int, isorder bool, isrefund bool, isinhouse bool, isnew bool, isrecommend bool, validdays int, watchers int, validtime time.Time, exittime time.Time, links string, info string, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【套餐表】信息
	setmealInfo, err := dao.SetmealDao.Get(setmealid)
	if err != nil {
		return false, errors.New("查询【套餐表】信息出错:" + err.Error())
	}
	if setmealInfo.Setmealid <= 0 {
		return false, errors.New("【套餐表】信息不存在")
	}
	if setmealInfo.DeleteStatus != 1 {
		return false, errors.New("【套餐表】信息已被删除")
	}
	//endregion

	//region 判断shopid值是否存在
	shop, err := dao.ShopDao.Get(shopid)
	if err != nil {
		return false, err
	}
	if shop.Shopid <= 0 {
		return false, errors.New("shopid值不存在")
	}
	if shop.DeleteStatus != 1 {
		return false, errors.New("shopid值已删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if setmealInfo.Setmealid == setmealid &&
		setmealInfo.Shopid == shopid &&
		setmealInfo.Pkgcode == pkgcode &&
		setmealInfo.Setmealtype == setmealtype &&
		setmealInfo.Title == title &&
		setmealInfo.People == people &&
		setmealInfo.Isorder == isorder &&
		setmealInfo.Isrefund == isrefund &&
		setmealInfo.Isinhouse == isinhouse &&
		setmealInfo.Isnew == isnew &&
		setmealInfo.Isrecommend == isrecommend &&
		setmealInfo.Validdays == validdays &&
		setmealInfo.Watchers == watchers &&
		setmealInfo.Validtime == validtime &&
		setmealInfo.Exittime == exittime &&
		setmealInfo.Links == links &&
		setmealInfo.Info == info &&
		setmealInfo.Adder == adder &&
		setmealInfo.Addtime == addtime &&
		setmealInfo.Moder == moder &&
		setmealInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【套餐表】信息
	setmealInfo.Shopid = shopid
	setmealInfo.Pkgcode = pkgcode
	setmealInfo.Setmealtype = setmealtype
	setmealInfo.Title = title
	setmealInfo.People = people
	setmealInfo.Isorder = isorder
	setmealInfo.Isrefund = isrefund
	setmealInfo.Isinhouse = isinhouse
	setmealInfo.Isnew = isnew
	setmealInfo.Isrecommend = isrecommend
	setmealInfo.Validdays = validdays
	setmealInfo.Watchers = watchers
	setmealInfo.Validtime = validtime
	setmealInfo.Exittime = exittime
	setmealInfo.Links = links
	setmealInfo.Info = info
	setmealInfo.Adder = adder
	setmealInfo.Addtime = addtime
	setmealInfo.Moder = moder
	setmealInfo.Modtime = modtime
	setmealInfo.DeleteStatus = 1
	//endregion

	//region 修改【套餐表】信息
	return dao.SetmealDao.Update(&setmealInfo)
	//endregion

}

//删除【套餐表】信息
func DeleteSetmealInfo(setmealid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 根据【套餐ID】查询【套餐明细表】总数
	{
		count, err := dao.SetmealdetailDao.GetRowCountBySetmealid(setmealid)
		if err != nil {
			return false, errors.New("查询【套餐明细表】总数出错:" + err.Error())
		}
		if count > 0 {
			return false, errors.New("【套餐明细表】存在相关记录")
		}
	}
	//endregion

	//region 查询【套餐表】信息
	{
		setmealInfo, err := dao.SetmealDao.Get(setmealid)
		if err != nil {
			return false, errors.New("查询【套餐表】信息出错:" + err.Error())
		}
		if setmealInfo.Setmealid <= 0 {
			return false, errors.New("没有找到【套餐表】信息")
		}
		if setmealInfo.DeleteStatus != 1 {
			return false, errors.New("【套餐表】信息已被删除")
		}
	}
	//endregion

	//region 删除【套餐表】信息
	return dao.SetmealDao.Delete(setmealid)
	//endregion

}

package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【套餐管理】信息
func GetGwPackageInfo(packageid int) (gwPackage model.GwPackage, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【套餐管理】信息
	gwPackage, err = dao.GwPackageDao.Get(packageid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【套餐ID】批量获取【套餐管理】列表
func GetInGwPackageList(packageids []int) (gwPackages []model.GwPackage, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	gwPackages, err = dao.GwPackageDao.GetIn(packageids)
	if err != nil {
		return
	}
	count = len(gwPackages)
	//endregion

	return
}

// 分页查询【套餐管理】列表
func GetGwPackageList(pageIndex, pageSize int) (gwPackages []model.GwPackage, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.GwPackageDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	gwPackages, err = dao.GwPackageDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【套餐管理】信息
func InsertGwPackageInfo(shop_code string, pkg_code string, pkg_type string, pkg_title string, pkg_people int, pkg_isorder int8, pkg_isrefund int8, pkg_isinhouse int8, pkg_isnew int8, pkg_isrecommend int8, pkg_validdays int, pkg_follows int, pkg_createtime time.Time, pkg_validtime time.Time, pkg_exittime time.Time, pkg_links string, pkg_info string, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//var err error
	isExist := false
	var gwPackageInfo model.GwPackage

	//region 构造【套餐管理】信息
	gwPackageInfo.ShopCode = shop_code
	gwPackageInfo.PkgCode = pkg_code
	gwPackageInfo.PkgType = pkg_type
	gwPackageInfo.PkgTitle = pkg_title
	gwPackageInfo.PkgPeople = pkg_people
	gwPackageInfo.PkgIsorder = pkg_isorder
	gwPackageInfo.PkgIsrefund = pkg_isrefund
	gwPackageInfo.PkgIsinhouse = pkg_isinhouse
	gwPackageInfo.PkgIsnew = pkg_isnew
	gwPackageInfo.PkgIsrecommend = pkg_isrecommend
	gwPackageInfo.PkgValiddays = pkg_validdays
	gwPackageInfo.PkgFollows = pkg_follows
	gwPackageInfo.PkgCreatetime = pkg_createtime
	gwPackageInfo.PkgValidtime = pkg_validtime
	gwPackageInfo.PkgExittime = pkg_exittime
	gwPackageInfo.PkgLinks = pkg_links
	gwPackageInfo.PkgInfo = pkg_info
	gwPackageInfo.Adder = adder
	gwPackageInfo.Addtime = addtime
	gwPackageInfo.Moder = moder
	gwPackageInfo.Modtime = modtime
	gwPackageInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【套餐管理】信息
		isSuccess, err := dao.GwPackageDao.Update(&gwPackageInfo)
		if err != nil {
			return -1, errors.New("插入【套餐管理】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【套餐管理】失败")
		}
		return gwPackageInfo.Packageid, nil
		//endregion

	} else {

		//region 插入数据库
		packageid, err := dao.GwPackageDao.Insert(&gwPackageInfo)
		if err != nil {
			return -1, errors.New("插入【套餐管理】出错:" + err.Error())
		}
		if packageid <= 0 {
			return -1, errors.New("插入【套餐管理】失败")
		}
		return int(packageid), nil
		//endregion

	}

}

//修改【套餐管理】信息
func UpdateGwPackageInfo(packageid int, shop_code string, pkg_code string, pkg_type string, pkg_title string, pkg_people int, pkg_isorder int8, pkg_isrefund int8, pkg_isinhouse int8, pkg_isnew int8, pkg_isrecommend int8, pkg_validdays int, pkg_follows int, pkg_createtime time.Time, pkg_validtime time.Time, pkg_exittime time.Time, pkg_links string, pkg_info string, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【套餐管理】信息
	gwPackageInfo, err := dao.GwPackageDao.Get(packageid)
	if err != nil {
		return false, errors.New("查询【套餐管理】信息出错:" + err.Error())
	}
	if gwPackageInfo.Packageid <= 0 {
		return false, errors.New("【套餐管理】信息不存在")
	}
	if gwPackageInfo.DeleteStatus != 1 {
		return false, errors.New("【套餐管理】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if gwPackageInfo.Packageid == packageid &&
		gwPackageInfo.ShopCode == shop_code &&
		gwPackageInfo.PkgCode == pkg_code &&
		gwPackageInfo.PkgType == pkg_type &&
		gwPackageInfo.PkgTitle == pkg_title &&
		gwPackageInfo.PkgPeople == pkg_people &&
		gwPackageInfo.PkgIsorder == pkg_isorder &&
		gwPackageInfo.PkgIsrefund == pkg_isrefund &&
		gwPackageInfo.PkgIsinhouse == pkg_isinhouse &&
		gwPackageInfo.PkgIsnew == pkg_isnew &&
		gwPackageInfo.PkgIsrecommend == pkg_isrecommend &&
		gwPackageInfo.PkgValiddays == pkg_validdays &&
		gwPackageInfo.PkgFollows == pkg_follows &&
		gwPackageInfo.PkgCreatetime == pkg_createtime &&
		gwPackageInfo.PkgValidtime == pkg_validtime &&
		gwPackageInfo.PkgExittime == pkg_exittime &&
		gwPackageInfo.PkgLinks == pkg_links &&
		gwPackageInfo.PkgInfo == pkg_info &&
		gwPackageInfo.Adder == adder &&
		gwPackageInfo.Addtime == addtime &&
		gwPackageInfo.Moder == moder &&
		gwPackageInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【套餐管理】信息
	gwPackageInfo.ShopCode = shop_code
	gwPackageInfo.PkgCode = pkg_code
	gwPackageInfo.PkgType = pkg_type
	gwPackageInfo.PkgTitle = pkg_title
	gwPackageInfo.PkgPeople = pkg_people
	gwPackageInfo.PkgIsorder = pkg_isorder
	gwPackageInfo.PkgIsrefund = pkg_isrefund
	gwPackageInfo.PkgIsinhouse = pkg_isinhouse
	gwPackageInfo.PkgIsnew = pkg_isnew
	gwPackageInfo.PkgIsrecommend = pkg_isrecommend
	gwPackageInfo.PkgValiddays = pkg_validdays
	gwPackageInfo.PkgFollows = pkg_follows
	gwPackageInfo.PkgCreatetime = pkg_createtime
	gwPackageInfo.PkgValidtime = pkg_validtime
	gwPackageInfo.PkgExittime = pkg_exittime
	gwPackageInfo.PkgLinks = pkg_links
	gwPackageInfo.PkgInfo = pkg_info
	gwPackageInfo.Adder = adder
	gwPackageInfo.Addtime = addtime
	gwPackageInfo.Moder = moder
	gwPackageInfo.Modtime = modtime
	gwPackageInfo.DeleteStatus = 1
	//endregion

	//region 修改【套餐管理】信息
	return dao.GwPackageDao.Update(&gwPackageInfo)
	//endregion

}

//删除【套餐管理】信息
func DeleteGwPackageInfo(packageid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【套餐管理】信息
	{
		gwPackageInfo, err := dao.GwPackageDao.Get(packageid)
		if err != nil {
			return false, errors.New("查询【套餐管理】信息出错:" + err.Error())
		}
		if gwPackageInfo.Packageid <= 0 {
			return false, errors.New("没有找到【套餐管理】信息")
		}
		if gwPackageInfo.DeleteStatus != 1 {
			return false, errors.New("【套餐管理】信息已被删除")
		}
	}
	//endregion

	//region 删除【套餐管理】信息
	return dao.GwPackageDao.Delete(packageid)
	//endregion

}

package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【套餐明细管理】信息
func GetGwPackagedetailInfo(packagedetailid int) (gwPackagedetail model.GwPackagedetail, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【套餐明细管理】信息
	gwPackagedetail, err = dao.GwPackagedetailDao.Get(packagedetailid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【套餐明细ID】批量获取【套餐明细管理】列表
func GetInGwPackagedetailList(packagedetailids []int) (gwPackagedetails []model.GwPackagedetail, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	gwPackagedetails, err = dao.GwPackagedetailDao.GetIn(packagedetailids)
	if err != nil {
		return
	}
	count = len(gwPackagedetails)
	//endregion

	return
}

// 分页查询【套餐明细管理】列表
func GetGwPackagedetailList(pageIndex, pageSize int) (gwPackagedetails []model.GwPackagedetail, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.GwPackagedetailDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	gwPackagedetails, err = dao.GwPackagedetailDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【套餐明细管理】信息
func InsertGwPackagedetailInfo(pkg_code string, pkgdetail_name string, pkgdetail_nums int, pkgdetail_price float64, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//var err error
	isExist := false
	var gwPackagedetailInfo model.GwPackagedetail

	//region 构造【套餐明细管理】信息
	gwPackagedetailInfo.PkgCode = pkg_code
	gwPackagedetailInfo.PkgdetailName = pkgdetail_name
	gwPackagedetailInfo.PkgdetailNums = pkgdetail_nums
	gwPackagedetailInfo.PkgdetailPrice = pkgdetail_price
	gwPackagedetailInfo.Adder = adder
	gwPackagedetailInfo.Addtime = addtime
	gwPackagedetailInfo.Moder = moder
	gwPackagedetailInfo.Modtime = modtime
	gwPackagedetailInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【套餐明细管理】信息
		isSuccess, err := dao.GwPackagedetailDao.Update(&gwPackagedetailInfo)
		if err != nil {
			return -1, errors.New("插入【套餐明细管理】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【套餐明细管理】失败")
		}
		return gwPackagedetailInfo.Packagedetailid, nil
		//endregion

	} else {

		//region 插入数据库
		packagedetailid, err := dao.GwPackagedetailDao.Insert(&gwPackagedetailInfo)
		if err != nil {
			return -1, errors.New("插入【套餐明细管理】出错:" + err.Error())
		}
		if packagedetailid <= 0 {
			return -1, errors.New("插入【套餐明细管理】失败")
		}
		return int(packagedetailid), nil
		//endregion

	}

}

//修改【套餐明细管理】信息
func UpdateGwPackagedetailInfo(packagedetailid int, pkg_code string, pkgdetail_name string, pkgdetail_nums int, pkgdetail_price float64, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【套餐明细管理】信息
	gwPackagedetailInfo, err := dao.GwPackagedetailDao.Get(packagedetailid)
	if err != nil {
		return false, errors.New("查询【套餐明细管理】信息出错:" + err.Error())
	}
	if gwPackagedetailInfo.Packagedetailid <= 0 {
		return false, errors.New("【套餐明细管理】信息不存在")
	}
	if gwPackagedetailInfo.DeleteStatus != 1 {
		return false, errors.New("【套餐明细管理】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if gwPackagedetailInfo.Packagedetailid == packagedetailid &&
		gwPackagedetailInfo.PkgCode == pkg_code &&
		gwPackagedetailInfo.PkgdetailName == pkgdetail_name &&
		gwPackagedetailInfo.PkgdetailNums == pkgdetail_nums &&
		gwPackagedetailInfo.PkgdetailPrice == pkgdetail_price &&
		gwPackagedetailInfo.Adder == adder &&
		gwPackagedetailInfo.Addtime == addtime &&
		gwPackagedetailInfo.Moder == moder &&
		gwPackagedetailInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【套餐明细管理】信息
	gwPackagedetailInfo.PkgCode = pkg_code
	gwPackagedetailInfo.PkgdetailName = pkgdetail_name
	gwPackagedetailInfo.PkgdetailNums = pkgdetail_nums
	gwPackagedetailInfo.PkgdetailPrice = pkgdetail_price
	gwPackagedetailInfo.Adder = adder
	gwPackagedetailInfo.Addtime = addtime
	gwPackagedetailInfo.Moder = moder
	gwPackagedetailInfo.Modtime = modtime
	gwPackagedetailInfo.DeleteStatus = 1
	//endregion

	//region 修改【套餐明细管理】信息
	return dao.GwPackagedetailDao.Update(&gwPackagedetailInfo)
	//endregion

}

//删除【套餐明细管理】信息
func DeleteGwPackagedetailInfo(packagedetailid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【套餐明细管理】信息
	{
		gwPackagedetailInfo, err := dao.GwPackagedetailDao.Get(packagedetailid)
		if err != nil {
			return false, errors.New("查询【套餐明细管理】信息出错:" + err.Error())
		}
		if gwPackagedetailInfo.Packagedetailid <= 0 {
			return false, errors.New("没有找到【套餐明细管理】信息")
		}
		if gwPackagedetailInfo.DeleteStatus != 1 {
			return false, errors.New("【套餐明细管理】信息已被删除")
		}
	}
	//endregion

	//region 删除【套餐明细管理】信息
	return dao.GwPackagedetailDao.Delete(packagedetailid)
	//endregion

}

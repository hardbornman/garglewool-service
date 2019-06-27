package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
)

// 获取【用户角色表】信息
func GetUserroleInfo(userroleid int) (userrole model.Userrole, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【用户角色表】信息
	userrole, err = dao.UserroleDao.Get(userroleid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【用户角色ID】批量获取【用户角色表】列表
func GetInUserroleList(userroleids []int) (userroles []model.Userrole, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	userroles, err = dao.UserroleDao.GetIn(userroleids)
	if err != nil {
		return
	}
	count = len(userroles)
	//endregion

	return
}

// 分页查询【用户角色表】列表
func GetUserroleList(pageIndex, pageSize int) (userroles []model.Userrole, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.UserroleDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	userroles, err = dao.UserroleDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【用户角色表】信息
func InsertUserroleInfo(rolename string, desc string, remark string) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//var err error
	isExist := false
	var userroleInfo model.Userrole

	//region 构造【用户角色表】信息
	userroleInfo.Rolename = rolename
	userroleInfo.Desc = desc
	userroleInfo.Remark = remark
	userroleInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【用户角色表】信息
		isSuccess, err := dao.UserroleDao.Update(&userroleInfo)
		if err != nil {
			return -1, errors.New("插入【用户角色表】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【用户角色表】失败")
		}
		return userroleInfo.Userroleid, nil
		//endregion

	} else {

		//region 插入数据库
		userroleid, err := dao.UserroleDao.Insert(&userroleInfo)
		if err != nil {
			return -1, errors.New("插入【用户角色表】出错:" + err.Error())
		}
		if userroleid <= 0 {
			return -1, errors.New("插入【用户角色表】失败")
		}
		return int(userroleid), nil
		//endregion

	}

}

//修改【用户角色表】信息
func UpdateUserroleInfo(userroleid int, rolename string, desc string, remark string) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【用户角色表】信息
	userroleInfo, err := dao.UserroleDao.Get(userroleid)
	if err != nil {
		return false, errors.New("查询【用户角色表】信息出错:" + err.Error())
	}
	if userroleInfo.Userroleid <= 0 {
		return false, errors.New("【用户角色表】信息不存在")
	}
	if userroleInfo.DeleteStatus != 1 {
		return false, errors.New("【用户角色表】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if userroleInfo.Userroleid == userroleid &&
		userroleInfo.Rolename == rolename &&
		userroleInfo.Desc == desc &&
		userroleInfo.Remark == remark {
		return true, nil
	}
	//endregion

	//region 构造【用户角色表】信息
	userroleInfo.Rolename = rolename
	userroleInfo.Desc = desc
	userroleInfo.Remark = remark
	userroleInfo.DeleteStatus = 1
	//endregion

	//region 修改【用户角色表】信息
	return dao.UserroleDao.Update(&userroleInfo)
	//endregion

}

//删除【用户角色表】信息
func DeleteUserroleInfo(userroleid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 根据【用户角色ID】查询【商家用户表】总数
	{
		count, err := dao.MerchantDao.GetRowCountByUserroleid(userroleid)
		if err != nil {
			return false, errors.New("查询【商家用户表】总数出错:" + err.Error())
		}
		if count > 0 {
			return false, errors.New("【商家用户表】存在相关记录")
		}
	}
	//endregion

	//region 查询【用户角色表】信息
	{
		userroleInfo, err := dao.UserroleDao.Get(userroleid)
		if err != nil {
			return false, errors.New("查询【用户角色表】信息出错:" + err.Error())
		}
		if userroleInfo.Userroleid <= 0 {
			return false, errors.New("没有找到【用户角色表】信息")
		}
		if userroleInfo.DeleteStatus != 1 {
			return false, errors.New("【用户角色表】信息已被删除")
		}
	}
	//endregion

	//region 删除【用户角色表】信息
	return dao.UserroleDao.Delete(userroleid)
	//endregion

}

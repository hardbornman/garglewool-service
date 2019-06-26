package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【用户管理】信息
func GetGwUserInfo(userid int) (gwUser model.GwUser, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【用户管理】信息
	gwUser, err = dao.GwUserDao.Get(userid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【用户ID】批量获取【用户管理】列表
func GetInGwUserList(userids []int) (gwUsers []model.GwUser, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	gwUsers, err = dao.GwUserDao.GetIn(userids)
	if err != nil {
		return
	}
	count = len(gwUsers)
	//endregion

	return
}

// 分页查询【用户管理】列表
func GetGwUserList(pageIndex, pageSize int) (gwUsers []model.GwUser, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.GwUserDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	gwUsers, err = dao.GwUserDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【用户管理】信息
func InsertGwUserInfo(user_name string, user_phone string, user_golds int, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//var err error
	isExist := false
	var gwUserInfo model.GwUser

	//region 构造【用户管理】信息
	gwUserInfo.UserName = user_name
	gwUserInfo.UserPhone = user_phone
	gwUserInfo.UserGolds = user_golds
	gwUserInfo.Adder = adder
	gwUserInfo.Addtime = addtime
	gwUserInfo.Moder = moder
	gwUserInfo.Modtime = modtime
	gwUserInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【用户管理】信息
		isSuccess, err := dao.GwUserDao.Update(&gwUserInfo)
		if err != nil {
			return -1, errors.New("插入【用户管理】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【用户管理】失败")
		}
		return gwUserInfo.Userid, nil
		//endregion

	} else {

		//region 插入数据库
		userid, err := dao.GwUserDao.Insert(&gwUserInfo)
		if err != nil {
			return -1, errors.New("插入【用户管理】出错:" + err.Error())
		}
		if userid <= 0 {
			return -1, errors.New("插入【用户管理】失败")
		}
		return int(userid), nil
		//endregion

	}

}

//修改【用户管理】信息
func UpdateGwUserInfo(userid int, user_name string, user_phone string, user_golds int, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【用户管理】信息
	gwUserInfo, err := dao.GwUserDao.Get(userid)
	if err != nil {
		return false, errors.New("查询【用户管理】信息出错:" + err.Error())
	}
	if gwUserInfo.Userid <= 0 {
		return false, errors.New("【用户管理】信息不存在")
	}
	if gwUserInfo.DeleteStatus != 1 {
		return false, errors.New("【用户管理】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if gwUserInfo.Userid == userid &&
		gwUserInfo.UserName == user_name &&
		gwUserInfo.UserPhone == user_phone &&
		gwUserInfo.UserGolds == user_golds &&
		gwUserInfo.Adder == adder &&
		gwUserInfo.Addtime == addtime &&
		gwUserInfo.Moder == moder &&
		gwUserInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【用户管理】信息
	gwUserInfo.UserName = user_name
	gwUserInfo.UserPhone = user_phone
	gwUserInfo.UserGolds = user_golds
	gwUserInfo.Adder = adder
	gwUserInfo.Addtime = addtime
	gwUserInfo.Moder = moder
	gwUserInfo.Modtime = modtime
	gwUserInfo.DeleteStatus = 1
	//endregion

	//region 修改【用户管理】信息
	return dao.GwUserDao.Update(&gwUserInfo)
	//endregion

}

//删除【用户管理】信息
func DeleteGwUserInfo(userid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【用户管理】信息
	{
		gwUserInfo, err := dao.GwUserDao.Get(userid)
		if err != nil {
			return false, errors.New("查询【用户管理】信息出错:" + err.Error())
		}
		if gwUserInfo.Userid <= 0 {
			return false, errors.New("没有找到【用户管理】信息")
		}
		if gwUserInfo.DeleteStatus != 1 {
			return false, errors.New("【用户管理】信息已被删除")
		}
	}
	//endregion

	//region 删除【用户管理】信息
	return dao.GwUserDao.Delete(userid)
	//endregion

}

// 登录【用户表】信息
func Login(loginName string, loginPwd string) (user model.GwUser, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【聊天用户表】信息
	user, err = dao.GwUserDao.Login(loginName, loginPwd)
	if err != nil {
		return
	}
	//endregion

	return
}

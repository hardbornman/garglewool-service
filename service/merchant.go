package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【商家用户表】信息
func GetMerchantInfo(merchantid int) (merchant model.Merchant, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【商家用户表】信息
	merchant, err = dao.MerchantDao.Get(merchantid)
	if err != nil {
		return
	}
	//endregion

	//region 获取【用户角色表】信息
	if merchant.Userroleid > 0 {
		var userrole model.Userrole
		userrole, err = dao.UserroleDao.Get(merchant.Userroleid)
		if err != nil {
			return
		}
		merchant.Rolename = userrole.Rolename
	}
	//endregion

	return
}

// 根据【商家ID】批量获取【商家用户表】列表
func GetInMerchantList(merchantids []int) (merchants []model.Merchant, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	merchants, err = dao.MerchantDao.GetIn(merchantids)
	if err != nil {
		return
	}
	count = len(merchants)
	if len(merchants) > 0 && merchants[0].Merchantid > 0 {

		//region 查询外键表【用户角色表】列表
		userroleids := make([]int, 0)
		var userroles []model.Userrole
		for _, merchant := range merchants {
			if merchant.Userroleid > 0 {
				userroleids = append(userroleids, merchant.Userroleid)
			}
		}
		if len(userroleids) > 0 {
			userroles, err = dao.UserroleDao.GetIn(userroleids)
			if err != nil {
				return
			}
		}
		//endregion

		//region 处理外键表
		for i, merchant := range merchants {

			//region 处理【用户角色表】列表
			if len(userroles) > 0 {
				for _, userrole := range userroles {
					if merchant.Userroleid == userrole.Userroleid {
						merchants[i].Rolename = userrole.Rolename
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

// 分页查询【商家用户表】列表
func GetMerchantList(pageIndex, pageSize int) (merchants []model.Merchant, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.MerchantDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	merchants, err = dao.MerchantDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	if len(merchants) > 0 && merchants[0].Merchantid > 0 {

		//region 查询外键表【用户角色表】列表
		userroleids := make([]int, 0)
		var userroles []model.Userrole
		for _, merchant := range merchants {
			if merchant.Userroleid > 0 {
				userroleids = append(userroleids, merchant.Userroleid)
			}
		}
		if len(userroleids) > 0 {
			userroles, err = dao.UserroleDao.GetIn(userroleids)
			if err != nil {
				return
			}
		}
		//endregion

		//region 处理外键表
		for i, merchant := range merchants {

			//region 处理【用户角色表】列表
			if len(userroles) > 0 {
				for _, userrole := range userroles {
					if merchant.Userroleid == userrole.Userroleid {
						merchants[i].Rolename = userrole.Rolename
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

// 插入【商家用户表】信息
func InsertMerchantInfo(merchantname string, phone string, userroleid int, loginpwd string, loginaccount string, nick string, wechataccount string, wechatsign string, remark string, lastlogintime time.Time, lastloginaddr string, rigstertime time.Time, enable bool, addr string) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	var err error
	isExist := false
	var merchantInfo model.Merchant

	//region 判断userroleid值是否存在
	userrole, err := dao.UserroleDao.Get(userroleid)
	if err != nil {
		return -1, err
	}
	if userrole.Userroleid <= 0 {
		return -1, errors.New("userroleid值不存在")
	}
	if userrole.DeleteStatus != 1 {
		return -1, errors.New("userroleid值已删除")
	}
	//endregion

	//region 构造【商家用户表】信息
	merchantInfo.Merchantname = merchantname
	merchantInfo.Phone = phone
	merchantInfo.Userroleid = userroleid
	merchantInfo.Loginaccount = loginaccount
	merchantInfo.Nick = nick
	merchantInfo.Wechataccount = wechataccount
	merchantInfo.Wechatsign = wechatsign
	merchantInfo.Remark = remark
	merchantInfo.Lastlogintime = lastlogintime
	merchantInfo.Lastloginaddr = lastloginaddr
	merchantInfo.Rigstertime = rigstertime
	merchantInfo.Enable = enable
	merchantInfo.Addr = addr
	merchantInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【商家用户表】信息
		isSuccess, err := dao.MerchantDao.Update(&merchantInfo)
		if err != nil {
			return -1, errors.New("插入【商家用户表】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【商家用户表】失败")
		}
		return merchantInfo.Merchantid, nil
		//endregion

	} else {

		//region 插入数据库
		merchantid, err := dao.MerchantDao.Insert(&merchantInfo)
		if err != nil {
			return -1, errors.New("插入【商家用户表】出错:" + err.Error())
		}
		if merchantid <= 0 {
			return -1, errors.New("插入【商家用户表】失败")
		}
		return int(merchantid), nil
		//endregion

	}

}

//修改【商家用户表】信息
func UpdateMerchantInfo(merchantid int, merchantname string, phone string, userroleid int, loginpwd string, loginaccount string, nick string, wechataccount string, wechatsign string, remark string, lastlogintime time.Time, lastloginaddr string, rigstertime time.Time, enable bool, addr string) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【商家用户表】信息
	merchantInfo, err := dao.MerchantDao.Get(merchantid)
	if err != nil {
		return false, errors.New("查询【商家用户表】信息出错:" + err.Error())
	}
	if merchantInfo.Merchantid <= 0 {
		return false, errors.New("【商家用户表】信息不存在")
	}
	if merchantInfo.DeleteStatus != 1 {
		return false, errors.New("【商家用户表】信息已被删除")
	}
	//endregion

	//region 判断userroleid值是否存在
	userrole, err := dao.UserroleDao.Get(userroleid)
	if err != nil {
		return false, err
	}
	if userrole.Userroleid <= 0 {
		return false, errors.New("userroleid值不存在")
	}
	if userrole.DeleteStatus != 1 {
		return false, errors.New("userroleid值已删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if merchantInfo.Merchantid == merchantid &&
		merchantInfo.Merchantname == merchantname &&
		merchantInfo.Phone == phone &&
		merchantInfo.Userroleid == userroleid &&
		merchantInfo.Loginpwd == loginpwd &&
		merchantInfo.Loginaccount == loginaccount &&
		merchantInfo.Nick == nick &&
		merchantInfo.Wechataccount == wechataccount &&
		merchantInfo.Wechatsign == wechatsign &&
		merchantInfo.Remark == remark &&
		merchantInfo.Lastlogintime == lastlogintime &&
		merchantInfo.Lastloginaddr == lastloginaddr &&
		merchantInfo.Rigstertime == rigstertime &&
		merchantInfo.Enable == enable &&
		merchantInfo.Addr == addr {
		return true, nil
	}
	//endregion

	//region 构造【商家用户表】信息
	merchantInfo.Merchantname = merchantname
	merchantInfo.Phone = phone
	merchantInfo.Userroleid = userroleid
	merchantInfo.Loginaccount = loginaccount
	merchantInfo.Nick = nick
	merchantInfo.Wechataccount = wechataccount
	merchantInfo.Wechatsign = wechatsign
	merchantInfo.Remark = remark
	merchantInfo.Lastlogintime = lastlogintime
	merchantInfo.Lastloginaddr = lastloginaddr
	merchantInfo.Rigstertime = rigstertime
	merchantInfo.Enable = enable
	merchantInfo.Addr = addr
	merchantInfo.DeleteStatus = 1
	//endregion

	//region 修改【商家用户表】信息
	return dao.MerchantDao.Update(&merchantInfo)
	//endregion

}

//删除【商家用户表】信息
func DeleteMerchantInfo(merchantid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 根据【商家ID】查询【店铺表】总数
	{
		count, err := dao.ShopDao.GetRowCountByMerchantid(merchantid)
		if err != nil {
			return false, errors.New("查询【店铺表】总数出错:" + err.Error())
		}
		if count > 0 {
			return false, errors.New("【店铺表】存在相关记录")
		}
	}
	//endregion

	//region 查询【商家用户表】信息
	{
		merchantInfo, err := dao.MerchantDao.Get(merchantid)
		if err != nil {
			return false, errors.New("查询【商家用户表】信息出错:" + err.Error())
		}
		if merchantInfo.Merchantid <= 0 {
			return false, errors.New("没有找到【商家用户表】信息")
		}
		if merchantInfo.DeleteStatus != 1 {
			return false, errors.New("【商家用户表】信息已被删除")
		}
	}
	//endregion

	//region 删除【商家用户表】信息
	return dao.MerchantDao.Delete(merchantid)
	//endregion

}

// 登录【用户表】信息
func Login(loginAccount string, loginPwd string) (merchant model.Merchant, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【聊天用户表】信息
	merchant, err = dao.MerchantDao.Login(loginAccount, loginPwd)
	if err != nil {
		return
	}
	//endregion

	return
}

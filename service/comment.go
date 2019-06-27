package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【评论表】信息
func GetCommentInfo(id int) (comment model.Comment, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【评论表】信息
	comment, err = dao.CommentDao.Get(id)
	if err != nil {
		return
	}
	//endregion

	//region 获取【店铺表】信息
	if comment.Shopid > 0 {
		var shop model.Shop
		shop, err = dao.ShopDao.Get(comment.Shopid)
		if err != nil {
			return
		}
		comment.Shopname = shop.Shopname
	}
	//endregion

	return
}

// 根据【评论ID】批量获取【评论表】列表
func GetInCommentList(ids []int) (comments []model.Comment, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	comments, err = dao.CommentDao.GetIn(ids)
	if err != nil {
		return
	}
	count = len(comments)
	if len(comments) > 0 && comments[0].Id > 0 {

		//region 查询外键表【店铺表】列表
		shopids := make([]int, 0)
		var shops []model.Shop
		for _, comment := range comments {
			if comment.Shopid > 0 {
				shopids = append(shopids, comment.Shopid)
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
		for i, comment := range comments {

			//region 处理【店铺表】列表
			if len(shops) > 0 {
				for _, shop := range shops {
					if comment.Shopid == shop.Shopid {
						comments[i].Shopname = shop.Shopname
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

// 分页查询【评论表】列表
func GetCommentList(pageIndex, pageSize int) (comments []model.Comment, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.CommentDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	comments, err = dao.CommentDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	if len(comments) > 0 && comments[0].Id > 0 {

		//region 查询外键表【店铺表】列表
		shopids := make([]int, 0)
		var shops []model.Shop
		for _, comment := range comments {
			if comment.Shopid > 0 {
				shopids = append(shopids, comment.Shopid)
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
		for i, comment := range comments {

			//region 处理【店铺表】列表
			if len(shops) > 0 {
				for _, shop := range shops {
					if comment.Shopid == shop.Shopid {
						comments[i].Shopname = shop.Shopname
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

// 插入【评论表】信息
func InsertCommentInfo(ordercode string, commentor int, star int, info string, images string, adder int, addtime time.Time, moder int, modtime time.Time, shopid int) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	var err error
	isExist := false
	var commentInfo model.Comment

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

	//region 构造【评论表】信息
	commentInfo.Ordercode = ordercode
	commentInfo.Commentor = commentor
	commentInfo.Star = star
	commentInfo.Info = info
	commentInfo.Images = images
	commentInfo.Adder = adder
	commentInfo.Addtime = addtime
	commentInfo.Moder = moder
	commentInfo.Modtime = modtime
	commentInfo.Shopid = shopid
	commentInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【评论表】信息
		isSuccess, err := dao.CommentDao.Update(&commentInfo)
		if err != nil {
			return -1, errors.New("插入【评论表】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【评论表】失败")
		}
		return commentInfo.Id, nil
		//endregion

	} else {

		//region 插入数据库
		id, err := dao.CommentDao.Insert(&commentInfo)
		if err != nil {
			return -1, errors.New("插入【评论表】出错:" + err.Error())
		}
		if id <= 0 {
			return -1, errors.New("插入【评论表】失败")
		}
		return int(id), nil
		//endregion

	}

}

//修改【评论表】信息
func UpdateCommentInfo(id int, ordercode string, commentor int, star int, info string, images string, adder int, addtime time.Time, moder int, modtime time.Time, shopid int) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【评论表】信息
	commentInfo, err := dao.CommentDao.Get(id)
	if err != nil {
		return false, errors.New("查询【评论表】信息出错:" + err.Error())
	}
	if commentInfo.Id <= 0 {
		return false, errors.New("【评论表】信息不存在")
	}
	if commentInfo.DeleteStatus != 1 {
		return false, errors.New("【评论表】信息已被删除")
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
	if commentInfo.Id == id &&
		commentInfo.Ordercode == ordercode &&
		commentInfo.Commentor == commentor &&
		commentInfo.Star == star &&
		commentInfo.Info == info &&
		commentInfo.Images == images &&
		commentInfo.Adder == adder &&
		commentInfo.Addtime == addtime &&
		commentInfo.Moder == moder &&
		commentInfo.Modtime == modtime &&
		commentInfo.Shopid == shopid {
		return true, nil
	}
	//endregion

	//region 构造【评论表】信息
	commentInfo.Ordercode = ordercode
	commentInfo.Commentor = commentor
	commentInfo.Star = star
	commentInfo.Info = info
	commentInfo.Images = images
	commentInfo.Adder = adder
	commentInfo.Addtime = addtime
	commentInfo.Moder = moder
	commentInfo.Modtime = modtime
	commentInfo.Shopid = shopid
	commentInfo.DeleteStatus = 1
	//endregion

	//region 修改【评论表】信息
	return dao.CommentDao.Update(&commentInfo)
	//endregion

}

//删除【评论表】信息
func DeleteCommentInfo(id int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【评论表】信息
	{
		commentInfo, err := dao.CommentDao.Get(id)
		if err != nil {
			return false, errors.New("查询【评论表】信息出错:" + err.Error())
		}
		if commentInfo.Id <= 0 {
			return false, errors.New("没有找到【评论表】信息")
		}
		if commentInfo.DeleteStatus != 1 {
			return false, errors.New("【评论表】信息已被删除")
		}
	}
	//endregion

	//region 删除【评论表】信息
	return dao.CommentDao.Delete(id)
	//endregion

}

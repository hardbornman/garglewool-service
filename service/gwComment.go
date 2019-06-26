package service

import (
	"errors"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"time"
)

// 获取【评论管理】信息
func GetGwCommentInfo(commentid int) (gwComment model.GwComment, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【评论管理】信息
	gwComment, err = dao.GwCommentDao.Get(commentid)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【评论ID】批量获取【评论管理】列表
func GetInGwCommentList(commentids []int) (gwComments []model.GwComment, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询列表
	gwComments, err = dao.GwCommentDao.GetIn(commentids)
	if err != nil {
		return
	}
	count = len(gwComments)
	//endregion

	return
}

// 分页查询【评论管理】列表
func GetGwCommentList(pageIndex, pageSize int) (gwComments []model.GwComment, count int, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询总数
	count, err = dao.GwCommentDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	gwComments, err = dao.GwCommentDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【评论管理】信息
func InsertGwCommentInfo(order_code string, order_commentor int, comment_star int, comment_info string, comment_images string, comment_createtime time.Time, adder int, addtime time.Time, moder int, modtime time.Time) (int, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//var err error
	isExist := false
	var gwCommentInfo model.GwComment

	//region 构造【评论管理】信息
	gwCommentInfo.OrderCode = order_code
	gwCommentInfo.OrderCommentor = order_commentor
	gwCommentInfo.CommentStar = comment_star
	gwCommentInfo.CommentInfo = comment_info
	gwCommentInfo.CommentImages = comment_images
	gwCommentInfo.CommentCreatetime = comment_createtime
	gwCommentInfo.Adder = adder
	gwCommentInfo.Addtime = addtime
	gwCommentInfo.Moder = moder
	gwCommentInfo.Modtime = modtime
	gwCommentInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【评论管理】信息
		isSuccess, err := dao.GwCommentDao.Update(&gwCommentInfo)
		if err != nil {
			return -1, errors.New("插入【评论管理】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【评论管理】失败")
		}
		return gwCommentInfo.Commentid, nil
		//endregion

	} else {

		//region 插入数据库
		commentid, err := dao.GwCommentDao.Insert(&gwCommentInfo)
		if err != nil {
			return -1, errors.New("插入【评论管理】出错:" + err.Error())
		}
		if commentid <= 0 {
			return -1, errors.New("插入【评论管理】失败")
		}
		return int(commentid), nil
		//endregion

	}

}

//修改【评论管理】信息
func UpdateGwCommentInfo(commentid int, order_code string, order_commentor int, comment_star int, comment_info string, comment_images string, comment_createtime time.Time, adder int, addtime time.Time, moder int, modtime time.Time) (bool, error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【评论管理】信息
	gwCommentInfo, err := dao.GwCommentDao.Get(commentid)
	if err != nil {
		return false, errors.New("查询【评论管理】信息出错:" + err.Error())
	}
	if gwCommentInfo.Commentid <= 0 {
		return false, errors.New("【评论管理】信息不存在")
	}
	if gwCommentInfo.DeleteStatus != 1 {
		return false, errors.New("【评论管理】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if gwCommentInfo.Commentid == commentid &&
		gwCommentInfo.OrderCode == order_code &&
		gwCommentInfo.OrderCommentor == order_commentor &&
		gwCommentInfo.CommentStar == comment_star &&
		gwCommentInfo.CommentInfo == comment_info &&
		gwCommentInfo.CommentImages == comment_images &&
		gwCommentInfo.CommentCreatetime == comment_createtime &&
		gwCommentInfo.Adder == adder &&
		gwCommentInfo.Addtime == addtime &&
		gwCommentInfo.Moder == moder &&
		gwCommentInfo.Modtime == modtime {
		return true, nil
	}
	//endregion

	//region 构造【评论管理】信息
	gwCommentInfo.OrderCode = order_code
	gwCommentInfo.OrderCommentor = order_commentor
	gwCommentInfo.CommentStar = comment_star
	gwCommentInfo.CommentInfo = comment_info
	gwCommentInfo.CommentImages = comment_images
	gwCommentInfo.CommentCreatetime = comment_createtime
	gwCommentInfo.Adder = adder
	gwCommentInfo.Addtime = addtime
	gwCommentInfo.Moder = moder
	gwCommentInfo.Modtime = modtime
	gwCommentInfo.DeleteStatus = 1
	//endregion

	//region 修改【评论管理】信息
	return dao.GwCommentDao.Update(&gwCommentInfo)
	//endregion

}

//删除【评论管理】信息
func DeleteGwCommentInfo(commentid int) (isSuccess bool, err error) {

	//region 验证请求路径
	validCallPath()
	//endregion

	//region 查询【评论管理】信息
	{
		gwCommentInfo, err := dao.GwCommentDao.Get(commentid)
		if err != nil {
			return false, errors.New("查询【评论管理】信息出错:" + err.Error())
		}
		if gwCommentInfo.Commentid <= 0 {
			return false, errors.New("没有找到【评论管理】信息")
		}
		if gwCommentInfo.DeleteStatus != 1 {
			return false, errors.New("【评论管理】信息已被删除")
		}
	}
	//endregion

	//region 删除【评论管理】信息
	return dao.GwCommentDao.Delete(commentid)
	//endregion

}

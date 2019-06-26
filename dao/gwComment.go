package dao

import (
	"bytes"
	"database/sql"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

var GwCommentDao = gwComment{}

type gwComment struct{}

// 根据【评论ID】查询【评论管理】表中是否存在相关记录
func (d *gwComment) Exist(commentid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from gw_comment where commentid=?", commentid)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	count := 0
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return false, err
		}
		return count > 0, nil
	}
	return false, nil
}

// 插入单条记录到【评论管理】表中
func (d *gwComment) Insert(m *model.GwComment) (int64, error) {
	result, err := garglewool.Exec("insert into gw_comment(order_code,order_commentor,comment_star,comment_info,comment_images,comment_createtime,adder,addtime,moder,modtime,deleteStatus) values(?,?,?,?,?,?,?,?,?,?,?)", m.OrderCode, m.OrderCommentor, m.CommentStar, m.CommentInfo, m.CommentImages, m.CommentCreatetime, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【评论ID】修改【评论管理】表的单条记录
func (d *gwComment) Update(m *model.GwComment) (bool, error) {
	result, err := garglewool.Exec("update gw_comment set order_code=?, order_commentor=?, comment_star=?, comment_info=?, comment_images=?, comment_createtime=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=? where commentid=?", m.OrderCode, m.OrderCommentor, m.CommentStar, m.CommentInfo, m.CommentImages, m.CommentCreatetime, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Commentid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【评论ID】软删除【评论管理】表中的单条记录
func (d *gwComment) Delete(commentid int) (bool, error) {
	result, err := garglewool.Exec("update gw_comment set deleteStatus=2 where commentid=?", commentid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【评论ID】数组软删除【评论管理】表中的多条记录
func (d *gwComment) DeleteIn(commentids []int) (count int64, err error) {
	if len(commentids) <= 0 {
		return count, errors.New("commentids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update gw_comment set deleteStatus=2")
	sql_str.WriteString(" where commentid in(")
	question_mark := strings.Repeat("?,", len(commentids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	var result sql.Result
	vals := make([]interface{}, 0, len(commentids))
	for _, v := range commentids {
		vals = append(vals, v)
	}
	result, err = garglewool.Exec(sql_str.String(), vals...)
	if err != nil {
		return count, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return count, err
	}
	return affected, nil
}

// 根据【评论ID】查询【评论管理】表中的单条记录
func (d *gwComment) Get(commentid int) (gwComment model.GwComment, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ commentid, order_code, order_commentor, comment_star, comment_info, comment_images, comment_createtime, adder, addtime, moder, modtime, deleteStatus from gw_comment where commentid=?", commentid)
	if err != nil {
		return gwComment, err
	}
	defer rows.Close()
	gwComments, err := d._RowsToArray(rows)
	if err != nil {
		return gwComment, err
	}
	if len(gwComments) <= 0 {
		return gwComment, err
	}
	return gwComments[0], nil
}

// 根据【评论ID】数组查询【评论管理】表中的多条记录
func (d *gwComment) GetIn(commentids []int) (gwComments []model.GwComment, err error) {
	if len(commentids) <= 0 {
		return gwComments, errors.New("commentids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ commentid, order_code, order_commentor, comment_star, comment_info, comment_images, comment_createtime, adder, addtime, moder, modtime, deleteStatus from ")
	sql_str.WriteString("gw_comment")
	sql_str.WriteString(" where commentid in(")
	param_keys := strings.Repeat("?,", len(commentids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(commentids))
	for _, v := range commentids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return gwComments, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【评论管理】表总记录数
func (d *gwComment) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from gw_comment")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return -1, err
		}
		return count, nil
	}
	//endregion

	return -1, nil
}

// 查询【评论管理】列表
func (d *gwComment) GetRowList(pageIndex, pageSize int) (gwComments []model.GwComment, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ commentid, order_code, order_commentor, comment_star, comment_info, comment_images, comment_createtime, adder, addtime, moder, modtime, deleteStatus from gw_comment")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by commentid desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return gwComments, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【评论管理】表记录
func (d *gwComment) _RowsToArray(rows *sqlx.Rows) (gwComments []model.GwComment, err error) {
	for rows.Next() {
		m := model.GwComment{}
		err = rows.Scan(&m.Commentid, &m.OrderCode, &m.OrderCommentor, &m.CommentStar, &m.CommentInfo, &m.CommentImages, &m.CommentCreatetime, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus)
		if err != nil {
			return gwComments, err
		}
		gwComments = append(gwComments, m)
	}
	return gwComments, err
}

package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

var CommentDao = comment{}

type comment struct{}

// 根据【评论ID】查询【评论表】表中是否存在相关记录
func (d *comment) Exist(id int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from comment where id=?", id)
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

// 插入单条记录到【评论表】表中
func (d *comment) Insert(m *model.Comment) (int64, error) {
	result, err := garglewool.Exec("insert into comment(ordercode,commentor,star,info,images,adder,addtime,moder,modtime,deleteStatus,shopid) values(?,?,?,?,?,?,?,?,?,?,?)", m.Ordercode, m.Commentor, m.Star, m.Info, m.Images, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Shopid)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【评论ID】修改【评论表】表的单条记录
func (d *comment) Update(m *model.Comment) (bool, error) {
	result, err := garglewool.Exec("update comment set ordercode=?, commentor=?, star=?, info=?, images=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=?, shopid=? where id=?", m.Ordercode, m.Commentor, m.Star, m.Info, m.Images, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Shopid, m.Id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【评论ID】软删除【评论表】表中的单条记录
func (d *comment) Delete(id int) (bool, error) {
	result, err := garglewool.Exec("update comment set deleteStatus=2 where id=?", id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【评论ID】数组软删除【评论表】表中的多条记录
func (d *comment) DeleteIn(ids []int) (count int64, err error) {
	if len(ids) <= 0 {
		return count, errors.New("ids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update comment set deleteStatus=2")
	sql_str.WriteString(" where id in(")
	question_mark := strings.Repeat("?,", len(ids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	vals := make([]interface{}, 0, len(ids))
	for _, v := range ids {
		vals = append(vals, v)
	}
	result, err := garglewool.Exec(sql_str.String(), vals...)
	if err != nil {
		return count, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return count, err
	}
	return affected, nil
}

// 根据【评论ID】查询【评论表】表中的单条记录
func (d *comment) Get(id int) (comment model.Comment, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ id, ordercode, commentor, star, info, images, adder, addtime, moder, modtime, deleteStatus, shopid from comment where id=?", id)
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	comments, err := d._RowsToArray(rows)
	if err != nil {
		return comment, err
	}
	if len(comments) <= 0 {
		return comment, err
	}
	return comments[0], nil
}

// 根据【评论ID】数组查询【评论表】表中的多条记录
func (d *comment) GetIn(ids []int) (comments []model.Comment, err error) {
	if len(ids) <= 0 {
		return comments, errors.New("ids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ id, ordercode, commentor, star, info, images, adder, addtime, moder, modtime, deleteStatus, shopid from ")
	sql_str.WriteString("comment")
	sql_str.WriteString(" where id in(")
	param_keys := strings.Repeat("?,", len(ids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(ids))
	for _, v := range ids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return comments, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【评论表】表总记录数
func (d *comment) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from comment")

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

// 根据【店铺ID】查询【评论表】总记录数
func (d *comment) GetRowCountByShopid(shopid int) (count int, err error) {
	rows, err := garglewool.Queryx("select count(0) Count from comment where shopid=? and deleteStatus = 1", shopid)
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
	return -1, err
}

// 查询【评论表】列表
func (d *comment) GetRowList(pageIndex, pageSize int) (comments []model.Comment, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ id, ordercode, commentor, star, info, images, adder, addtime, moder, modtime, deleteStatus, shopid from comment")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by id desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return comments, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【评论表】表记录
func (d *comment) _RowsToArray(rows *sqlx.Rows) (comments []model.Comment, err error) {
	for rows.Next() {
		m := model.Comment{}
		err = rows.Scan(&m.Id, &m.Ordercode, &m.Commentor, &m.Star, &m.Info, &m.Images, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus, &m.Shopid)
		if err != nil {
			return comments, err
		}
		comments = append(comments, m)
	}
	return comments, err
}

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

var GwUserDao = gwUser{}

type gwUser struct{}

// 根据【用户ID】查询【用户管理】表中是否存在相关记录
func (d *gwUser) Exist(userid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from gw_user where userid=?", userid)
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

// 插入单条记录到【用户管理】表中
func (d *gwUser) Insert(m *model.GwUser) (int64, error) {
	result, err := garglewool.Exec("insert into gw_user(user_name,user_phone,user_golds,adder,addtime,moder,modtime,deleteStatus) values(?,?,?,?,?,?,?,?)", m.UserName, m.UserPhone, m.UserGolds, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【用户ID】修改【用户管理】表的单条记录
func (d *gwUser) Update(m *model.GwUser) (bool, error) {
	result, err := garglewool.Exec("update gw_user set user_name=?, user_phone=?, user_golds=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=? where userid=?", m.UserName, m.UserPhone, m.UserGolds, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Userid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【用户ID】软删除【用户管理】表中的单条记录
func (d *gwUser) Delete(userid int) (bool, error) {
	result, err := garglewool.Exec("update gw_user set deleteStatus=2 where userid=?", userid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【用户ID】数组软删除【用户管理】表中的多条记录
func (d *gwUser) DeleteIn(userids []int) (count int64, err error) {
	if len(userids) <= 0 {
		return count, errors.New("userids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update gw_user set deleteStatus=2")
	sql_str.WriteString(" where userid in(")
	question_mark := strings.Repeat("?,", len(userids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	var result sql.Result
	vals := make([]interface{}, 0, len(userids))
	for _, v := range userids {
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

// 根据【用户ID】查询【用户管理】表中的单条记录
func (d *gwUser) Get(userid int) (gwUser model.GwUser, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ userid, user_name, user_phone, user_golds, adder, addtime, moder, modtime, deleteStatus from gw_user where userid=?", userid)
	if err != nil {
		return gwUser, err
	}
	defer rows.Close()
	gwUsers, err := d._RowsToArray(rows)
	if err != nil {
		return gwUser, err
	}
	if len(gwUsers) <= 0 {
		return gwUser, err
	}
	return gwUsers[0], nil
}

// 根据【用户ID】数组查询【用户管理】表中的多条记录
func (d *gwUser) GetIn(userids []int) (gwUsers []model.GwUser, err error) {
	if len(userids) <= 0 {
		return gwUsers, errors.New("userids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ userid, user_name, user_phone, user_golds, adder, addtime, moder, modtime, deleteStatus from ")
	sql_str.WriteString("gw_user")
	sql_str.WriteString(" where userid in(")
	param_keys := strings.Repeat("?,", len(userids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(userids))
	for _, v := range userids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return gwUsers, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【用户管理】表总记录数
func (d *gwUser) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from gw_user")

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

// 查询【用户管理】列表
func (d *gwUser) GetRowList(pageIndex, pageSize int) (gwUsers []model.GwUser, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ userid, user_name, user_phone, user_golds, adder, addtime, moder, modtime, deleteStatus from gw_user")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by userid desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return gwUsers, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【用户管理】表记录
func (d *gwUser) _RowsToArray(rows *sqlx.Rows) (gwUsers []model.GwUser, err error) {
	for rows.Next() {
		m := model.GwUser{}
		err = rows.Scan(&m.Userid, &m.UserName, &m.UserPhone, &m.UserGolds, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus)
		if err != nil {
			return gwUsers, err
		}
		gwUsers = append(gwUsers, m)
	}
	return gwUsers, err
}

// 查询【用户表】表中是否存在相关记录
func (d *gwUser) Login(loginName string, loginPwd string) (user model.GwUser, err error) {
	rows, err := garglewool.Queryx("select *  from gwuser where loginname=? and loginpwd=?", loginName, loginPwd)
	if err != nil {
		return user, err
	}
	defer rows.Close()
	users, err := d._RowsToArray(rows)
	if err != nil {
		return user, err
	}
	if len(users) <= 0 {
		return user, err
	}
	return users[0], nil
}

package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

var UserroleDao = userrole{}

type userrole struct{}

// 根据【用户角色ID】查询【用户角色表】表中是否存在相关记录
func (d *userrole) Exist(userroleid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from userrole where userroleid=?", userroleid)
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

// 插入单条记录到【用户角色表】表中
func (d *userrole) Insert(m *model.Userrole) (int64, error) {
	result, err := garglewool.Exec("insert into userrole(rolename,desc,remark,deleteStatus) values(?,?,?,?)", m.Rolename, m.Desc, m.Remark, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【用户角色ID】修改【用户角色表】表的单条记录
func (d *userrole) Update(m *model.Userrole) (bool, error) {
	result, err := garglewool.Exec("update userrole set rolename=?, desc=?, remark=?, deleteStatus=? where userroleid=?", m.Rolename, m.Desc, m.Remark, m.DeleteStatus, m.Userroleid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【用户角色ID】软删除【用户角色表】表中的单条记录
func (d *userrole) Delete(userroleid int) (bool, error) {
	result, err := garglewool.Exec("update userrole set deleteStatus=2 where userroleid=?", userroleid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【用户角色ID】数组软删除【用户角色表】表中的多条记录
func (d *userrole) DeleteIn(userroleids []int) (count int64, err error) {
	if len(userroleids) <= 0 {
		return count, errors.New("userroleids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update userrole set deleteStatus=2")
	sql_str.WriteString(" where userroleid in(")
	question_mark := strings.Repeat("?,", len(userroleids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	vals := make([]interface{}, 0, len(userroleids))
	for _, v := range userroleids {
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

// 根据【用户角色ID】查询【用户角色表】表中的单条记录
func (d *userrole) Get(userroleid int) (userrole model.Userrole, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ userroleid, rolename, desc, remark, deleteStatus, createtime, updatetime from userrole where userroleid=?", userroleid)
	if err != nil {
		return userrole, err
	}
	defer rows.Close()
	userroles, err := d._RowsToArray(rows)
	if err != nil {
		return userrole, err
	}
	if len(userroles) <= 0 {
		return userrole, err
	}
	return userroles[0], nil
}

// 根据【用户角色ID】数组查询【用户角色表】表中的多条记录
func (d *userrole) GetIn(userroleids []int) (userroles []model.Userrole, err error) {
	if len(userroleids) <= 0 {
		return userroles, errors.New("userroleids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ userroleid, rolename, desc, remark, deleteStatus, createtime, updatetime from ")
	sql_str.WriteString("userrole")
	sql_str.WriteString(" where userroleid in(")
	param_keys := strings.Repeat("?,", len(userroleids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(userroleids))
	for _, v := range userroleids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return userroles, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【用户角色表】表总记录数
func (d *userrole) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from userrole")

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

// 查询【用户角色表】列表
func (d *userrole) GetRowList(pageIndex, pageSize int) (userroles []model.Userrole, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ userroleid, rolename, desc, remark, deleteStatus, createtime, updatetime from userrole")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by userroleid desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return userroles, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【用户角色表】表记录
func (d *userrole) _RowsToArray(rows *sqlx.Rows) (userroles []model.Userrole, err error) {
	for rows.Next() {
		m := model.Userrole{}
		err = rows.Scan(&m.Userroleid, &m.Rolename, &m.Desc, &m.Remark, &m.DeleteStatus, &m.Createtime, &m.Updatetime)
		if err != nil {
			return userroles, err
		}
		userroles = append(userroles, m)
	}
	return userroles, err
}

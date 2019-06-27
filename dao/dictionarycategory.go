package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

var DictionarycategoryDao = dictionarycategory{}

type dictionarycategory struct{}

// 根据【字典分类ID】查询【字典分类】表中是否存在相关记录
func (d *dictionarycategory) Exist(dictionarycategoryid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from dictionarycategory where dictionarycategoryid=?", dictionarycategoryid)
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

// 插入单条记录到【字典分类】表中
func (d *dictionarycategory) Insert(m *model.Dictionarycategory) (int64, error) {
	result, err := garglewool.Exec("insert into dictionarycategory(categorykey,categoryvalue,adder,addtime,moder,modtime,deleteStatus) values(?,?,?,?,?,?,?)", m.Categorykey, m.Categoryvalue, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【字典分类ID】修改【字典分类】表的单条记录
func (d *dictionarycategory) Update(m *model.Dictionarycategory) (bool, error) {
	result, err := garglewool.Exec("update dictionarycategory set categorykey=?, categoryvalue=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=? where dictionarycategoryid=?", m.Categorykey, m.Categoryvalue, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Dictionarycategoryid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【字典分类ID】软删除【字典分类】表中的单条记录
func (d *dictionarycategory) Delete(dictionarycategoryid int) (bool, error) {
	result, err := garglewool.Exec("update dictionarycategory set deleteStatus=2 where dictionarycategoryid=?", dictionarycategoryid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【字典分类ID】数组软删除【字典分类】表中的多条记录
func (d *dictionarycategory) DeleteIn(dictionarycategoryids []int) (count int64, err error) {
	if len(dictionarycategoryids) <= 0 {
		return count, errors.New("dictionarycategoryids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update dictionarycategory set deleteStatus=2")
	sql_str.WriteString(" where dictionarycategoryid in(")
	question_mark := strings.Repeat("?,", len(dictionarycategoryids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	vals := make([]interface{}, 0, len(dictionarycategoryids))
	for _, v := range dictionarycategoryids {
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

// 根据【字典分类ID】查询【字典分类】表中的单条记录
func (d *dictionarycategory) Get(dictionarycategoryid int) (dictionarycategory model.Dictionarycategory, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ dictionarycategoryid, categorykey, categoryvalue, adder, addtime, moder, modtime, deleteStatus from dictionarycategory where dictionarycategoryid=?", dictionarycategoryid)
	if err != nil {
		return dictionarycategory, err
	}
	defer rows.Close()
	dictionarycategorys, err := d._RowsToArray(rows)
	if err != nil {
		return dictionarycategory, err
	}
	if len(dictionarycategorys) <= 0 {
		return dictionarycategory, err
	}
	return dictionarycategorys[0], nil
}

// 根据【字典分类ID】数组查询【字典分类】表中的多条记录
func (d *dictionarycategory) GetIn(dictionarycategoryids []int) (dictionarycategorys []model.Dictionarycategory, err error) {
	if len(dictionarycategoryids) <= 0 {
		return dictionarycategorys, errors.New("dictionarycategoryids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ dictionarycategoryid, categorykey, categoryvalue, adder, addtime, moder, modtime, deleteStatus from ")
	sql_str.WriteString("dictionarycategory")
	sql_str.WriteString(" where dictionarycategoryid in(")
	param_keys := strings.Repeat("?,", len(dictionarycategoryids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(dictionarycategoryids))
	for _, v := range dictionarycategoryids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return dictionarycategorys, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【字典分类】表总记录数
func (d *dictionarycategory) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from dictionarycategory")

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

// 查询【字典分类】列表
func (d *dictionarycategory) GetRowList(pageIndex, pageSize int) (dictionarycategorys []model.Dictionarycategory, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ dictionarycategoryid, categorykey, categoryvalue, adder, addtime, moder, modtime, deleteStatus from dictionarycategory")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by dictionarycategoryid desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return dictionarycategorys, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【字典分类】表记录
func (d *dictionarycategory) _RowsToArray(rows *sqlx.Rows) (dictionarycategorys []model.Dictionarycategory, err error) {
	for rows.Next() {
		m := model.Dictionarycategory{}
		err = rows.Scan(&m.Dictionarycategoryid, &m.Categorykey, &m.Categoryvalue, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus)
		if err != nil {
			return dictionarycategorys, err
		}
		dictionarycategorys = append(dictionarycategorys, m)
	}
	return dictionarycategorys, err
}

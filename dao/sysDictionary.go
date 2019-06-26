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

var SysDictionaryDao = sysDictionary{}

type sysDictionary struct{}

// 根据【字典ID】查询【字典表】表中是否存在相关记录
func (d *sysDictionary) Exist(dictionaryid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from sys_dictionary where dictionaryid=?", dictionaryid)
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

// 插入单条记录到【字典表】表中
func (d *sysDictionary) Insert(m *model.SysDictionary) (int64, error) {
	result, err := garglewool.Exec("insert into sys_dictionary(categorykey,dictkey,dictvalue,adder,addtime,moder,modtime,deleteStatus) values(?,?,?,?,?,?,?,?)", m.Categorykey, m.Dictkey, m.Dictvalue, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【字典ID】修改【字典表】表的单条记录
func (d *sysDictionary) Update(m *model.SysDictionary) (bool, error) {
	result, err := garglewool.Exec("update sys_dictionary set categorykey=?, dictkey=?, dictvalue=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=? where dictionaryid=?", m.Categorykey, m.Dictkey, m.Dictvalue, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Dictionaryid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【字典ID】软删除【字典表】表中的单条记录
func (d *sysDictionary) Delete(dictionaryid int) (bool, error) {
	result, err := garglewool.Exec("update sys_dictionary set deleteStatus=2 where dictionaryid=?", dictionaryid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【字典ID】数组软删除【字典表】表中的多条记录
func (d *sysDictionary) DeleteIn(dictionaryids []int) (count int64, err error) {
	if len(dictionaryids) <= 0 {
		return count, errors.New("dictionaryids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update sys_dictionary set deleteStatus=2")
	sql_str.WriteString(" where dictionaryid in(")
	question_mark := strings.Repeat("?,", len(dictionaryids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	var result sql.Result
	vals := make([]interface{}, 0, len(dictionaryids))
	for _, v := range dictionaryids {
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

// 根据【字典ID】查询【字典表】表中的单条记录
func (d *sysDictionary) Get(dictionaryid int) (sysDictionary model.SysDictionary, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ dictionaryid, categorykey, dictkey, dictvalue, adder, addtime, moder, modtime, deleteStatus from sys_dictionary where dictionaryid=?", dictionaryid)
	if err != nil {
		return sysDictionary, err
	}
	defer rows.Close()
	sysDictionarys, err := d._RowsToArray(rows)
	if err != nil {
		return sysDictionary, err
	}
	if len(sysDictionarys) <= 0 {
		return sysDictionary, err
	}
	return sysDictionarys[0], nil
}

// 根据【字典ID】数组查询【字典表】表中的多条记录
func (d *sysDictionary) GetIn(dictionaryids []int) (sysDictionarys []model.SysDictionary, err error) {
	if len(dictionaryids) <= 0 {
		return sysDictionarys, errors.New("dictionaryids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ dictionaryid, categorykey, dictkey, dictvalue, adder, addtime, moder, modtime, deleteStatus from ")
	sql_str.WriteString("sys_dictionary")
	sql_str.WriteString(" where dictionaryid in(")
	param_keys := strings.Repeat("?,", len(dictionaryids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(dictionaryids))
	for _, v := range dictionaryids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return sysDictionarys, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【字典表】表总记录数
func (d *sysDictionary) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from sys_dictionary")

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

// 查询【字典表】列表
func (d *sysDictionary) GetRowList(pageIndex, pageSize int) (sysDictionarys []model.SysDictionary, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ dictionaryid, categorykey, dictkey, dictvalue, adder, addtime, moder, modtime, deleteStatus from sys_dictionary")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by dictionaryid desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return sysDictionarys, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【字典表】表记录
func (d *sysDictionary) _RowsToArray(rows *sqlx.Rows) (sysDictionarys []model.SysDictionary, err error) {
	for rows.Next() {
		m := model.SysDictionary{}
		err = rows.Scan(&m.Dictionaryid, &m.Categorykey, &m.Dictkey, &m.Dictvalue, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus)
		if err != nil {
			return sysDictionarys, err
		}
		sysDictionarys = append(sysDictionarys, m)
	}
	return sysDictionarys, err
}

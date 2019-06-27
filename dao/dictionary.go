package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

var DictionaryDao = dictionary{}

type dictionary struct{}

// 根据【字典ID】查询【字典表】表中是否存在相关记录
func (d *dictionary) Exist(dictionaryid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from dictionary where dictionaryid=?", dictionaryid)
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
func (d *dictionary) Insert(m *model.Dictionary) (int64, error) {
	result, err := garglewool.Exec("insert into dictionary(categorykey,dictkey,dictvalue,adder,addtime,moder,modtime,deleteStatus) values(?,?,?,?,?,?,?,?)", m.Categorykey, m.Dictkey, m.Dictvalue, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【字典ID】修改【字典表】表的单条记录
func (d *dictionary) Update(m *model.Dictionary) (bool, error) {
	result, err := garglewool.Exec("update dictionary set categorykey=?, dictkey=?, dictvalue=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=? where dictionaryid=?", m.Categorykey, m.Dictkey, m.Dictvalue, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Dictionaryid)
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
func (d *dictionary) Delete(dictionaryid int) (bool, error) {
	result, err := garglewool.Exec("update dictionary set deleteStatus=2 where dictionaryid=?", dictionaryid)
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
func (d *dictionary) DeleteIn(dictionaryids []int) (count int64, err error) {
	if len(dictionaryids) <= 0 {
		return count, errors.New("dictionaryids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update dictionary set deleteStatus=2")
	sql_str.WriteString(" where dictionaryid in(")
	question_mark := strings.Repeat("?,", len(dictionaryids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	vals := make([]interface{}, 0, len(dictionaryids))
	for _, v := range dictionaryids {
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

// 根据【字典ID】查询【字典表】表中的单条记录
func (d *dictionary) Get(dictionaryid int) (dictionary model.Dictionary, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ dictionaryid, categorykey, dictkey, dictvalue, adder, addtime, moder, modtime, deleteStatus from dictionary where dictionaryid=?", dictionaryid)
	if err != nil {
		return dictionary, err
	}
	defer rows.Close()
	dictionarys, err := d._RowsToArray(rows)
	if err != nil {
		return dictionary, err
	}
	if len(dictionarys) <= 0 {
		return dictionary, err
	}
	return dictionarys[0], nil
}

// 根据【字典ID】数组查询【字典表】表中的多条记录
func (d *dictionary) GetIn(dictionaryids []int) (dictionarys []model.Dictionary, err error) {
	if len(dictionaryids) <= 0 {
		return dictionarys, errors.New("dictionaryids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ dictionaryid, categorykey, dictkey, dictvalue, adder, addtime, moder, modtime, deleteStatus from ")
	sql_str.WriteString("dictionary")
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
		return dictionarys, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【字典表】表总记录数
func (d *dictionary) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from dictionary")

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
func (d *dictionary) GetRowList(pageIndex, pageSize int) (dictionarys []model.Dictionary, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ dictionaryid, categorykey, dictkey, dictvalue, adder, addtime, moder, modtime, deleteStatus from dictionary")

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
		return dictionarys, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【字典表】表记录
func (d *dictionary) _RowsToArray(rows *sqlx.Rows) (dictionarys []model.Dictionary, err error) {
	for rows.Next() {
		m := model.Dictionary{}
		err = rows.Scan(&m.Dictionaryid, &m.Categorykey, &m.Dictkey, &m.Dictvalue, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus)
		if err != nil {
			return dictionarys, err
		}
		dictionarys = append(dictionarys, m)
	}
	return dictionarys, err
}

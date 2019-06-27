package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

var SetmealdetailDao = setmealdetail{}

type setmealdetail struct{}

// 根据【套餐明细ID】查询【套餐明细表】表中是否存在相关记录
func (d *setmealdetail) Exist(setmealdetailid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from setmealdetail where setmealdetailid=?", setmealdetailid)
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

// 插入单条记录到【套餐明细表】表中
func (d *setmealdetail) Insert(m *model.Setmealdetail) (int64, error) {
	result, err := garglewool.Exec("insert into setmealdetail(setmealid,name,nums,price,adder,addtime,moder,modtime,deleteStatus) values(?,?,?,?,?,?,?,?,?)", m.Setmealid, m.Name, m.Nums, m.Price, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【套餐明细ID】修改【套餐明细表】表的单条记录
func (d *setmealdetail) Update(m *model.Setmealdetail) (bool, error) {
	result, err := garglewool.Exec("update setmealdetail set setmealid=?, name=?, nums=?, price=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=? where setmealdetailid=?", m.Setmealid, m.Name, m.Nums, m.Price, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Setmealdetailid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【套餐明细ID】软删除【套餐明细表】表中的单条记录
func (d *setmealdetail) Delete(setmealdetailid int) (bool, error) {
	result, err := garglewool.Exec("update setmealdetail set deleteStatus=2 where setmealdetailid=?", setmealdetailid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【套餐明细ID】数组软删除【套餐明细表】表中的多条记录
func (d *setmealdetail) DeleteIn(setmealdetailids []int) (count int64, err error) {
	if len(setmealdetailids) <= 0 {
		return count, errors.New("setmealdetailids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update setmealdetail set deleteStatus=2")
	sql_str.WriteString(" where setmealdetailid in(")
	question_mark := strings.Repeat("?,", len(setmealdetailids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	vals := make([]interface{}, 0, len(setmealdetailids))
	for _, v := range setmealdetailids {
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

// 根据【套餐明细ID】查询【套餐明细表】表中的单条记录
func (d *setmealdetail) Get(setmealdetailid int) (setmealdetail model.Setmealdetail, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ setmealdetailid, setmealid, name, nums, price, adder, addtime, moder, modtime, deleteStatus from setmealdetail where setmealdetailid=?", setmealdetailid)
	if err != nil {
		return setmealdetail, err
	}
	defer rows.Close()
	setmealdetails, err := d._RowsToArray(rows)
	if err != nil {
		return setmealdetail, err
	}
	if len(setmealdetails) <= 0 {
		return setmealdetail, err
	}
	return setmealdetails[0], nil
}

// 根据【套餐明细ID】数组查询【套餐明细表】表中的多条记录
func (d *setmealdetail) GetIn(setmealdetailids []int) (setmealdetails []model.Setmealdetail, err error) {
	if len(setmealdetailids) <= 0 {
		return setmealdetails, errors.New("setmealdetailids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ setmealdetailid, setmealid, name, nums, price, adder, addtime, moder, modtime, deleteStatus from ")
	sql_str.WriteString("setmealdetail")
	sql_str.WriteString(" where setmealdetailid in(")
	param_keys := strings.Repeat("?,", len(setmealdetailids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(setmealdetailids))
	for _, v := range setmealdetailids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return setmealdetails, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【套餐明细表】表总记录数
func (d *setmealdetail) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from setmealdetail")

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

// 根据【套餐ID】查询【套餐明细表】总记录数
func (d *setmealdetail) GetRowCountBySetmealid(setmealid int) (count int, err error) {
	rows, err := garglewool.Queryx("select count(0) Count from setmealdetail where setmealid=? and deleteStatus = 1", setmealid)
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

// 查询【套餐明细表】列表
func (d *setmealdetail) GetRowList(pageIndex, pageSize int) (setmealdetails []model.Setmealdetail, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ setmealdetailid, setmealid, name, nums, price, adder, addtime, moder, modtime, deleteStatus from setmealdetail")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by setmealdetailid desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return setmealdetails, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【套餐明细表】表记录
func (d *setmealdetail) _RowsToArray(rows *sqlx.Rows) (setmealdetails []model.Setmealdetail, err error) {
	for rows.Next() {
		m := model.Setmealdetail{}
		err = rows.Scan(&m.Setmealdetailid, &m.Setmealid, &m.Name, &m.Nums, &m.Price, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus)
		if err != nil {
			return setmealdetails, err
		}
		setmealdetails = append(setmealdetails, m)
	}
	return setmealdetails, err
}

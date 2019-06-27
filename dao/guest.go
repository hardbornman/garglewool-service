package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

var GuestDao = guest{}

type guest struct{}

// 根据【用户ID】查询【买家客户表】表中是否存在相关记录
func (d *guest) Exist(guestid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from guest where guestid=?", guestid)
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

// 插入单条记录到【买家客户表】表中
func (d *guest) Insert(m *model.Guest) (int64, error) {
	result, err := garglewool.Exec("insert into guest(name,password,phone,golds,adder,addtime,moder,modtime,deleteStatus) values(?,?,?,?,?,?,?,?,?)", m.Name, m.Password, m.Phone, m.Golds, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【用户ID】修改【买家客户表】表的单条记录
func (d *guest) Update(m *model.Guest) (bool, error) {
	result, err := garglewool.Exec("update guest set name=?, password=?, phone=?, golds=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=? where guestid=?", m.Name, m.Password, m.Phone, m.Golds, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Guestid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【用户ID】软删除【买家客户表】表中的单条记录
func (d *guest) Delete(guestid int) (bool, error) {
	result, err := garglewool.Exec("update guest set deleteStatus=2 where guestid=?", guestid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【用户ID】数组软删除【买家客户表】表中的多条记录
func (d *guest) DeleteIn(guestids []int) (count int64, err error) {
	if len(guestids) <= 0 {
		return count, errors.New("guestids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update guest set deleteStatus=2")
	sql_str.WriteString(" where guestid in(")
	question_mark := strings.Repeat("?,", len(guestids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	vals := make([]interface{}, 0, len(guestids))
	for _, v := range guestids {
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

// 根据【用户ID】查询【买家客户表】表中的单条记录
func (d *guest) Get(guestid int) (guest model.Guest, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ guestid, name, password, phone, golds, adder, addtime, moder, modtime, deleteStatus from guest where guestid=?", guestid)
	if err != nil {
		return guest, err
	}
	defer rows.Close()
	guests, err := d._RowsToArray(rows)
	if err != nil {
		return guest, err
	}
	if len(guests) <= 0 {
		return guest, err
	}
	return guests[0], nil
}

// 根据【用户ID】数组查询【买家客户表】表中的多条记录
func (d *guest) GetIn(guestids []int) (guests []model.Guest, err error) {
	if len(guestids) <= 0 {
		return guests, errors.New("guestids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ guestid, name, password, phone, golds, adder, addtime, moder, modtime, deleteStatus from ")
	sql_str.WriteString("guest")
	sql_str.WriteString(" where guestid in(")
	param_keys := strings.Repeat("?,", len(guestids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(guestids))
	for _, v := range guestids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return guests, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【买家客户表】表总记录数
func (d *guest) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from guest")

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

// 查询【买家客户表】列表
func (d *guest) GetRowList(pageIndex, pageSize int) (guests []model.Guest, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ guestid, name, password, phone, golds, adder, addtime, moder, modtime, deleteStatus from guest")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by guestid desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return guests, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【买家客户表】表记录
func (d *guest) _RowsToArray(rows *sqlx.Rows) (guests []model.Guest, err error) {
	for rows.Next() {
		m := model.Guest{}
		err = rows.Scan(&m.Guestid, &m.Name, &m.Password, &m.Phone, &m.Golds, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus)
		if err != nil {
			return guests, err
		}
		guests = append(guests, m)
	}
	return guests, err
}

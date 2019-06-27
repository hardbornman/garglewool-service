package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

var VoucherDao = voucher{}

type voucher struct{}

// 根据【抵用券ID】查询【抵用券管理】表中是否存在相关记录
func (d *voucher) Exist(voucherid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from voucher where voucherid=?", voucherid)
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

// 插入单条记录到【抵用券管理】表中
func (d *voucher) Insert(m *model.Voucher) (int64, error) {
	result, err := garglewool.Exec("insert into voucher(code,guestid,quota,reduce,validdays,isinvalid,adder,addtime,moder,modtime,deleteStatus) values(?,?,?,?,?,?,?,?,?,?,?)", m.Code, m.Guestid, m.Quota, m.Reduce, m.Validdays, m.Isinvalid, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【抵用券ID】修改【抵用券管理】表的单条记录
func (d *voucher) Update(m *model.Voucher) (bool, error) {
	result, err := garglewool.Exec("update voucher set code=?, guestid=?, quota=?, reduce=?, validdays=?, isinvalid=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=? where voucherid=?", m.Code, m.Guestid, m.Quota, m.Reduce, m.Validdays, m.Isinvalid, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Voucherid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【抵用券ID】软删除【抵用券管理】表中的单条记录
func (d *voucher) Delete(voucherid int) (bool, error) {
	result, err := garglewool.Exec("update voucher set deleteStatus=2 where voucherid=?", voucherid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【抵用券ID】数组软删除【抵用券管理】表中的多条记录
func (d *voucher) DeleteIn(voucherids []int) (count int64, err error) {
	if len(voucherids) <= 0 {
		return count, errors.New("voucherids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update voucher set deleteStatus=2")
	sql_str.WriteString(" where voucherid in(")
	question_mark := strings.Repeat("?,", len(voucherids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	vals := make([]interface{}, 0, len(voucherids))
	for _, v := range voucherids {
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

// 根据【抵用券ID】查询【抵用券管理】表中的单条记录
func (d *voucher) Get(voucherid int) (voucher model.Voucher, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ voucherid, code, guestid, quota, reduce, createtime, validdays, isinvalid, adder, addtime, moder, modtime, deleteStatus from voucher where voucherid=?", voucherid)
	if err != nil {
		return voucher, err
	}
	defer rows.Close()
	vouchers, err := d._RowsToArray(rows)
	if err != nil {
		return voucher, err
	}
	if len(vouchers) <= 0 {
		return voucher, err
	}
	return vouchers[0], nil
}

// 根据【抵用券ID】数组查询【抵用券管理】表中的多条记录
func (d *voucher) GetIn(voucherids []int) (vouchers []model.Voucher, err error) {
	if len(voucherids) <= 0 {
		return vouchers, errors.New("voucherids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ voucherid, code, guestid, quota, reduce, createtime, validdays, isinvalid, adder, addtime, moder, modtime, deleteStatus from ")
	sql_str.WriteString("voucher")
	sql_str.WriteString(" where voucherid in(")
	param_keys := strings.Repeat("?,", len(voucherids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(voucherids))
	for _, v := range voucherids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return vouchers, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【抵用券管理】表总记录数
func (d *voucher) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from voucher")

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

// 根据【用户id】查询【抵用券管理】总记录数
func (d *voucher) GetRowCountByGuestid(guestid int) (count int, err error) {
	rows, err := garglewool.Queryx("select count(0) Count from voucher where guestid=? and deleteStatus = 1", guestid)
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

// 查询【抵用券管理】列表
func (d *voucher) GetRowList(pageIndex, pageSize int) (vouchers []model.Voucher, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ voucherid, code, guestid, quota, reduce, createtime, validdays, isinvalid, adder, addtime, moder, modtime, deleteStatus from voucher")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by voucherid desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return vouchers, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【抵用券管理】表记录
func (d *voucher) _RowsToArray(rows *sqlx.Rows) (vouchers []model.Voucher, err error) {
	for rows.Next() {
		m := model.Voucher{}
		err = rows.Scan(&m.Voucherid, &m.Code, &m.Guestid, &m.Quota, &m.Reduce, &m.Createtime, &m.Validdays, &m.Isinvalid, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus)
		if err != nil {
			return vouchers, err
		}
		vouchers = append(vouchers, m)
	}
	return vouchers, err
}

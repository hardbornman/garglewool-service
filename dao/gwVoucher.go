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

var GwVoucherDao = gwVoucher{}

type gwVoucher struct{}

// 根据【抵用券ID】查询【抵用券管理】表中是否存在相关记录
func (d *gwVoucher) Exist(voucherid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from gw_voucher where voucherid=?", voucherid)
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
func (d *gwVoucher) Insert(m *model.GwVoucher) (int64, error) {
	result, err := garglewool.Exec("insert into gw_voucher(voucher_code,voucher_userid,voucher_quota,voucher_reduce,voucher_createtime,voucher_validdays,voucher_isinvalid,adder,addtime,moder,modtime,deleteStatus) values(?,?,?,?,?,?,?,?,?,?,?,?)", m.VoucherCode, m.VoucherUserid, m.VoucherQuota, m.VoucherReduce, m.VoucherCreatetime, m.VoucherValiddays, m.VoucherIsinvalid, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【抵用券ID】修改【抵用券管理】表的单条记录
func (d *gwVoucher) Update(m *model.GwVoucher) (bool, error) {
	result, err := garglewool.Exec("update gw_voucher set voucher_code=?, voucher_userid=?, voucher_quota=?, voucher_reduce=?, voucher_createtime=?, voucher_validdays=?, voucher_isinvalid=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=? where voucherid=?", m.VoucherCode, m.VoucherUserid, m.VoucherQuota, m.VoucherReduce, m.VoucherCreatetime, m.VoucherValiddays, m.VoucherIsinvalid, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Voucherid)
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
func (d *gwVoucher) Delete(voucherid int) (bool, error) {
	result, err := garglewool.Exec("update gw_voucher set deleteStatus=2 where voucherid=?", voucherid)
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
func (d *gwVoucher) DeleteIn(voucherids []int) (count int64, err error) {
	if len(voucherids) <= 0 {
		return count, errors.New("voucherids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update gw_voucher set deleteStatus=2")
	sql_str.WriteString(" where voucherid in(")
	question_mark := strings.Repeat("?,", len(voucherids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	var result sql.Result
	vals := make([]interface{}, 0, len(voucherids))
	for _, v := range voucherids {
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

// 根据【抵用券ID】查询【抵用券管理】表中的单条记录
func (d *gwVoucher) Get(voucherid int) (gwVoucher model.GwVoucher, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ voucherid, voucher_code, voucher_userid, voucher_quota, voucher_reduce, voucher_createtime, voucher_validdays, voucher_isinvalid, adder, addtime, moder, modtime, deleteStatus from gw_voucher where voucherid=?", voucherid)
	if err != nil {
		return gwVoucher, err
	}
	defer rows.Close()
	gwVouchers, err := d._RowsToArray(rows)
	if err != nil {
		return gwVoucher, err
	}
	if len(gwVouchers) <= 0 {
		return gwVoucher, err
	}
	return gwVouchers[0], nil
}

// 根据【抵用券ID】数组查询【抵用券管理】表中的多条记录
func (d *gwVoucher) GetIn(voucherids []int) (gwVouchers []model.GwVoucher, err error) {
	if len(voucherids) <= 0 {
		return gwVouchers, errors.New("voucherids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ voucherid, voucher_code, voucher_userid, voucher_quota, voucher_reduce, voucher_createtime, voucher_validdays, voucher_isinvalid, adder, addtime, moder, modtime, deleteStatus from ")
	sql_str.WriteString("gw_voucher")
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
		return gwVouchers, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【抵用券管理】表总记录数
func (d *gwVoucher) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from gw_voucher")

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

// 查询【抵用券管理】列表
func (d *gwVoucher) GetRowList(pageIndex, pageSize int) (gwVouchers []model.GwVoucher, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ voucherid, voucher_code, voucher_userid, voucher_quota, voucher_reduce, voucher_createtime, voucher_validdays, voucher_isinvalid, adder, addtime, moder, modtime, deleteStatus from gw_voucher")

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
		return gwVouchers, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【抵用券管理】表记录
func (d *gwVoucher) _RowsToArray(rows *sqlx.Rows) (gwVouchers []model.GwVoucher, err error) {
	for rows.Next() {
		m := model.GwVoucher{}
		err = rows.Scan(&m.Voucherid, &m.VoucherCode, &m.VoucherUserid, &m.VoucherQuota, &m.VoucherReduce, &m.VoucherCreatetime, &m.VoucherValiddays, &m.VoucherIsinvalid, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus)
		if err != nil {
			return gwVouchers, err
		}
		gwVouchers = append(gwVouchers, m)
	}
	return gwVouchers, err
}

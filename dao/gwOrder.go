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

var GwOrderDao = gwOrder{}

type gwOrder struct{}

// 根据【订单ID】查询【订单管理】表中是否存在相关记录
func (d *gwOrder) Exist(orderid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from gw_order where orderid=?", orderid)
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

// 插入单条记录到【订单管理】表中
func (d *gwOrder) Insert(m *model.GwOrder) (int64, error) {
	result, err := garglewool.Exec("insert into gw_order(pkg_code,order_code,order_buyer,order_paytype,order_totalprice,order_payprice,order_paytime,order_isinvalid,order_isused,order_isrefund,order_refundprice,order_refundtime,order_remark,adder,addtime,moder,modtime,deleteStatus) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", m.PkgCode, m.OrderCode, m.OrderBuyer, m.OrderPaytype, m.OrderTotalprice, m.OrderPayprice, m.OrderPaytime, m.OrderIsinvalid, m.OrderIsused, m.OrderIsrefund, m.OrderRefundprice, m.OrderRefundtime, m.OrderRemark, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【订单ID】修改【订单管理】表的单条记录
func (d *gwOrder) Update(m *model.GwOrder) (bool, error) {
	result, err := garglewool.Exec("update gw_order set pkg_code=?, order_code=?, order_buyer=?, order_paytype=?, order_totalprice=?, order_payprice=?, order_paytime=?, order_isinvalid=?, order_isused=?, order_isrefund=?, order_refundprice=?, order_refundtime=?, order_remark=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=? where orderid=?", m.PkgCode, m.OrderCode, m.OrderBuyer, m.OrderPaytype, m.OrderTotalprice, m.OrderPayprice, m.OrderPaytime, m.OrderIsinvalid, m.OrderIsused, m.OrderIsrefund, m.OrderRefundprice, m.OrderRefundtime, m.OrderRemark, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Orderid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【订单ID】软删除【订单管理】表中的单条记录
func (d *gwOrder) Delete(orderid int) (bool, error) {
	result, err := garglewool.Exec("update gw_order set deleteStatus=2 where orderid=?", orderid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【订单ID】数组软删除【订单管理】表中的多条记录
func (d *gwOrder) DeleteIn(orderids []int) (count int64, err error) {
	if len(orderids) <= 0 {
		return count, errors.New("orderids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update gw_order set deleteStatus=2")
	sql_str.WriteString(" where orderid in(")
	question_mark := strings.Repeat("?,", len(orderids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	var result sql.Result
	vals := make([]interface{}, 0, len(orderids))
	for _, v := range orderids {
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

// 根据【订单ID】查询【订单管理】表中的单条记录
func (d *gwOrder) Get(orderid int) (gwOrder model.GwOrder, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ orderid, pkg_code, order_code, order_buyer, order_paytype, order_totalprice, order_payprice, order_paytime, order_isinvalid, order_isused, order_isrefund, order_refundprice, order_refundtime, order_remark, adder, addtime, moder, modtime, deleteStatus from gw_order where orderid=?", orderid)
	if err != nil {
		return gwOrder, err
	}
	defer rows.Close()
	gwOrders, err := d._RowsToArray(rows)
	if err != nil {
		return gwOrder, err
	}
	if len(gwOrders) <= 0 {
		return gwOrder, err
	}
	return gwOrders[0], nil
}

// 根据【订单ID】数组查询【订单管理】表中的多条记录
func (d *gwOrder) GetIn(orderids []int) (gwOrders []model.GwOrder, err error) {
	if len(orderids) <= 0 {
		return gwOrders, errors.New("orderids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ orderid, pkg_code, order_code, order_buyer, order_paytype, order_totalprice, order_payprice, order_paytime, order_isinvalid, order_isused, order_isrefund, order_refundprice, order_refundtime, order_remark, adder, addtime, moder, modtime, deleteStatus from ")
	sql_str.WriteString("gw_order")
	sql_str.WriteString(" where orderid in(")
	param_keys := strings.Repeat("?,", len(orderids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(orderids))
	for _, v := range orderids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return gwOrders, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【订单管理】表总记录数
func (d *gwOrder) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from gw_order")

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

// 查询【订单管理】列表
func (d *gwOrder) GetRowList(pageIndex, pageSize int) (gwOrders []model.GwOrder, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ orderid, pkg_code, order_code, order_buyer, order_paytype, order_totalprice, order_payprice, order_paytime, order_isinvalid, order_isused, order_isrefund, order_refundprice, order_refundtime, order_remark, adder, addtime, moder, modtime, deleteStatus from gw_order")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by orderid desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return gwOrders, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【订单管理】表记录
func (d *gwOrder) _RowsToArray(rows *sqlx.Rows) (gwOrders []model.GwOrder, err error) {
	for rows.Next() {
		m := model.GwOrder{}
		err = rows.Scan(&m.Orderid, &m.PkgCode, &m.OrderCode, &m.OrderBuyer, &m.OrderPaytype, &m.OrderTotalprice, &m.OrderPayprice, &m.OrderPaytime, &m.OrderIsinvalid, &m.OrderIsused, &m.OrderIsrefund, &m.OrderRefundprice, &m.OrderRefundtime, &m.OrderRemark, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus)
		if err != nil {
			return gwOrders, err
		}
		gwOrders = append(gwOrders, m)
	}
	return gwOrders, err
}

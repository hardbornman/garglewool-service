package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

var OrderDao = order{}

type order struct{}

// 根据【订单ID】查询【订单表】表中是否存在相关记录
func (d *order) Exist(orderid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from order where orderid=?", orderid)
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

// 插入单条记录到【订单表】表中
func (d *order) Insert(m *model.Order) (int64, error) {
	result, err := garglewool.Exec("insert into order(pkgcode,ordercode,buyer,paytype,totalprice,payprice,paytime,isinvalid,isused,isrefund,refundprice,refundtime,remark,adder,addtime,moder,modtime,deleteStatus) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", m.Pkgcode, m.Ordercode, m.Buyer, m.Paytype, m.Totalprice, m.Payprice, m.Paytime, m.Isinvalid, m.Isused, m.Isrefund, m.Refundprice, m.Refundtime, m.Remark, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【订单ID】修改【订单表】表的单条记录
func (d *order) Update(m *model.Order) (bool, error) {
	result, err := garglewool.Exec("update order set pkgcode=?, ordercode=?, buyer=?, paytype=?, totalprice=?, payprice=?, paytime=?, isinvalid=?, isused=?, isrefund=?, refundprice=?, refundtime=?, remark=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=? where orderid=?", m.Pkgcode, m.Ordercode, m.Buyer, m.Paytype, m.Totalprice, m.Payprice, m.Paytime, m.Isinvalid, m.Isused, m.Isrefund, m.Refundprice, m.Refundtime, m.Remark, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Orderid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【订单ID】软删除【订单表】表中的单条记录
func (d *order) Delete(orderid int) (bool, error) {
	result, err := garglewool.Exec("update order set deleteStatus=2 where orderid=?", orderid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【订单ID】数组软删除【订单表】表中的多条记录
func (d *order) DeleteIn(orderids []int) (count int64, err error) {
	if len(orderids) <= 0 {
		return count, errors.New("orderids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update order set deleteStatus=2")
	sql_str.WriteString(" where orderid in(")
	question_mark := strings.Repeat("?,", len(orderids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	vals := make([]interface{}, 0, len(orderids))
	for _, v := range orderids {
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

// 根据【订单ID】查询【订单表】表中的单条记录
func (d *order) Get(orderid int) (order model.Order, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ orderid, pkgcode, ordercode, buyer, paytype, totalprice, payprice, paytime, isinvalid, isused, isrefund, refundprice, refundtime, remark, adder, addtime, moder, modtime, deleteStatus from order where orderid=?", orderid)
	if err != nil {
		return order, err
	}
	defer rows.Close()
	orders, err := d._RowsToArray(rows)
	if err != nil {
		return order, err
	}
	if len(orders) <= 0 {
		return order, err
	}
	return orders[0], nil
}

// 根据【订单ID】数组查询【订单表】表中的多条记录
func (d *order) GetIn(orderids []int) (orders []model.Order, err error) {
	if len(orderids) <= 0 {
		return orders, errors.New("orderids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ orderid, pkgcode, ordercode, buyer, paytype, totalprice, payprice, paytime, isinvalid, isused, isrefund, refundprice, refundtime, remark, adder, addtime, moder, modtime, deleteStatus from ")
	sql_str.WriteString("order")
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
		return orders, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【订单表】表总记录数
func (d *order) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from order")

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

// 查询【订单表】列表
func (d *order) GetRowList(pageIndex, pageSize int) (orders []model.Order, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ orderid, pkgcode, ordercode, buyer, paytype, totalprice, payprice, paytime, isinvalid, isused, isrefund, refundprice, refundtime, remark, adder, addtime, moder, modtime, deleteStatus from order")

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
		return orders, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【订单表】表记录
func (d *order) _RowsToArray(rows *sqlx.Rows) (orders []model.Order, err error) {
	for rows.Next() {
		m := model.Order{}
		err = rows.Scan(&m.Orderid, &m.Pkgcode, &m.Ordercode, &m.Buyer, &m.Paytype, &m.Totalprice, &m.Payprice, &m.Paytime, &m.Isinvalid, &m.Isused, &m.Isrefund, &m.Refundprice, &m.Refundtime, &m.Remark, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus)
		if err != nil {
			return orders, err
		}
		orders = append(orders, m)
	}
	return orders, err
}

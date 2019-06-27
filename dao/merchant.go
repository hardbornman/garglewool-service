package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

var MerchantDao = merchant{}

type merchant struct{}

// 根据【商家ID】查询【商家用户表】表中是否存在相关记录
func (d *merchant) Exist(merchantid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from merchant where merchantid=?", merchantid)
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

// 插入单条记录到【商家用户表】表中
func (d *merchant) Insert(m *model.Merchant) (int64, error) {
	result, err := garglewool.Exec("insert into merchant(merchantname,phone,deleteStatus,userroleid,loginpwd,loginaccount,nick,wechataccount,wechatsign,remark,lastlogintime,lastloginaddr,rigstertime,enable,addr) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", m.Merchantname, m.Phone, m.DeleteStatus, m.Userroleid, m.Loginpwd, m.Loginaccount, m.Nick, m.Wechataccount, m.Wechatsign, m.Remark, m.Lastlogintime, m.Lastloginaddr, m.Rigstertime, m.Enable, m.Addr)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【商家ID】修改【商家用户表】表的单条记录
func (d *merchant) Update(m *model.Merchant) (bool, error) {
	result, err := garglewool.Exec("update merchant set merchantname=?, phone=?, deleteStatus=?, userroleid=?, loginpwd=?, loginaccount=?, nick=?, wechataccount=?, wechatsign=?, remark=?, lastlogintime=?, lastloginaddr=?, rigstertime=?, enable=?, addr=? where merchantid=?", m.Merchantname, m.Phone, m.DeleteStatus, m.Userroleid, m.Loginpwd, m.Loginaccount, m.Nick, m.Wechataccount, m.Wechatsign, m.Remark, m.Lastlogintime, m.Lastloginaddr, m.Rigstertime, m.Enable, m.Addr, m.Merchantid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【商家ID】软删除【商家用户表】表中的单条记录
func (d *merchant) Delete(merchantid int) (bool, error) {
	result, err := garglewool.Exec("update merchant set deleteStatus=2 where merchantid=?", merchantid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【商家ID】数组软删除【商家用户表】表中的多条记录
func (d *merchant) DeleteIn(merchantids []int) (count int64, err error) {
	if len(merchantids) <= 0 {
		return count, errors.New("merchantids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update merchant set deleteStatus=2")
	sql_str.WriteString(" where merchantid in(")
	question_mark := strings.Repeat("?,", len(merchantids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	vals := make([]interface{}, 0, len(merchantids))
	for _, v := range merchantids {
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

// 根据【商家ID】查询【商家用户表】表中的单条记录
func (d *merchant) Get(merchantid int) (merchant model.Merchant, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ merchantid, merchantname, phone, createtime, updatetime, deleteStatus, userroleid, loginpwd, loginaccount, nick, wechataccount, wechatsign, remark, lastlogintime, lastloginaddr, rigstertime, enable, addr from merchant where merchantid=?", merchantid)
	if err != nil {
		return merchant, err
	}
	defer rows.Close()
	merchants, err := d._RowsToArray(rows)
	if err != nil {
		return merchant, err
	}
	if len(merchants) <= 0 {
		return merchant, err
	}
	return merchants[0], nil
}

// 根据【商家ID】数组查询【商家用户表】表中的多条记录
func (d *merchant) GetIn(merchantids []int) (merchants []model.Merchant, err error) {
	if len(merchantids) <= 0 {
		return merchants, errors.New("merchantids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ merchantid, merchantname, phone, createtime, updatetime, deleteStatus, userroleid, loginpwd, loginaccount, nick, wechataccount, wechatsign, remark, lastlogintime, lastloginaddr, rigstertime, enable, addr from ")
	sql_str.WriteString("merchant")
	sql_str.WriteString(" where merchantid in(")
	param_keys := strings.Repeat("?,", len(merchantids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(merchantids))
	for _, v := range merchantids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return merchants, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【商家用户表】表总记录数
func (d *merchant) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from merchant")

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

// 根据【用户角色ID】查询【商家用户表】总记录数
func (d *merchant) GetRowCountByUserroleid(userroleid int) (count int, err error) {
	rows, err := garglewool.Queryx("select count(0) Count from merchant where userroleid=? and deleteStatus = 1", userroleid)
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

// 查询【商家用户表】列表
func (d *merchant) GetRowList(pageIndex, pageSize int) (merchants []model.Merchant, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ merchantid, merchantname, phone, createtime, updatetime, deleteStatus, userroleid, loginpwd, loginaccount, nick, wechataccount, wechatsign, remark, lastlogintime, lastloginaddr, rigstertime, enable, addr from merchant")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by merchantid desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return merchants, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【商家用户表】表记录
func (d *merchant) _RowsToArray(rows *sqlx.Rows) (merchants []model.Merchant, err error) {
	for rows.Next() {
		m := model.Merchant{}
		err = rows.Scan(&m.Merchantid, &m.Merchantname, &m.Phone, &m.Createtime, &m.Updatetime, &m.DeleteStatus, &m.Userroleid, &m.Loginpwd, &m.Loginaccount, &m.Nick, &m.Wechataccount, &m.Wechatsign, &m.Remark, &m.Lastlogintime, &m.Lastloginaddr, &m.Rigstertime, &m.Enable, &m.Addr)
		if err != nil {
			return merchants, err
		}
		merchants = append(merchants, m)
	}
	return merchants, err
}

// 查询【用户表】表中是否存在相关记录
func (d *merchant) Login(loginAccount string, loginPwd string) (merchant model.Merchant, err error) {
	rows, err := garglewool.Queryx("select *  from merchant where loginaccount=? and loginpwd=?", loginAccount, loginPwd)
	if err != nil {
		return merchant, err
	}
	defer rows.Close()
	merchants, err := d._RowsToArray(rows)
	if err != nil {
		return merchant, err
	}
	if len(merchants) <= 0 {
		return merchant, err
	}
	return merchants[0], nil
}

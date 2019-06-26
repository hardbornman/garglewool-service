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

var GwShopDao = gwShop{}

type gwShop struct{}

// 根据【店铺ID】查询【店铺管理】表中是否存在相关记录
func (d *gwShop) Exist(shopid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from gw_shop where shopid=?", shopid)
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

// 插入单条记录到【店铺管理】表中
func (d *gwShop) Insert(m *model.GwShop) (int64, error) {
	result, err := garglewool.Exec("insert into gw_shop(shop_code,shop_name,shop_province,shop_city,shop_district,shop_address,shop_phone,shop_createtime,shop_exittime,adder,addtime,moder,modtime,deleteStatus) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)", m.ShopCode, m.ShopName, m.ShopProvince, m.ShopCity, m.ShopDistrict, m.ShopAddress, m.ShopPhone, m.ShopCreatetime, m.ShopExittime, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【店铺ID】修改【店铺管理】表的单条记录
func (d *gwShop) Update(m *model.GwShop) (bool, error) {
	result, err := garglewool.Exec("update gw_shop set shop_code=?, shop_name=?, shop_province=?, shop_city=?, shop_district=?, shop_address=?, shop_phone=?, shop_createtime=?, shop_exittime=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=? where shopid=?", m.ShopCode, m.ShopName, m.ShopProvince, m.ShopCity, m.ShopDistrict, m.ShopAddress, m.ShopPhone, m.ShopCreatetime, m.ShopExittime, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Shopid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【店铺ID】软删除【店铺管理】表中的单条记录
func (d *gwShop) Delete(shopid int) (bool, error) {
	result, err := garglewool.Exec("update gw_shop set deleteStatus=2 where shopid=?", shopid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【店铺ID】数组软删除【店铺管理】表中的多条记录
func (d *gwShop) DeleteIn(shopids []int) (count int64, err error) {
	if len(shopids) <= 0 {
		return count, errors.New("shopids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update gw_shop set deleteStatus=2")
	sql_str.WriteString(" where shopid in(")
	question_mark := strings.Repeat("?,", len(shopids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	var result sql.Result
	vals := make([]interface{}, 0, len(shopids))
	for _, v := range shopids {
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

// 根据【店铺ID】查询【店铺管理】表中的单条记录
func (d *gwShop) Get(shopid int) (gwShop model.GwShop, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ shopid, shop_code, shop_name, shop_province, shop_city, shop_district, shop_address, shop_phone, shop_createtime, shop_exittime, adder, addtime, moder, modtime, deleteStatus from gw_shop where shopid=?", shopid)
	if err != nil {
		return gwShop, err
	}
	defer rows.Close()
	gwShops, err := d._RowsToArray(rows)
	if err != nil {
		return gwShop, err
	}
	if len(gwShops) <= 0 {
		return gwShop, err
	}
	return gwShops[0], nil
}

// 根据【店铺ID】数组查询【店铺管理】表中的多条记录
func (d *gwShop) GetIn(shopids []int) (gwShops []model.GwShop, err error) {
	if len(shopids) <= 0 {
		return gwShops, errors.New("shopids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ shopid, shop_code, shop_name, shop_province, shop_city, shop_district, shop_address, shop_phone, shop_createtime, shop_exittime, adder, addtime, moder, modtime, deleteStatus from ")
	sql_str.WriteString("gw_shop")
	sql_str.WriteString(" where shopid in(")
	param_keys := strings.Repeat("?,", len(shopids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(shopids))
	for _, v := range shopids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return gwShops, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【店铺管理】表总记录数
func (d *gwShop) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from gw_shop")

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

// 查询【店铺管理】列表
func (d *gwShop) GetRowList(pageIndex, pageSize int) (gwShops []model.GwShop, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ shopid, shop_code, shop_name, shop_province, shop_city, shop_district, shop_address, shop_phone, shop_createtime, shop_exittime, adder, addtime, moder, modtime, deleteStatus from gw_shop")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by shopid desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return gwShops, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【店铺管理】表记录
func (d *gwShop) _RowsToArray(rows *sqlx.Rows) (gwShops []model.GwShop, err error) {
	for rows.Next() {
		m := model.GwShop{}
		err = rows.Scan(&m.Shopid, &m.ShopCode, &m.ShopName, &m.ShopProvince, &m.ShopCity, &m.ShopDistrict, &m.ShopAddress, &m.ShopPhone, &m.ShopCreatetime, &m.ShopExittime, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus)
		if err != nil {
			return gwShops, err
		}
		gwShops = append(gwShops, m)
	}
	return gwShops, err
}

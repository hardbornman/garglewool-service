package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

var ShopDao = shop{}

type shop struct{}

// 根据【店铺ID】查询【店铺表】表中是否存在相关记录
func (d *shop) Exist(shopid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from shop where shopid=?", shopid)
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

// 插入单条记录到【店铺表】表中
func (d *shop) Insert(m *model.Shop) (int64, error) {
	result, err := garglewool.Exec("insert into shop(shopcode,shopname,province,city,district,address,phone,leaguetime,exittime,adder,addtime,moder,modtime,deleteStatus,merchantid,longtitude,latitude) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", m.Shopcode, m.Shopname, m.Province, m.City, m.District, m.Address, m.Phone, m.Leaguetime, m.Exittime, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Merchantid, m.Longtitude, m.Latitude)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【店铺ID】修改【店铺表】表的单条记录
func (d *shop) Update(m *model.Shop) (bool, error) {
	result, err := garglewool.Exec("update shop set shopcode=?, shopname=?, province=?, city=?, district=?, address=?, phone=?, leaguetime=?, exittime=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=?, merchantid=?, longtitude=?, latitude=? where shopid=?", m.Shopcode, m.Shopname, m.Province, m.City, m.District, m.Address, m.Phone, m.Leaguetime, m.Exittime, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Merchantid, m.Longtitude, m.Latitude, m.Shopid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【店铺ID】软删除【店铺表】表中的单条记录
func (d *shop) Delete(shopid int) (bool, error) {
	result, err := garglewool.Exec("update shop set deleteStatus=2 where shopid=?", shopid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【店铺ID】数组软删除【店铺表】表中的多条记录
func (d *shop) DeleteIn(shopids []int) (count int64, err error) {
	if len(shopids) <= 0 {
		return count, errors.New("shopids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update shop set deleteStatus=2")
	sql_str.WriteString(" where shopid in(")
	question_mark := strings.Repeat("?,", len(shopids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	vals := make([]interface{}, 0, len(shopids))
	for _, v := range shopids {
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

// 根据【店铺ID】查询【店铺表】表中的单条记录
func (d *shop) Get(shopid int) (shop model.Shop, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ shopid, shopcode, shopname, province, city, district, address, phone, leaguetime, exittime, adder, addtime, moder, modtime, deleteStatus, merchantid, longtitude, latitude from shop where shopid=?", shopid)
	if err != nil {
		return shop, err
	}
	defer rows.Close()
	shops, err := d._RowsToArray(rows)
	if err != nil {
		return shop, err
	}
	if len(shops) <= 0 {
		return shop, err
	}
	return shops[0], nil
}

// 根据【店铺ID】数组查询【店铺表】表中的多条记录
func (d *shop) GetIn(shopids []int) (shops []model.Shop, err error) {
	if len(shopids) <= 0 {
		return shops, errors.New("shopids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ shopid, shopcode, shopname, province, city, district, address, phone, leaguetime, exittime, adder, addtime, moder, modtime, deleteStatus, merchantid, longtitude, latitude from ")
	sql_str.WriteString("shop")
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
		return shops, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【店铺表】表总记录数
//func (d *shop) GetRowCount() (count int, err error) {
//	sqlString := golibs.NewStringBuilder()
//	params := make([]interface{}, 0)
//	conditions := 0
//
//	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from shop")
//
//	//region 处理deleteStatus
//	if conditions > 0 {
//		sqlString.Append(" and ")
//	} else {
//		sqlString.Append(" where ")
//	}
//	sqlString.Append("deleteStatus = 1")
//	//endregion
//
//	//region Query
//	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
//	if err != nil {
//		return -1, err
//	}
//	defer rows.Close()
//	if rows.Next() {
//		err = rows.Scan(&count)
//		if err != nil {
//			return -1, err
//		}
//		return count, nil
//	}
//	//endregion
//
//	return -1, nil
//}

// 根据【商家ID】查询【店铺表】总记录数
func (d *shop) GetRowCountByMerchantid(merchantid int) (count int, err error) {
	rows, err := garglewool.Queryx("select count(0) Count from shop where merchantid=? and deleteStatus = 1", merchantid)
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

// 查询【店铺表】列表
//func (d *shop) GetRowList(pageIndex, pageSize int) (shops []model.Shop, err error) {
//	sqlString := golibs.NewStringBuilder()
//	params := make([]interface{}, 0)
//	conditions := 0
//
//	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ shopid, shopcode, shopname, province, city, district, address, phone, leaguetime, exittime, adder, addtime, moder, modtime, deleteStatus, merchantid, longtitude, latitude from shop")
//
//	//region 处理deleteStatus
//	if conditions > 0 {
//		sqlString.Append(" and ")
//	} else {
//		sqlString.Append(" where ")
//	}
//	sqlString.Append("deleteStatus = 1")
//	//endregion
//
//	//region order by
//	sqlString.Append(" order by shopid desc")
//	//endregion
//
//	//region limit
//	sqlString.Append(" limit ?,?")
//	params = append(params, (pageIndex-1)*pageSize)
//	params = append(params, pageSize)
//	//endregion
//
//	//region Query
//	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
//	if err != nil {
//		return shops, err
//	}
//	defer rows.Close()
//	return d._RowsToArray(rows)
//	//endregion
//}
// 查询【店铺表】表总记录数,使用条件：【店铺ID】,【店铺名称】,【店铺详细地址】,【加盟平台日期】
func (d *shop) GetRowCount(shopid int, shopname string, address string, leaguetime time.Time) (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from shop")

	//region 处理shopid
	if shopid > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("shopid=?")
		params = append(params, shopid)
		conditions++
	}
	//endregion

	//region 处理shopname
	if golibs.Length(shopname) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("shopname=?")
		params = append(params, shopname)
		conditions++
	}
	//endregion

	//region 处理address
	if golibs.Length(address) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("address=?")
		params = append(params, address)
		conditions++
	}
	//endregion

	//region 处理leaguetime
	if leaguetime.After(time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local)) {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("leaguetime=?")
		params = append(params, leaguetime)
		conditions++
	}
	//endregion

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

// 查询【店铺表】列表,使用条件：【店铺ID】,【店铺名称】,【店铺详细地址】,【加盟平台日期】
func (d *shop) GetRowList(shopid int, shopname string, address string, leaguetime time.Time, pageIndex, pageSize int) (shops []model.Shop, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ shopid, shopcode, shopname, province, city, district, address, phone, leaguetime, exittime, adder, addtime, moder, modtime, deleteStatus, merchantid, longtitude, latitude from shop")

	//region 处理shopid
	if shopid > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("shopid=?")
		params = append(params, shopid)
		conditions++
	}
	//endregion

	//region 处理shopname
	if golibs.Length(shopname) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("shopname=?")
		params = append(params, shopname)
		conditions++
	}
	//endregion

	//region 处理address
	if golibs.Length(address) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("address=?")
		params = append(params, address)
		conditions++
	}
	//endregion

	//region 处理leaguetime
	if leaguetime.After(time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local)) {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("leaguetime=?")
		params = append(params, leaguetime)
		conditions++
	}
	//endregion

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
		return shops, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【店铺表】表记录
func (d *shop) _RowsToArray(rows *sqlx.Rows) (shops []model.Shop, err error) {
	for rows.Next() {
		m := model.Shop{}
		err = rows.Scan(&m.Shopid, &m.Shopcode, &m.Shopname, &m.Province, &m.City, &m.District, &m.Address, &m.Phone, &m.Leaguetime, &m.Exittime, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus, &m.Merchantid, &m.Longtitude, &m.Latitude)
		if err != nil {
			return shops, err
		}
		shops = append(shops, m)
	}
	return shops, err
}

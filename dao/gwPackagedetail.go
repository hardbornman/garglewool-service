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

var GwPackagedetailDao = gwPackagedetail{}

type gwPackagedetail struct{}

// 根据【套餐明细ID】查询【套餐明细管理】表中是否存在相关记录
func (d *gwPackagedetail) Exist(packagedetailid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from gw_packagedetail where packagedetailid=?", packagedetailid)
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

// 插入单条记录到【套餐明细管理】表中
func (d *gwPackagedetail) Insert(m *model.GwPackagedetail) (int64, error) {
	result, err := garglewool.Exec("insert into gw_packagedetail(pkg_code,pkgdetail_name,pkgdetail_nums,pkgdetail_price,adder,addtime,moder,modtime,deleteStatus) values(?,?,?,?,?,?,?,?,?)", m.PkgCode, m.PkgdetailName, m.PkgdetailNums, m.PkgdetailPrice, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【套餐明细ID】修改【套餐明细管理】表的单条记录
func (d *gwPackagedetail) Update(m *model.GwPackagedetail) (bool, error) {
	result, err := garglewool.Exec("update gw_packagedetail set pkg_code=?, pkgdetail_name=?, pkgdetail_nums=?, pkgdetail_price=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=? where packagedetailid=?", m.PkgCode, m.PkgdetailName, m.PkgdetailNums, m.PkgdetailPrice, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Packagedetailid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【套餐明细ID】软删除【套餐明细管理】表中的单条记录
func (d *gwPackagedetail) Delete(packagedetailid int) (bool, error) {
	result, err := garglewool.Exec("update gw_packagedetail set deleteStatus=2 where packagedetailid=?", packagedetailid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【套餐明细ID】数组软删除【套餐明细管理】表中的多条记录
func (d *gwPackagedetail) DeleteIn(packagedetailids []int) (count int64, err error) {
	if len(packagedetailids) <= 0 {
		return count, errors.New("packagedetailids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update gw_packagedetail set deleteStatus=2")
	sql_str.WriteString(" where packagedetailid in(")
	question_mark := strings.Repeat("?,", len(packagedetailids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	var result sql.Result
	vals := make([]interface{}, 0, len(packagedetailids))
	for _, v := range packagedetailids {
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

// 根据【套餐明细ID】查询【套餐明细管理】表中的单条记录
func (d *gwPackagedetail) Get(packagedetailid int) (gwPackagedetail model.GwPackagedetail, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ packagedetailid, pkg_code, pkgdetail_name, pkgdetail_nums, pkgdetail_price, adder, addtime, moder, modtime, deleteStatus from gw_packagedetail where packagedetailid=?", packagedetailid)
	if err != nil {
		return gwPackagedetail, err
	}
	defer rows.Close()
	gwPackagedetails, err := d._RowsToArray(rows)
	if err != nil {
		return gwPackagedetail, err
	}
	if len(gwPackagedetails) <= 0 {
		return gwPackagedetail, err
	}
	return gwPackagedetails[0], nil
}

// 根据【套餐明细ID】数组查询【套餐明细管理】表中的多条记录
func (d *gwPackagedetail) GetIn(packagedetailids []int) (gwPackagedetails []model.GwPackagedetail, err error) {
	if len(packagedetailids) <= 0 {
		return gwPackagedetails, errors.New("packagedetailids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ packagedetailid, pkg_code, pkgdetail_name, pkgdetail_nums, pkgdetail_price, adder, addtime, moder, modtime, deleteStatus from ")
	sql_str.WriteString("gw_packagedetail")
	sql_str.WriteString(" where packagedetailid in(")
	param_keys := strings.Repeat("?,", len(packagedetailids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(packagedetailids))
	for _, v := range packagedetailids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return gwPackagedetails, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【套餐明细管理】表总记录数
func (d *gwPackagedetail) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from gw_packagedetail")

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

// 查询【套餐明细管理】列表
func (d *gwPackagedetail) GetRowList(pageIndex, pageSize int) (gwPackagedetails []model.GwPackagedetail, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ packagedetailid, pkg_code, pkgdetail_name, pkgdetail_nums, pkgdetail_price, adder, addtime, moder, modtime, deleteStatus from gw_packagedetail")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by packagedetailid desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return gwPackagedetails, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【套餐明细管理】表记录
func (d *gwPackagedetail) _RowsToArray(rows *sqlx.Rows) (gwPackagedetails []model.GwPackagedetail, err error) {
	for rows.Next() {
		m := model.GwPackagedetail{}
		err = rows.Scan(&m.Packagedetailid, &m.PkgCode, &m.PkgdetailName, &m.PkgdetailNums, &m.PkgdetailPrice, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus)
		if err != nil {
			return gwPackagedetails, err
		}
		gwPackagedetails = append(gwPackagedetails, m)
	}
	return gwPackagedetails, err
}

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

var GwPackageDao = gwPackage{}

type gwPackage struct{}

// 根据【套餐ID】查询【套餐管理】表中是否存在相关记录
func (d *gwPackage) Exist(packageid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from gw_package where packageid=?", packageid)
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

// 插入单条记录到【套餐管理】表中
func (d *gwPackage) Insert(m *model.GwPackage) (int64, error) {
	result, err := garglewool.Exec("insert into gw_package(shop_code,pkg_code,pkg_type,pkg_title,pkg_people,pkg_isorder,pkg_isrefund,pkg_isinhouse,pkg_isnew,pkg_isrecommend,pkg_validdays,pkg_follows,pkg_createtime,pkg_validtime,pkg_exittime,pkg_links,pkg_info,adder,addtime,moder,modtime,deleteStatus) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", m.ShopCode, m.PkgCode, m.PkgType, m.PkgTitle, m.PkgPeople, m.PkgIsorder, m.PkgIsrefund, m.PkgIsinhouse, m.PkgIsnew, m.PkgIsrecommend, m.PkgValiddays, m.PkgFollows, m.PkgCreatetime, m.PkgValidtime, m.PkgExittime, m.PkgLinks, m.PkgInfo, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【套餐ID】修改【套餐管理】表的单条记录
func (d *gwPackage) Update(m *model.GwPackage) (bool, error) {
	result, err := garglewool.Exec("update gw_package set shop_code=?, pkg_code=?, pkg_type=?, pkg_title=?, pkg_people=?, pkg_isorder=?, pkg_isrefund=?, pkg_isinhouse=?, pkg_isnew=?, pkg_isrecommend=?, pkg_validdays=?, pkg_follows=?, pkg_createtime=?, pkg_validtime=?, pkg_exittime=?, pkg_links=?, pkg_info=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=? where packageid=?", m.ShopCode, m.PkgCode, m.PkgType, m.PkgTitle, m.PkgPeople, m.PkgIsorder, m.PkgIsrefund, m.PkgIsinhouse, m.PkgIsnew, m.PkgIsrecommend, m.PkgValiddays, m.PkgFollows, m.PkgCreatetime, m.PkgValidtime, m.PkgExittime, m.PkgLinks, m.PkgInfo, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Packageid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【套餐ID】软删除【套餐管理】表中的单条记录
func (d *gwPackage) Delete(packageid int) (bool, error) {
	result, err := garglewool.Exec("update gw_package set deleteStatus=2 where packageid=?", packageid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【套餐ID】数组软删除【套餐管理】表中的多条记录
func (d *gwPackage) DeleteIn(packageids []int) (count int64, err error) {
	if len(packageids) <= 0 {
		return count, errors.New("packageids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update gw_package set deleteStatus=2")
	sql_str.WriteString(" where packageid in(")
	question_mark := strings.Repeat("?,", len(packageids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	var result sql.Result
	vals := make([]interface{}, 0, len(packageids))
	for _, v := range packageids {
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

// 根据【套餐ID】查询【套餐管理】表中的单条记录
func (d *gwPackage) Get(packageid int) (gwPackage model.GwPackage, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ packageid, shop_code, pkg_code, pkg_type, pkg_title, pkg_people, pkg_isorder, pkg_isrefund, pkg_isinhouse, pkg_isnew, pkg_isrecommend, pkg_validdays, pkg_follows, pkg_createtime, pkg_validtime, pkg_exittime, pkg_links, pkg_info, adder, addtime, moder, modtime, deleteStatus from gw_package where packageid=?", packageid)
	if err != nil {
		return gwPackage, err
	}
	defer rows.Close()
	gwPackages, err := d._RowsToArray(rows)
	if err != nil {
		return gwPackage, err
	}
	if len(gwPackages) <= 0 {
		return gwPackage, err
	}
	return gwPackages[0], nil
}

// 根据【套餐ID】数组查询【套餐管理】表中的多条记录
func (d *gwPackage) GetIn(packageids []int) (gwPackages []model.GwPackage, err error) {
	if len(packageids) <= 0 {
		return gwPackages, errors.New("packageids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ packageid, shop_code, pkg_code, pkg_type, pkg_title, pkg_people, pkg_isorder, pkg_isrefund, pkg_isinhouse, pkg_isnew, pkg_isrecommend, pkg_validdays, pkg_follows, pkg_createtime, pkg_validtime, pkg_exittime, pkg_links, pkg_info, adder, addtime, moder, modtime, deleteStatus from ")
	sql_str.WriteString("gw_package")
	sql_str.WriteString(" where packageid in(")
	param_keys := strings.Repeat("?,", len(packageids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(packageids))
	for _, v := range packageids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return gwPackages, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【套餐管理】表总记录数
func (d *gwPackage) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from gw_package")

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

// 查询【套餐管理】列表
func (d *gwPackage) GetRowList(pageIndex, pageSize int) (gwPackages []model.GwPackage, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ packageid, shop_code, pkg_code, pkg_type, pkg_title, pkg_people, pkg_isorder, pkg_isrefund, pkg_isinhouse, pkg_isnew, pkg_isrecommend, pkg_validdays, pkg_follows, pkg_createtime, pkg_validtime, pkg_exittime, pkg_links, pkg_info, adder, addtime, moder, modtime, deleteStatus from gw_package")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by packageid desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return gwPackages, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【套餐管理】表记录
func (d *gwPackage) _RowsToArray(rows *sqlx.Rows) (gwPackages []model.GwPackage, err error) {
	for rows.Next() {
		m := model.GwPackage{}
		err = rows.Scan(&m.Packageid, &m.ShopCode, &m.PkgCode, &m.PkgType, &m.PkgTitle, &m.PkgPeople, &m.PkgIsorder, &m.PkgIsrefund, &m.PkgIsinhouse, &m.PkgIsnew, &m.PkgIsrecommend, &m.PkgValiddays, &m.PkgFollows, &m.PkgCreatetime, &m.PkgValidtime, &m.PkgExittime, &m.PkgLinks, &m.PkgInfo, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus)
		if err != nil {
			return gwPackages, err
		}
		gwPackages = append(gwPackages, m)
	}
	return gwPackages, err
}

package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

var SProvinceDao = sProvince{}

type sProvince struct{}

// 根据【省ID】查询【省，直辖市】表中是否存在相关记录
func (d *sProvince) Exist(provinceId int64) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from s_province where provinceId=?", provinceId)
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

// 插入单条记录到【省，直辖市】表中
func (d *sProvince) Insert(m *model.SProvince) (bool, error) {
	result, err := garglewool.Exec("insert into s_province(provinceId,provinceName,dateCreated,dateUpdated) values(?,?,?,?)", m.ProvinceId, m.ProvinceName, m.DateCreated, m.DateUpdated)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【省ID】修改【省，直辖市】表的单条记录
func (d *sProvince) Update(m *model.SProvince) (bool, error) {
	result, err := garglewool.Exec("update s_province set provinceName=?, dateCreated=?, dateUpdated=? where provinceId=?", m.ProvinceName, m.DateCreated, m.DateUpdated, m.ProvinceId)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 插入或修改【省，直辖市】表的单条记录
func (d *sProvince) InsertUpdate(m *model.SProvince) (bool, error) {
	result, err := garglewool.Exec("insert into s_province(provinceId,provinceName,dateCreated,dateUpdated) values(?,?,?,?) ON DUPLICATE KEY UPDATE provinceName=?,dateCreated=?,dateUpdated=?", m.ProvinceId, m.ProvinceName, m.DateCreated, m.DateUpdated, m.ProvinceName, m.DateCreated, m.DateUpdated)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【省ID】查询【省，直辖市】表中的单条记录
func (d *sProvince) Get(provinceId int64) (sProvince model.SProvince, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ provinceId, provinceName, dateCreated, dateUpdated from s_province where provinceId=?", provinceId)
	if err != nil {
		return sProvince, err
	}
	defer rows.Close()
	sProvinces, err := d._RowsToArray(rows)
	if err != nil {
		return sProvince, err
	}
	if len(sProvinces) <= 0 {
		return sProvince, err
	}
	return sProvinces[0], nil
}

// 根据【省ID】数组查询【省，直辖市】表中的多条记录
func (d *sProvince) GetIn(provinceIds []int64) (sProvinces []model.SProvince, err error) {
	if len(provinceIds) <= 0 {
		return sProvinces, errors.New("provinceIds is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ provinceId, provinceName, dateCreated, dateUpdated from ")
	sql_str.WriteString("s_province")
	sql_str.WriteString(" where provinceId in(")
	param_keys := strings.Repeat("?,", len(provinceIds))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(provinceIds))
	for _, v := range provinceIds {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return sProvinces, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【省，直辖市】表总记录数,使用条件：【省名称】
func (d *sProvince) GetRowCount(provinceName string) (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from s_province")

	//region 处理provinceName
	if golibs.Length(provinceName) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("provinceName=?")
		params = append(params, provinceName)
		conditions++
	}
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

// 查询【省，直辖市】列表,使用条件：【省名称】
func (d *sProvince) GetRowList(provinceName string, pageIndex, pageSize int) (sProvinces []model.SProvince, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ provinceId, provinceName, dateCreated, dateUpdated from s_province")

	//region 处理provinceName
	if golibs.Length(provinceName) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("provinceName=?")
		params = append(params, provinceName)
		conditions++
	}
	//endregion

	//region order by
	sqlString.Append(" order by provinceId desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return sProvinces, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【省，直辖市】表记录
func (d *sProvince) _RowsToArray(rows *sqlx.Rows) (sProvinces []model.SProvince, err error) {
	for rows.Next() {
		m := model.SProvince{}
		err = rows.Scan(&m.ProvinceId, &m.ProvinceName, &m.DateCreated, &m.DateUpdated)
		if err != nil {
			return sProvinces, err
		}
		sProvinces = append(sProvinces, m)
	}
	return sProvinces, err
}

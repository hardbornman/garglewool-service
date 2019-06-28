package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

var SDistrictDao = sDistrict{}

type sDistrict struct{}

// 根据【区ID】查询【区县】表中是否存在相关记录
func (d *sDistrict) Exist(districtId int64) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from s_district where districtId=?", districtId)
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

// 插入单条记录到【区县】表中
func (d *sDistrict) Insert(m *model.SDistrict) (bool, error) {
	result, err := garglewool.Exec("insert into s_district(districtId,districtName,cityId,dateCreated,dateUpdated) values(?,?,?,?,?)", m.DistrictId, m.DistrictName, m.CityId, m.DateCreated, m.DateUpdated)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【区ID】修改【区县】表的单条记录
func (d *sDistrict) Update(m *model.SDistrict) (bool, error) {
	result, err := garglewool.Exec("update s_district set districtName=?, cityId=?, dateCreated=?, dateUpdated=? where districtId=?", m.DistrictName, m.CityId, m.DateCreated, m.DateUpdated, m.DistrictId)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 插入或修改【区县】表的单条记录
func (d *sDistrict) InsertUpdate(m *model.SDistrict) (bool, error) {
	result, err := garglewool.Exec("insert into s_district(districtId,districtName,cityId,dateCreated,dateUpdated) values(?,?,?,?,?) ON DUPLICATE KEY UPDATE districtName=?,cityId=?,dateCreated=?,dateUpdated=?", m.DistrictId, m.DistrictName, m.CityId, m.DateCreated, m.DateUpdated, m.DistrictName, m.CityId, m.DateCreated, m.DateUpdated)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【区ID】查询【区县】表中的单条记录
func (d *sDistrict) Get(districtId int64) (sDistrict model.SDistrict, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ districtId, districtName, cityId, dateCreated, dateUpdated from s_district where districtId=?", districtId)
	if err != nil {
		return sDistrict, err
	}
	defer rows.Close()
	sDistricts, err := d._RowsToArray(rows)
	if err != nil {
		return sDistrict, err
	}
	if len(sDistricts) <= 0 {
		return sDistrict, err
	}
	return sDistricts[0], nil
}

// 根据【区ID】数组查询【区县】表中的多条记录
func (d *sDistrict) GetIn(districtIds []int64) (sDistricts []model.SDistrict, err error) {
	if len(districtIds) <= 0 {
		return sDistricts, errors.New("districtIds is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ districtId, districtName, cityId, dateCreated, dateUpdated from ")
	sql_str.WriteString("s_district")
	sql_str.WriteString(" where districtId in(")
	param_keys := strings.Repeat("?,", len(districtIds))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(districtIds))
	for _, v := range districtIds {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return sDistricts, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【区县】表总记录数,使用条件：【区名称】,【市ID】
func (d *sDistrict) GetRowCount(districtName string, cityId int64) (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from s_district")

	//region 处理districtName
	if golibs.Length(districtName) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("districtName=?")
		params = append(params, districtName)
		conditions++
	}
	//endregion

	//region 处理cityId
	if cityId > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("cityId=?")
		params = append(params, cityId)
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

// 根据【市ID】查询【区县】总记录数
func (d *sDistrict) GetRowCountByCityId(cityId int64) (count int, err error) {
	rows, err := garglewool.Queryx("select count(0) Count from s_district where cityId=?", cityId)
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

// 查询【区县】列表,使用条件：【区名称】,【市ID】
func (d *sDistrict) GetRowList(districtName string, cityId int64, pageIndex, pageSize int) (sDistricts []model.SDistrict, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ districtId, districtName, cityId, dateCreated, dateUpdated from s_district")

	//region 处理districtName
	if golibs.Length(districtName) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("districtName=?")
		params = append(params, districtName)
		conditions++
	}
	//endregion

	//region 处理cityId
	if cityId > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("cityId=?")
		params = append(params, cityId)
		conditions++
	}
	//endregion

	//region order by
	sqlString.Append(" order by districtId desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return sDistricts, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【区县】表记录
func (d *sDistrict) _RowsToArray(rows *sqlx.Rows) (sDistricts []model.SDistrict, err error) {
	for rows.Next() {
		m := model.SDistrict{}
		err = rows.Scan(&m.DistrictId, &m.DistrictName, &m.CityId, &m.DateCreated, &m.DateUpdated)
		if err != nil {
			return sDistricts, err
		}
		sDistricts = append(sDistricts, m)
	}
	return sDistricts, err
}

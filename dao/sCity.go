package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

var SCityDao = sCity{}

type sCity struct{}

// 根据【市ID】查询【市】表中是否存在相关记录
func (d *sCity) Exist(cityId int64) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from s_city where cityId=?", cityId)
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

// 插入单条记录到【市】表中
func (d *sCity) Insert(m *model.SCity) (bool, error) {
	result, err := garglewool.Exec("insert into s_city(cityId,cityName,zipCode,provinceId,dateCreated,dateUpdated) values(?,?,?,?,?,?)", m.CityId, m.CityName, m.ZipCode, m.ProvinceId, m.DateCreated, m.DateUpdated)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【市ID】修改【市】表的单条记录
func (d *sCity) Update(m *model.SCity) (bool, error) {
	result, err := garglewool.Exec("update s_city set cityName=?, zipCode=?, provinceId=?, dateCreated=?, dateUpdated=? where cityId=?", m.CityName, m.ZipCode, m.ProvinceId, m.DateCreated, m.DateUpdated, m.CityId)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 插入或修改【市】表的单条记录
func (d *sCity) InsertUpdate(m *model.SCity) (bool, error) {
	result, err := garglewool.Exec("insert into s_city(cityId,cityName,zipCode,provinceId,dateCreated,dateUpdated) values(?,?,?,?,?,?) ON DUPLICATE KEY UPDATE cityName=?,zipCode=?,provinceId=?,dateCreated=?,dateUpdated=?", m.CityId, m.CityName, m.ZipCode, m.ProvinceId, m.DateCreated, m.DateUpdated, m.CityName, m.ZipCode, m.ProvinceId, m.DateCreated, m.DateUpdated)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【市ID】查询【市】表中的单条记录
func (d *sCity) Get(cityId int64) (sCity model.SCity, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ cityId, cityName, zipCode, provinceId, dateCreated, dateUpdated from s_city where cityId=?", cityId)
	if err != nil {
		return sCity, err
	}
	defer rows.Close()
	sCitys, err := d._RowsToArray(rows)
	if err != nil {
		return sCity, err
	}
	if len(sCitys) <= 0 {
		return sCity, err
	}
	return sCitys[0], nil
}

// 根据【市ID】数组查询【市】表中的多条记录
func (d *sCity) GetIn(cityIds []int64) (sCitys []model.SCity, err error) {
	if len(cityIds) <= 0 {
		return sCitys, errors.New("cityIds is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ cityId, cityName, zipCode, provinceId, dateCreated, dateUpdated from ")
	sql_str.WriteString("s_city")
	sql_str.WriteString(" where cityId in(")
	param_keys := strings.Repeat("?,", len(cityIds))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(cityIds))
	for _, v := range cityIds {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return sCitys, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【市】表总记录数,使用条件：【市名称】,【省ID】
func (d *sCity) GetRowCount(cityName string, provinceId int64) (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from s_city")

	//region 处理cityName
	if golibs.Length(cityName) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("cityName=?")
		params = append(params, cityName)
		conditions++
	}
	//endregion

	//region 处理provinceId
	if provinceId > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("provinceId=?")
		params = append(params, provinceId)
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

// 根据【省ID】查询【市】总记录数
func (d *sCity) GetRowCountByProvinceId(provinceId int64) (count int, err error) {
	rows, err := garglewool.Queryx("select count(0) Count from s_city where provinceId=?", provinceId)
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

// 查询【市】列表,使用条件：【市名称】,【省ID】
func (d *sCity) GetRowList(cityName string, provinceId int64, pageIndex, pageSize int) (sCitys []model.SCity, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ cityId, cityName, zipCode, provinceId, dateCreated, dateUpdated from s_city")

	//region 处理cityName
	if golibs.Length(cityName) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("cityName=?")
		params = append(params, cityName)
		conditions++
	}
	//endregion

	//region 处理provinceId
	if provinceId > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("provinceId=?")
		params = append(params, provinceId)
		conditions++
	}
	//endregion

	//region order by
	sqlString.Append(" order by cityId desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return sCitys, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【市】表记录
func (d *sCity) _RowsToArray(rows *sqlx.Rows) (sCitys []model.SCity, err error) {
	for rows.Next() {
		m := model.SCity{}
		err = rows.Scan(&m.CityId, &m.CityName, &m.ZipCode, &m.ProvinceId, &m.DateCreated, &m.DateUpdated)
		if err != nil {
			return sCitys, err
		}
		sCitys = append(sCitys, m)
	}
	return sCitys, err
}

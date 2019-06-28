package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

var AreaDao = area{}

type area struct{}

// 根据【序号】查询【区域表】表中是否存在相关记录
func (d *area) Exist(areaid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from area where areaid=?", areaid)
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

// 插入单条记录到【区域表】表中
func (d *area) Insert(m *model.Area) (bool, error) {
	result, err := garglewool.Exec("insert into area(areaid,regionName,regionCode,parentCode,longitude,latitude) values(?,?,?,?,?,?)", m.Areaid, m.RegionName, m.RegionCode, m.ParentCode, m.Longitude, m.Latitude)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【序号】修改【区域表】表的单条记录
func (d *area) Update(m *model.Area) (bool, error) {
	result, err := garglewool.Exec("update area set regionName=?, regionCode=?, parentCode=?, longitude=?, latitude=? where areaid=?", m.RegionName, m.RegionCode, m.ParentCode, m.Longitude, m.Latitude, m.Areaid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 插入或修改【区域表】表的单条记录
func (d *area) InsertUpdate(m *model.Area) (bool, error) {
	result, err := garglewool.Exec("insert into area(areaid,regionName,regionCode,parentCode,longitude,latitude) values(?,?,?,?,?,?) ON DUPLICATE KEY UPDATE regionName=?,regionCode=?,parentCode=?,longitude=?,latitude=?", m.Areaid, m.RegionName, m.RegionCode, m.ParentCode, m.Longitude, m.Latitude, m.RegionName, m.RegionCode, m.ParentCode, m.Longitude, m.Latitude)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【序号】查询【区域表】表中的单条记录
func (d *area) Get(areaid int) (area model.Area, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ areaid, regionName, regionCode, parentCode, longitude, latitude from area where areaid=?", areaid)
	if err != nil {
		return area, err
	}
	defer rows.Close()
	areas, err := d._RowsToArray(rows)
	if err != nil {
		return area, err
	}
	if len(areas) <= 0 {
		return area, err
	}
	return areas[0], nil
}

// 根据【序号】数组查询【区域表】表中的多条记录
func (d *area) GetIn(areaids []int) (areas []model.Area, err error) {
	if len(areaids) <= 0 {
		return areas, errors.New("areaids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ areaid, regionName, regionCode, parentCode, longitude, latitude from ")
	sql_str.WriteString("area")
	sql_str.WriteString(" where areaid in(")
	param_keys := strings.Repeat("?,", len(areaids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(areaids))
	for _, v := range areaids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return areas, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【区域表】表总记录数,使用条件：【区域名称】,【经度】,【纬度】
func (d *area) GetRowCount(regionName string, longitude string, latitude string) (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from area")

	//region 处理regionName
	if golibs.Length(regionName) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("regionName=?")
		params = append(params, regionName)
		conditions++
	}
	//endregion

	//region 处理longitude
	if golibs.Length(longitude) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("longitude=?")
		params = append(params, longitude)
		conditions++
	}
	//endregion

	//region 处理latitude
	if golibs.Length(latitude) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("latitude=?")
		params = append(params, latitude)
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

// 查询【区域表】列表,使用条件：【区域名称】,【经度】,【纬度】
func (d *area) GetRowList(regionName string, longitude string, latitude string, pageIndex, pageSize int) (areas []model.Area, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ areaid, regionName, regionCode, parentCode, longitude, latitude from area")

	//region 处理regionName
	if golibs.Length(regionName) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("regionName=?")
		params = append(params, regionName)
		conditions++
	}
	//endregion

	//region 处理longitude
	if golibs.Length(longitude) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("longitude=?")
		params = append(params, longitude)
		conditions++
	}
	//endregion

	//region 处理latitude
	if golibs.Length(latitude) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("latitude=?")
		params = append(params, latitude)
		conditions++
	}
	//endregion

	//region order by
	sqlString.Append(" order by areaid desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return areas, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【区域表】表记录
func (d *area) _RowsToArray(rows *sqlx.Rows) (areas []model.Area, err error) {
	for rows.Next() {
		m := model.Area{}
		err = rows.Scan(&m.Areaid, &m.RegionName, &m.RegionCode, &m.ParentCode, &m.Longitude, &m.Latitude)
		if err != nil {
			return areas, err
		}
		areas = append(areas, m)
	}
	return areas, err
}

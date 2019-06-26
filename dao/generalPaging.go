package dao

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ha666/golibs"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
	"strings"
	"time"
)

var GeneralPagingDao = generalPaging{}

type generalPaging struct{}

var (
	FieldTypes map[string]string
)

// 初始化所有字段
func initFields(dbName string) {
	FieldTypes = make(map[string]string)
	loadAllTables(dbName)
}

func loadAllTables(dbName string) {
	//region 加载全部表
	tables, err := TableDao.GetTables(dbName)
	if err != nil {
		log.Fatalf("加载所有表时出错:%s", err.Error())
	}
	if len(tables) <= 0 || golibs.Length(tables[0].TableName) <= 0 {
		log.Fatalln("没有找到任何表")
	}
	for _, v := range tables {
		loadAllFields(dbName, v.TableName)
	}
	//endregion
}

func loadAllFields(dbName, tableName string) {

	//reigon 加载所有列
	columns, err := ColumnDao.GetColumns(dbName, tableName)
	if err != nil {
		log.Fatalf("加载所有列时出错:%s", err.Error())
	}
	if len(columns) <= 0 || golibs.Length(columns[0].ColumnName) <= 0 {
		log.Fatalln("没有找到任何字段")
	}
	//endregion

	//region 把字段类型写入缓存
	for i, v := range columns {
		if a, ok := FieldTypes[strings.ToLower(v.ColumnName)]; ok {
			if dbFieldType2GoType(v.DataType) != a {
				log.Fatalf("发现重复：%d\t%s\t%s\t%s\n", i, v.ColumnName, dbFieldType2GoType(v.DataType), a)
			}
		} else {
			FieldTypes[strings.ToLower(v.ColumnName)] = dbFieldType2GoType(v.DataType)
		}
	}
	//endregion
}

// 查询总数
func (d *generalPaging) GeneralPagingTotal(sql string) (total int, err error) {
	rows, err := garglewool.Queryx(sql)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&total)
		if err != nil {
			return -1, err
		}
		return total, nil
	}
	return -1, nil
}

// 查询列表
func (d *generalPaging) GeneralPagingList(sql string) (collectionsArray []map[string]interface{}, err error) {
	rows, err := garglewool.Queryx(sql)
	if err != nil {
		return collectionsArray, err
	}
	defer rows.Close()
	collectionsArray, err = d._RowsToArray(rows)
	if err != nil {
		return
	}
	return collectionsArray, nil
}

func (d *generalPaging) _RowsToArray(rows *sqlx.Rows) (collectionsArray []map[string]interface{}, err error) {
	columns, err := rows.Columns()
	if err != nil {
		return collectionsArray, err
	}
	collectionsArray = make([]map[string]interface{}, 0)
	for rows.Next() {
		a, err := rows.SliceScan()
		if err != nil {
			return collectionsArray, err
		}
		m := make(map[string]interface{}, len(columns))
		for i := 0; i < len(columns); i++ {
			s, ok := FieldTypes[strings.ToLower(columns[i])]
			if !ok {
				return collectionsArray, errors.New(fmt.Sprintf("没有找到字段%s的类型", columns[i]))
			}
			switch s {
			case "int", "int8":
				m[columns[i]], err = strconv.Atoi(golibs.SliceByteToString(a[i].([]uint8)))
				if err != nil {
					return collectionsArray, errors.New(fmt.Sprintf("解析int类型字段%s出错:%s", columns[i], err.Error()))
				}
			case "int64":
				m[columns[i]], err = strconv.ParseInt(golibs.SliceByteToString(a[i].([]uint8)), 10, 64)
				if err != nil {
					return collectionsArray, errors.New(fmt.Sprintf("解析int64类型字段%s出错:%s", columns[i], err.Error()))
				}
			case "string":
				m[columns[i]] = golibs.SliceByteToString(a[i].([]uint8))
			case "time.Time":
				m[columns[i]] = a[i].(time.Time).Format(golibs.Time_TIMEStandard)
			default:
				return collectionsArray, errors.New(fmt.Sprintf("不支持的类型:%s", columns[i]))
			}
		}
		collectionsArray = append(collectionsArray, m)
	}
	return collectionsArray, err
}

// 把数据库字段类型转成go数据类型
func dbFieldType2GoType(field_type string) string {
	type_result := ""
	switch field_type {
	case "bit":
		type_result = "bool"
	case "tinyint":
		type_result = "int8"
	case "smallint":
		type_result = "int16"
	case "int":
		type_result = "int"
	case "bigint":
		type_result = "int64"
	case "float", "decimal", "double", "numeric":
		type_result = "float64"
	case "char", "nchar", "varchar", "nvarchar", "text", "longtext", "mediumtext", "enum", "set":
		type_result = "string"
	case "blob", "longblob", "mediumblob", "tinyblob":
		type_result = "[]byte"
	case "date", "datetime", "datetime2", "timestamp":
		type_result = "time.Time"
	default:
		type_result = "不支持的类型:" + field_type
	}
	return type_result
}

// 把下划线命名法和小驼峰命名法转成大驼峰命名法
func ToBigHump(str *string) {
	var data bytes.Buffer
	is_capitalize := false
	for i, s := range *str {
		if string(s) == "_" {
			is_capitalize = true
		} else {
			if is_capitalize || i == 0 {
				data.WriteString(strings.ToUpper(string(s)))
			} else {
				data.WriteString(string(s))
			}
			is_capitalize = false
		}
	}
	*str = data.String()
}

package dao

import (
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
)

var ColumnDao = column{}

type column struct{}

func (d *column) GetColumns(dbName, table_name string) (column_infos []model.ColumnInfo, err error) {
	rows, err := garglewool.Queryx("SELECT COLUMN_NAME,ORDINAL_POSITION,IS_NULLABLE,DATA_TYPE,COLUMN_TYPE,COLUMN_KEY,EXTRA,COLUMN_COMMENT FROM information_schema.COLUMNS WHERE TABLE_SCHEMA=? AND TABLE_NAME=?;", dbName, table_name)
	defer rows.Close()
	if err != nil {
		return column_infos, err
	}
	return d._ColumnRowsToArray(rows)
}

func (d *column) _ColumnRowsToArray(rows *sqlx.Rows) (models []model.ColumnInfo, err error) {
	for rows.Next() {
		m := model.ColumnInfo{}
		err = rows.Scan(&m.ColumnName, &m.OrdinalPosition, &m.IsNullAble, &m.DataType, &m.ColumnType, &m.ColumnKey, &m.Extra, &m.ColumnComment)
		if err != nil {
			return models, err
		}
		models = append(models, m)
	}
	return models, err
}

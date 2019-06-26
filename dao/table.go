package dao

import (
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
)

var TableDao = table{}

type table struct{}

func (d *table) GetTables(dbName string) (table_infos []model.TableInfo, err error) {
	rows, err := garglewool.Queryx("SELECT TABLE_NAME,TABLE_TYPE,ENGINE,TABLE_COMMENT FROM information_schema.TABLES WHERE TABLE_SCHEMA=?", dbName)
	defer rows.Close()
	if err != nil {
		return table_infos, err
	}
	return d._TableRowsToArray(rows)
}

func (d *table) _TableRowsToArray(rows *sqlx.Rows) (models []model.TableInfo, err error) {
	for rows.Next() {
		m := model.TableInfo{}
		err = rows.Scan(&m.TableName, &m.TableType, &m.Engine, &m.TableComment)
		if err != nil {
			return models, err
		}
		models = append(models, m)
	}
	return models, err
}

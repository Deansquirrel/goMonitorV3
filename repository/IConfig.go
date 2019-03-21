package repository

import "database/sql"

type IConfig interface {
	GetSqlGetConfigList() string
	GetSqlGetConfig() string

	getConfigListByRows(rows *sql.Rows) ([]IConfigData, error)
}

type IConfigData interface {
	IsEqual(c interface{}) bool
}

package repository

import (
	"database/sql"
)

type IHis interface {
	GetSqlHisList() string
	GetSqlHisById() string
	GetSqlHisByConfigId() string
	GetSqlHisByTime() string
	GetSqlSetHis() string
	GetSqlClearHis() string

	getHisListByRows(rows *sql.Rows) ([]IHisData, error)
	getHisSetArgs(data interface{}) ([]interface{}, error)
}

type IHisData interface {
}

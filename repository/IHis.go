package repository

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV3/object"
)

type IHis interface {
	GetSqlHisList() string
	GetSqlHisById() string
	GetSqlHisByConfigId() string
	GetSqlHisByTime() string
	GetSqlSetHis() string
	GetSqlClearHis() string

	getHisListByRows(rows *sql.Rows) ([]object.IHisData, error)
	getHisSetArgs(data object.IHisData) ([]interface{}, error)
}

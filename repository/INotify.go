package repository

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV3/object"
)

type INotify interface {
	GetSqlGetConfig() string
	getConfigListByRows(rows *sql.Rows) ([]object.INotifyData, error)
}

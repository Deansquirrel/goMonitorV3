package repository

import "database/sql"

type INotify interface {
	GetSqlGetConfig() string
	getConfigListByRows(rows *sql.Rows) ([]INotifyData, error)
}

type INotifyData interface {
	GetNotifyId() string
}

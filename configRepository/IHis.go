package configRepository

import (
	"database/sql"
	"time"
)

type IHis interface {
	GetHisList() ([]*interface{}, error)
	GetHisById(id string) (*interface{}, error)
	GetHisByConfigId(id string) ([]*interface{}, error)
	GetHisByTime(begTime time.Time, endTime time.Time) ([]*interface{}, error)
	SetHis(data *interface{}) error
	ClearHis(t time.Duration) error
	getHisListByRows(rows *sql.Rows) ([]*interface{}, error)

	GetSqlHisList() string
	GetSqlHisById() string
	GetSqlHisByConfigId() string
	GetSqlHisByTime() string
	GetSqlSetHis() string
}

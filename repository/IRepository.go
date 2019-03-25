package repository

import (
	"github.com/Deansquirrel/goMonitorV3/object"
	"time"
)

type IConfigRepository interface {
	GetConfigList() ([]object.IConfigData, error)
	GetConfig(id string) (object.IConfigData, error)
}

type IHisRepository interface {
	GetHisList() ([]object.IHisData, error)

	GetHisById(id string) (object.IHisData, error)
	GetHisByConfigId(id string) ([]object.IHisData, error)
	GetHisByTime(begTime, endTime time.Time) ([]object.IHisData, error)
	SetHis(data object.IHisData) error
	ClearHis(t time.Duration) error
}

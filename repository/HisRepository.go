package repository

import (
	"github.com/Deansquirrel/goMonitorV3/object"
	"github.com/Deansquirrel/goToolCommon"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

type hisRepository struct {
	His IHis
}

func NewIntHisRepository() *hisRepository {
	return newHisRepository(&intHis{})
}

func NewWebStateHisRepository() *hisRepository {
	return newHisRepository(&webStateHis{})
}

func NewCrmDzXfTestHisRepository() *hisRepository {
	return newHisRepository(&crmDzXfTestHis{})
}

func newHisRepository(his IHis) *hisRepository {
	return &hisRepository{
		His: his,
	}
}

func (hr *hisRepository) GetHisList() ([]object.IHisData, error) {
	rows, err := comm.getRowsBySQL(hr.His.GetSqlHisList())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return hr.His.getHisListByRows(rows)
}

func (hr *hisRepository) GetHisById(id string) (object.IHisData, error) {
	rows, err := comm.getRowsBySQL(hr.His.GetSqlHisById(), id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return hr.His.getHisListByRows(rows)
}

func (hr *hisRepository) GetHisByConfigId(id string) ([]object.IHisData, error) {
	rows, err := comm.getRowsBySQL(hr.His.GetSqlHisByConfigId(), id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return hr.His.getHisListByRows(rows)
}

func (hr *hisRepository) GetHisByTime(begTime, endTime time.Time) ([]object.IHisData, error) {
	rows, err := comm.getRowsBySQL(hr.His.GetSqlHisByTime(), begTime, endTime)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return hr.His.getHisListByRows(rows)
}

func (hr *hisRepository) SetHis(data object.IHisData) error {
	args, err := hr.His.getHisSetArgs(data)
	if err != nil {
		return err
	}
	err = comm.setRowsBySQL(hr.His.GetSqlSetHis(), args...)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (hr *hisRepository) ClearHis(t time.Duration) error {
	dateP := goToolCommon.GetDateTimeStr(time.Now().Add(-t))
	err := comm.setRowsBySQL(hr.His.GetSqlClearHis(), dateP)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

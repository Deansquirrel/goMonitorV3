package repository

import (
	"github.com/Deansquirrel/goToolCommon"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

type hisRepository struct {
	His IHis
}

func NewHisRepository(his IHis) *hisRepository {
	return &hisRepository{
		His: his,
	}
}

func (hr *hisRepository) GetHisList() ([]IHisData, error) {
	rows, err := comm.getRowsBySQL(hr.His.GetSqlHisList())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return hr.His.getHisListByRows(rows)
}

func (hr *hisRepository) GetHisById(id string) (IHisData, error) {
	rows, err := comm.getRowsBySQL(hr.His.GetSqlHisById(), id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return hr.His.getHisListByRows(rows)
}

func (hr *hisRepository) GetHisByConfigId(id string) ([]IHisData, error) {
	rows, err := comm.getRowsBySQL(hr.His.GetSqlHisByConfigId(), id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return hr.His.getHisListByRows(rows)
}

func (hr *hisRepository) GetHisByTime(begTime, endTime time.Time) ([]IHisData, error) {
	rows, err := comm.getRowsBySQL(hr.His.GetSqlHisByTime(), begTime, endTime)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return hr.His.getHisListByRows(rows)
}

func (hr *hisRepository) SetHis(data IHisData) error {
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

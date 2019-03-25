package repository

import (
	"database/sql"
	"errors"
	"github.com/Deansquirrel/goMonitorV3/object"
	"time"
)
import log "github.com/Deansquirrel/goToolLog"

const SqlGetWebStateTaskHis = "" +
	"SELECT FId, FConfigId, FUseTime, FHttpCode, FContent, FOprTime" +
	" FROM WebStateTaskHis"
const SqlGetWebStateTaskHisById = "" +
	"SELECT FId, FConfigId, FUseTime, FHttpCode, FContent, FOprTime" +
	" FROM WebStateTaskHis" +
	" WHERE FId = ?"
const SqlGetWebStateTaskHisByConfigId = "" +
	"SELECT FId, FConfigId, FUseTime, FHttpCode, FContent, FOprTime" +
	" FROM WebStateTaskHis" +
	" WHERE FConfigId = ?"
const SqlGetWebStateTaskHisByTime = "" +
	"SELECT FId, FConfigId, FUseTime, FHttpCode, FContent, FOprTime" +
	" FROM WebStateTaskHis" +
	" WHERE [FOprTime] >= ? AND [FOprTime] <= ?"

const SqlSetWebStateTaskHis = "" +
	"INSERT INTO WebStateTaskHis (FId, FConfigId, FUseTime, FHttpCode, FContent)" +
	" VALUES (?,?,?,?,?)"
const SqlDelWebStateTaskHis = "" +
	"DELETE FROM WebStateTaskHis" +
	" WHERE FOprTime < ?"

type webStateHis struct {
}

func (wsh *webStateHis) GetSqlHisList() string {
	return SqlGetWebStateTaskHis
}

func (wsh *webStateHis) GetSqlHisById() string {
	return SqlGetWebStateTaskHisById
}

func (wsh *webStateHis) GetSqlHisByConfigId() string {
	return SqlGetWebStateTaskHisByConfigId
}

func (wsh *webStateHis) GetSqlHisByTime() string {
	return SqlGetWebStateTaskHisByTime
}

func (wsh *webStateHis) GetSqlSetHis() string {
	return SqlSetWebStateTaskHis
}

func (wsh *webStateHis) GetSqlClearHis() string {
	return SqlDelWebStateTaskHis
}

func (wsh *webStateHis) getHisListByRows(rows *sql.Rows) ([]object.IHisData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fConfigId, fContent string
	var fUseTime, fHttpCode int
	var fOprTime time.Time
	resultList := make([]object.IHisData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&fId, &fConfigId, &fUseTime, &fHttpCode, &fContent, &fOprTime)
		if err != nil {
			break
		}
		config := object.WebStateHisData{
			FId:       fId,
			FConfigId: fConfigId,
			FUseTime:  fUseTime,
			FHttpCode: fHttpCode,
			FContent:  fContent,
			FOprTime:  fOprTime,
		}
		resultList = append(resultList, &config)
	}
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	if rows.Err() != nil {
		log.Error(rows.Err().Error())
		return nil, rows.Err()
	}
	return resultList, nil
}

func (wsh *webStateHis) getHisSetArgs(data object.IHisData) ([]interface{}, error) {
	switch f := data.(type) {
	case object.WebStateHisData:
		result := make([]interface{}, 0)
		result = append(result, f.FId)
		result = append(result, f.FConfigId)
		result = append(result, f.FUseTime)
		result = append(result, f.FHttpCode)
		result = append(result, f.FContent)
		return result, nil
	default:
		return nil, errors.New("webStateHis getHisSetArgs 参数类型错误")
	}
}

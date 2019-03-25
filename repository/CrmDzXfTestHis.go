package repository

import (
	"database/sql"
	"errors"
	"github.com/Deansquirrel/goMonitorV3/object"
	"time"
)
import log "github.com/Deansquirrel/goToolLog"

const SqlGetCrmDzXfTestTaskHis = "" +
	"SELECT [FId], [FConfigId], [FUseTime], [FHttpCode], [FContent], [FOprTime]" +
	" FROM  CrmDzXfTestTaskHis"

const SqlGetCrmDzXfTestTaskHisById = "" +
	"SELECT [FId], [FConfigId], [FUseTime], [FHttpCode], [FContent], [FOprTime]" +
	" FROM  CrmDzXfTestTaskHis" +
	" WHERE FId = ?"

const SqlGetCrmDzXfTestTaskHisByConfigId = "" +
	"SELECT [FId], [FConfigId], [FUseTime], [FHttpCode], [FContent], [FOprTime]" +
	" FROM  CrmDzXfTestTaskHis" +
	" WHERE FConfigId = ?"

const SqlGetCrmDzXfTestTaskHisByTime = "" +
	"SELECT [FId], [FConfigId], [FUseTime], [FHttpCode], [FContent], [FOprTime]" +
	" FROM  CrmDzXfTestTaskHis" +
	" WHERE [FOprTime] >= ? AND [FOprTime] <= ?"

const SqlSetCrmDzXfTestTaskHis = "" +
	"INSERT INTO CrmDzXfTestTaskHis (FId, FConfigId, FUseTime, FHttpCode, FContent)" +
	" VALUES (?, ?, ?, ?, ?)"

const SqlDelCrmDzXfTestTaskHis = "" +
	"DELETE FROM CrmDzXfTestTaskHis" +
	" WHERE FOprTime < ?"

type crmDzXfTestHis struct {
}

func (config *crmDzXfTestHis) GetSqlHisList() string {
	return SqlGetCrmDzXfTestTaskHis
}

func (config *crmDzXfTestHis) GetSqlHisById() string {
	return SqlGetCrmDzXfTestTaskHisById
}

func (config *crmDzXfTestHis) GetSqlHisByConfigId() string {
	return SqlGetCrmDzXfTestTaskHisByConfigId
}

func (config *crmDzXfTestHis) GetSqlHisByTime() string {
	return SqlGetCrmDzXfTestTaskHisByTime
}

func (config *crmDzXfTestHis) GetSqlSetHis() string {
	return SqlSetCrmDzXfTestTaskHis
}

func (config *crmDzXfTestHis) GetSqlClearHis() string {
	return SqlDelCrmDzXfTestTaskHis
}

func (config *crmDzXfTestHis) getHisListByRows(rows *sql.Rows) ([]object.IHisData, error) {
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
		config := object.CrmDzXfTestHisData{
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

func (config *crmDzXfTestHis) getHisSetArgs(data object.IHisData) ([]interface{}, error) {
	switch f := data.(type) {
	case object.CrmDzXfTestHisData:
		result := make([]interface{}, 0)
		result = append(result, f.FId)
		result = append(result, f.FConfigId)
		result = append(result, f.FUseTime)
		result = append(result, f.FHttpCode)
		result = append(result, f.FContent)
		return result, nil
	default:
		return nil, errors.New("crmDzXfTestHis getHisSetArgs 参数类型错误")
	}
}

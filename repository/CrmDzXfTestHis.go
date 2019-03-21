package repository

import (
	"database/sql"
	"errors"
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

type CrmDzXfTestHis struct {
}

type CrmDzXfTestHisData struct {
	FId       string
	FConfigId string
	FUseTime  int
	FHttpCode int
	FContent  string
	FOprTime  time.Time
}

func (config *CrmDzXfTestHis) GetSqlHisList() string {
	return SqlGetCrmDzXfTestTaskHis
}

func (config *CrmDzXfTestHis) GetSqlHisById() string {
	return SqlGetCrmDzXfTestTaskHisById
}

func (config *CrmDzXfTestHis) GetSqlHisByConfigId() string {
	return SqlGetCrmDzXfTestTaskHisByConfigId
}

func (config *CrmDzXfTestHis) GetSqlHisByTime() string {
	return SqlGetCrmDzXfTestTaskHisByTime
}

func (config *CrmDzXfTestHis) GetSqlSetHis() string {
	return SqlSetCrmDzXfTestTaskHis
}

func (config *CrmDzXfTestHis) GetSqlClearHis() string {
	return SqlDelCrmDzXfTestTaskHis
}

func (config *CrmDzXfTestHis) getHisListByRows(rows *sql.Rows) ([]IHisData, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fConfigId, fContent string
	var fUseTime, fHttpCode int
	var fOprTime time.Time
	resultList := make([]IHisData, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(&fId, &fConfigId, &fUseTime, &fHttpCode, &fContent, &fOprTime)
		if err != nil {
			break
		}
		config := CrmDzXfTestHisData{
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

func (config *CrmDzXfTestHis) getHisSetArgs(data interface{}) ([]interface{}, error) {
	switch f := data.(type) {
	case CrmDzXfTestHisData:
		result := make([]interface{}, 0)
		result = append(result, f.FId)
		result = append(result, f.FConfigId)
		result = append(result, f.FUseTime)
		result = append(result, f.FHttpCode)
		result = append(result, f.FContent)
		return result, nil
	default:
		return nil, errors.New("CrmDzXfTestHis getHisSetArgs 参数类型错误")
	}
}
